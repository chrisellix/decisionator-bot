package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"strings"
)

const TOKEN = "7908084075:AAFo-E05orcfGSqjPk_Vfo0tZnTpU2saJLc"

var bot *tgbotapi.BotAPI

var ChatId int64

var forDecisinator = [5]string{"decisinator", "decinator", "decisin", "desinator", "decide"}

var answers = []string{
	"Yes",
	"No",
	"You can, you should, and if you’re brave enough to start, you will.",
	"The scariest moment is always just before you start.",
	"Get busy living or get busy dying.",
	"Monsters are real, and ghosts are real too. They live inside us, and sometimes, they win.",
	"You can't deny laughter; when it comes, it plops down in your favorite chair and stays as long as it wants.",
	"A person can’t change all at once.",
	"The trust of the innocent is the liar's most useful tool.",
	"We lie best when we lie to ourselves.",
	"There’s no harm in hoping for the best as long as you’re prepared for the worst.",
	"Home is where you dance with others, and dancing is life.",
	"We never know which lives we influence, or when, or why.",
	"A secret’s a strange thing.",
	"The world has teeth and it can bite you any time it wants.",
	"You can't be careful on a skateboard.",
	"Sometimes human places create inhuman monsters.",
	"Good books don't give up all their secrets at once.",
	"Remember, hope is a good thing, maybe the best of things, and no good thing ever dies.",
	"Perfect paranoia is perfect awareness.",
	"The mind can calculate, but the spirit yearns, and the heart knows what the heart knows.",
	"The most important things are the hardest to say.",
}

func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telegram API")
	}
}

func getDecisinatorsAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}
func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(ChatId, msg)
	bot.Send(msgConfig)
}

func isMessageForDecisinator(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}
	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range forDecisinator {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}

	}
	return false
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(ChatId, getDecisinatorsAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}
func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			ChatId = update.Message.Chat.ID
			sendMessage("Hi and welcome to Decisinator! You can ask your questions by " +
				"calling my name and get my opinion on that." +
				"For example: \"Decisinator, am I ready to change my current job?\" or " +
				"\"Decisinator, do I really wanna go to that party?\"")
		}

		if isMessageForDecisinator(&update) {
			sendAnswer(&update)
		}
	}
}
