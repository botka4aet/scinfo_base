package main

//https://habr.com/ru/post/351060/
//https://habr.com/ru/post/446468/
//go get -u
//go mod tidy
import (
	"os"
	"reflect"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TelegramBot() {
	Info := New("tele")

	//Создаем бота
	bot, err := tgbotapi.NewBotAPI(Info.TeleConfig.Token)
	if err != nil {
		panic(err)
	}

	//Устанавливаем время обновления
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	//Получаем обновления от бота
	updates := bot.GetUpdatesChan(updateConfig)
	cuserrights := false
	for update := range updates {
		if update.Message == nil {
			continue
		}

		cuserrights = Contains(Info.TeleConfig.telehigh, strconv.FormatInt(update.Message.From.ID, 10))
		//Проверяем что от пользователья пришло именно текстовое сообщение
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			tmsg := ""
			cmdmsg := strings.Split(update.Message.Text, " ")[0]
			switch cmdmsg {
			case "/start":

				if cuserrights {
					tmsg = "Команды админа: /add,/delete,/update,/updateforce"
				} else {
					tmsg = "Команды пользователя:"
				}

				//Отправлем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, tmsg)
				bot.Send(msg)

			default:

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда")
				bot.Send(msg)
			}
		} else {

			//Отправлем сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда")
			bot.Send(msg)
		}
	}
}

// Проверяем права
func Contains(a []string, b string) bool {
	for _, n := range a {
		if b == n {
			return true
		}
	}
	return false
}

func TGstart() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		// Now that we know we've gotten a new message, we can construct a
		// reply! We'll take the Chat ID and Text from the incoming message
		// and use it to create a new message.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// We'll also say that this message is a reply to the previous message.
		// For any other specifications than Chat ID or Text, you'll need to
		// set fields on the `MessageConfig`.
		msg.ReplyToMessageID = update.Message.MessageID

		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		if _, err := bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}
}
