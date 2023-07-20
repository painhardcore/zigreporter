# Zig reporter bot

This repository contains a simple Go application that fetches RSS feeds and sends new items to a @ziglang_en Telegram chat.

## How it works

The application uses the [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) to interact with the Telegram Bot API and the [mmcdole/gofeed](https://github.com/mmcdole/gofeed) library to parse RSS feeds.

Every 15 minutes (configured via GitHub Actions), the application fetches the RSS feeds and checks for new items. If a new item is found, a message is sent to the specified Telegram chat.

The last processed item from each feed is saved in a file, ensuring that only new items are sent to the Telegram chat.

## Setup

To run this application, you need to set the following environment variable:

- `TELEGRAM_BOT_TOKEN`: The token of your Telegram bot.

This token can be stored as a secret in your GitHub repository settings.

You can add more RSS feeds by updating the `feeds` map in `main.go`.

Each feed has a message format, which is a string with placeholders for the fields of the RSS feed items. You can customize the message format for each feed.

## Contribution

If you find any bugs or have a feature request, please open an issue on github!
