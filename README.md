<h1>Telegram Decision Bot 🤖</h1>

<h2><strong>Setup Instructions</strong></h2>
</h4>1. Create a Bot via BotFather</h4>
Go to @BotFather on Telegram

Send /newbot and follow the steps

Copy the bot token given at the end
<hr>
<h4>2. Add Your Token to the Code</h4>
In your Go file replace:

const BotToken = "YOUR_BOT_TOKEN_HERE"

with your actual token:

const BotToken = "123456789:ABCdefGhIjkLMNopQrstUVwxYZ"
<hr>
<h4>3. Install Telegram Bot API </h4>
go get github.com/go-telegram-bot-api/telegram-bot-api/v5
<hr>
4. Run the Bot
go run main.go

<hr>
<h4><5. Start Chatting</h4>
Open your bot on Telegram

Send /start

Ask a question

The bot will respond with a random decision! :) 
