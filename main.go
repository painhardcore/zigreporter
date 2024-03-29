package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gofeed "github.com/mmcdole/gofeed"
	"github.com/painhardcore/zigreporter/word"
)

type FeedConfig struct {
	Template string   // Message template
	Fields   []string // Fields to use in the template
}

const (
	// ChatID is the ID of the Telegram chat to send messages to
	ChatID int64 = -1001533442735
)

// RSS feeds and their message format
var feeds = map[string]FeedConfig{
	"https://ziglang.org/news/index.xml":       {Template: "Ziglang News 📰\n[%s](%s)\n\n%s [more](%s)", Fields: []string{"Title", "Link", "Description", "Link"}},
	"https://github.com/ziglang/zig/tags.atom": {Template: "🚀 New Zig release: *%s*\n\n[Link to release](%s)", Fields: []string{"Title", "Link"}},
}

// Retrieve the last processed item from the file
func getLastItem(url string) string {
	data, err := os.ReadFile("last" + string(os.PathSeparator) + getFileName(url))
	if err != nil {
		return ""
	}
	return string(data)
}

// Save the ID of the last processed item to the file
func setLastItem(url string, lastItem string) {
	_ = os.WriteFile("last"+string(os.PathSeparator)+getFileName(url), []byte(lastItem), 0o644)
}

// Generate a file name based on the feed URL
func getFileName(url string) string {
	// replace all slashes in the URL with underscores
	return strings.ReplaceAll(url, "/", "_") + ".txt"
}

// Send a message to a Telegram chat
func sendMessage(botToken string, chatID int64, text string, disablePreview bool) {
	bot, _ := tgbotapi.NewBotAPI(botToken)
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = disablePreview
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatalf("Error sending message: %s", err)
	}
}

// Process a feed: fetch the latest item and send it to the Telegram chat if it's new
func processFeed(fp *gofeed.Parser, url string, feedConfig FeedConfig, botToken string) {
	feed, _ := fp.ParseURL(url)

	if len(feed.Items) == 0 {
		fmt.Printf("Feed %s is empty.\n", url)
		return
	}

	latestItem := feed.Items[0]
	lastItem := getLastItem(url)

	if strings.Compare(lastItem, latestItem.GUID) == 0 {
		fmt.Printf("No new item found in feed %s\n", url)
		return
	}

	fmt.Printf("New items found in feed: %s\n", url)
	found := false
	// Iterate over the items in reverse order to send the oldest first
	for i := len(feed.Items) - 1; i >= 0; i-- {
		if strings.Compare(lastItem, feed.Items[i].GUID) == 0 {
			found = true
			continue
		}
		if found {
			fieldValues := getFieldValues(feed.Items[i], feedConfig.Fields)
			message := fmt.Sprintf(feedConfig.Template, fieldValues...)
			sendMessage(botToken, ChatID, message, false)
			setLastItem(url, feed.Items[i].GUID)
		}
	}
	// Save the latest item to the file
	setLastItem(url, latestItem.GUID)
}

// getFieldValues retrieves the values of the specified fields from a feed item.
// It takes an item of type *gofeed.Item and a slice of strings representing the fields to retrieve.
// It returns a slice of interface{} containing the values of the specified fields.
func getFieldValues(item *gofeed.Item, fields []string) []interface{} {
	fieldValues := make([]interface{}, len(fields))
	for i, fieldName := range fields {
		value := reflect.ValueOf(item).Elem().FieldByName(fieldName).Interface()
		if str, ok := value.(string); ok {
			value = escapeMarkdown(html.UnescapeString(str))
		}
		fieldValues[i] = value
	}
	return fieldValues
}

// Escape Markdown special characters in a string
func escapeMarkdown(text string) string {
	for _, char := range []string{"_", "*", "[", "]", "(", ")", "~", "`", "#"} {
		text = strings.ReplaceAll(text, char, "\\"+char)
	}
	return text
}

type ZigVersion struct {
	Master struct {
		Version string `json:"version"`
	} `json:"master"`
}

func processJSONFeed(url string, botToken string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var versionData ZigVersion
	if err := decoder.Decode(&versionData); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
		return
	}

	latestVersion := versionData.Master.Version
	lastVersion := getLastItem(url)

	if strings.Compare(lastVersion, latestVersion) == 0 {
		fmt.Printf("No new version found in feed %s\n", url)
		return
	}

	message := generateMessage(latestVersion, lastVersion)
	sendMessage(botToken, ChatID, message, true)
	setLastItem(url, latestVersion)
}

func generateMessage(latestVersion, lastVersion string) string {
	partsLatest := strings.Split(latestVersion, "+")
	partsLast := strings.Split(lastVersion, "+")

	if len(partsLatest) != 2 || len(partsLast) != 2 {
		return fmt.Sprintf("🚀 New dev version: [%s](https://ziglang.org/download)", latestVersion)
	}

	latestCommitHash := partsLatest[1]
	lastCommitHash := partsLast[1]

	return fmt.Sprintf(
		"🚀 New %s dev version: [%s](https://ziglang.org/download) | [Changes](https://github.com/ziglang/zig/compare/%s...%s)",
		word.RandomReleaseWord(), partsLatest[0], lastCommitHash, latestCommitHash)
}

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		fmt.Println("TELEGRAM_BOT_TOKEN is required")
		os.Exit(1)
	}

	fp := gofeed.NewParser()

	for url, feedConfig := range feeds {
		processFeed(fp, url, feedConfig, botToken)
	}
	processJSONFeed("https://ziglang.org/download/index.json", botToken)
}
