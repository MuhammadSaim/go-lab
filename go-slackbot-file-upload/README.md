# Slack File Upload Bot

This Go project uploads files to a Slack channel using the `slack-go` package.

## Features
- Uploads files to a specified Slack channel
- Shares the file link in the channel after upload
- Configurable using a `config` package

## Prerequisites
- Go 1.23+
- Slack bot token with the following scopes:
  - `files:write`
  - `files:read`
  - `chat:write`

## Configuration
Ensure you have a `config` package that provides the following fields:

```go
cfg.SLACK_BOT_TOKEN  // Your Slack bot token (xoxb-...)
cfg.SLACK_CHANNEL_ID // Channel ID where the file will be uploaded
```

## Project Structure
```plaintext
.
├── config/                  # Config package for managing credentials
├── files/                   # Directory to store files for upload
├── main.go                  # Main entry point of the project
├── README.md                # Project documentation
├── go.mod                   # Go module file
└── go.sum                   # Dependency lock file
```

## Usage
1. **Clone the repository:**
   ```bash
   git clone <repo-url>
   cd go-slackbot-file-upload
   ```
2. **Set up your Slack bot token and channel ID:**
   - Ensure you have a valid bot token and channel ID in your `config` package.

3. **Run the bot:**
   ```bash
   go run main.go
   ```

4. **Verify:**
   - The file should be uploaded to the specified Slack channel.

## Troubleshooting
- Ensure the bot has been invited to the channel using `/invite @BotName`.
- Confirm your bot token has the correct scopes.
- If the file uploads but is not visible, ensure the `PostMessage` method is called after the upload.

## Contributing
Feel free to fork and submit pull requests to improve the project.

## License
This project is open-source under the MIT License.


