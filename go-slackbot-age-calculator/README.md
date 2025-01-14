# Slack Age Calculator Bot

This is a simple Slack bot built using Go and the `slacker` package. The bot calculates a user's age based on their year of birth provided via a Slack message.

## Features
- Accepts the command: `my yob is <year>`.
- Validates the provided year.
- Returns the user's age if the year is valid.

---

## Prerequisites
- Go 1.23+
- A Slack workspace.
- A Slack app with necessary permissions.

---

## Installation

### 1. Clone the Repository
```bash
git clone <repository-url>
cd go-slackbot-age-calculator
```

### 2. Set Up Environment Variables
Create a `.env` file in the root directory and add your Slack tokens:

```plaintext
SLACK_BOT_TOKEN=xoxb-...
SLACK_APP_TOKEN=xapp-...
```

### 3. Install Dependencies
```bash
go mod tidy
```

---

## Usage

### 1. Run the Bot
```bash
go run main.go
```

### 2. Test the Bot in Slack
- Send a message to the bot: `my yob is 1990`
- The bot will respond with your calculated age.

---

## Code Structure
```plaintext
├── main.go               # Main bot logic
├── config
│   └── config.go         # Loads the configuration from the .env file
├── go.mod                # Go module dependencies
├── README.md             # Project documentation
└── .env                  # Environment variables
```

---

## Configuration
Ensure your Slack app has the following **Bot Token Scopes**:
- `chat:write`
- `app_mentions:read`

**Steps to Configure:**
1. Go to [Slack API Dashboard](https://api.slack.com/apps).
2. Select your app and navigate to `OAuth & Permissions`.
3. Add the scopes mentioned above.
4. Reinstall the app in your workspace to apply changes.

---

## Contributing
Feel free to fork the repository and create a pull request with any improvements or bug fixes.

---

## License
This project is licensed under the MIT License.
