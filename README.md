# Slack Age Calculator Bot

## Table of Contents

- [About](#about)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [MIT License](./LICENSE)

## About

The Slack Age Calculator Bot is a simple Slack bot that calculates a user's age based on their year of birth. The bot listens for a specific command in Slack, processes the input, and returns the user's age. This project demonstrates how to use the Slacker library for building Slack bots in Go, as well as handling environment variables and context management.

## Getting Started

These instructions will help you set up and run the Slack Age Calculator Bot on your local machine for development and testing purposes.

### Prerequisites

To run this project, you'll need the following:

- Go (Golang) installed on your machine.
- A Slack workspace with a bot token and app token.
- The `godotenv` package for managing environment variables.
- The `slacker` package for interacting with the Slack API.

To install the required Go packages, run:

```bash
go get -u github.com/joho/godotenv
go get -u github.com/shomali11/slacker
```

### Installing

1. Clone the repository:

   ```bash
   git clone https://github.com/daking24/slack-age-bot.git
   cd slack-age-calculator-bot
   ```

2. Create a `.env` file:

   ```env
   SLACK_BOT_TOKEN=your-slack-bot-token
   SLACK_APP_TOKEN=your-slack-app-token
   ```

   Replace your-slack-bot-token and your-slack-app-token with your actual Slack tokens.

3. Run the application:

   ```bash
   go run main.go
   ```

   The bot will start listening for the command `my year-of-birth is <year>` in your Slack workspace. When you provide a year, the bot will respond with your calculated age.

## Usage

To use the Slack Age Calculator Bot, invite it to a Slack channel and type the following command:

```text
my year-of-birth is 1998
```

The bot will reply with your age based on the current year (2024).

You can customize the bot's behavior by modifying the command definition in the `main.go` file. For example, you could add new commands or change the age calculation logic. The `printCommandEvents` function also logs command events to the console, which can be useful for debugging.
