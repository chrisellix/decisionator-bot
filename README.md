<h1>Telegram Decision Bot 🤖</h1>

Setup Instructions
1. Create a Bot via BotFather
Go to @BotFather on Telegram

Send /newbot and follow the steps

Copy the bot token given at the end

2. Add Your Token to the Code
In your Go file replace:

const BotToken = "YOUR_BOT_TOKEN_HERE"

with your actual token:

const BotToken = "123456789:ABCdefGhIjkLMNopQrstUVwxYZ"

3. Install Telegram Bot API
go get github.com/go-telegram-bot-api/telegram-bot-api/v5
4. Run the Bot
go run main.go

5. Start Chatting
Open your bot on Telegram

Send /start

Ask a question

The bot will respond with a random decision! :) 
