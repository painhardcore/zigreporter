package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gofeed "github.com/mmcdole/gofeed"
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
	ioutil.WriteFile("last"+string(os.PathSeparator)+getFileName(url), []byte(lastItem), 0o644)
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

	fmt.Printf("New item found in feed %s\n", url)

	fieldValues := getFieldValues(latestItem, feedConfig.Fields)
	message := fmt.Sprintf(feedConfig.Template, fieldValues...)
	sendMessage(botToken, ChatID, message, false)
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

	message := generateMessage(latestVersion)
	sendMessage(botToken, ChatID, message, true)
	setLastItem(url, latestVersion)
}

func generateMessage(version string) string {
	parts := strings.Split(version, "+")
	if len(parts) != 2 {
		return fmt.Sprintf("🚀 New dev version: [%s](https://ziglang.org/download)", version)
	}

	versionPart := parts[0]
	commitHash := parts[1]

	return fmt.Sprintf(
		"🚀 New dev version: [%s](https://ziglang.org/download) | Commit: [%s](https://github.com/ziglang/zig/commits/%s)",
		versionPart, commitHash, commitHash)
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
