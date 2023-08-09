# Zig Reporter Bot

This repository contains a Go application designed to fetch RSS feeds and relay new items to the @ziglang_en Telegram chat.

## Features

- Fetches Ziglang news from the official Zig website's RSS feed.
- Notifies about new Zig releases by probing tags from GitHub's Atom feed.
- Obtains development versions in a JSON format from Zig's official site.
- Sends the aggregated data in a well-structured message to a Telegram chat.
- Ensures no duplicate notifications by remembering the last fetched item.

## How it works

The application leverages [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) for Telegram Bot API interactions and the [mmcdole/gofeed](https://github.com/mmcdole/gofeed) library for parsing RSS feeds.

Set to run every 15 minutes via GitHub Actions, the application checks the RSS feeds and identifies any new items. Upon locating a new item, it dispatches a message to the designated Telegram chat. 

The last item from each feed is persistently stored in a file, guaranteeing only new items are relayed to the Telegram chat.

## Setup

 ### Export your Telegram bot token as an environment variable:

```bash
export TELEGRAM_BOT_TOKEN=your_bot_token_here
```
This token can also be securely stored as a secret in your GitHub repository settings.

### Configuration

- **Feeds**: The `feeds` map within `main.go` houses the RSS feeds and their associated message formats. You can conveniently expand or adjust this collection to include other sources. The map structure contains:
  - Feed URL as the key.
  - `FeedConfig` object detailing the message template and the fields to extract from the feed.
  
- **Telegram Chat ID**: The constant `ChatID` in the application corresponds to the Telegram chat ID where messages will be dispatched. Adjust this value if targeting a different chat.

- **Storage**: The application saves the ID of the last processed item from each feed in individual files, grouped under the `last` directory. Ensure that the application has the necessary permissions to read from and write to this directory.

### Running Locally

1. Clone the repository:
```bash
git clone https://github.com/your_repository_name/zig_reporter_bot.git
```
2. Navigate into the directory:
```bash
cd zig_reporter_bot
```
3. Run the application:
```bash
go run main.go
```

## Deployment

The application is configured to run through GitHub Actions, which activates every 15 minutes. Ensure that the `TELEGRAM_BOT_TOKEN` is securely saved within the secrets of your GitHub repository, so that GitHub Actions can retrieve it.

Should you wish to modify the frequency or behavior of this action, you can edit the respective `.yml` file located in the `.github/workflows` directory.

## Contribution

All forms of contributions to this project are highly encouraged and appreciated. If you discover bugs or have ideas for features, please don't hesitate to open an issue on GitHub. When planning to submit a pull request, do ensure to test your changes exhaustively.

## License

This project falls under the MIT License. For more details, refer to the `LICENSE` file in the repository.
