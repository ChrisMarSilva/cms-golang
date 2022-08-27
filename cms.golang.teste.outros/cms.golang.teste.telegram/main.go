package main

import (
	"log"
	"fmt"
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.telegram
// go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
// go mod tidy

// go run main.go

// https://go-telegram-bot-api.dev/

const tgbotapiKey = "long-long-tgbot-api-key"

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🏠 Главная"),
		tgbotapi.NewKeyboardButton("🗒 Запись"),
	),
)

var courseMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Golang"),
		tgbotapi.NewKeyboardButton("Intense golang"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("HighLoad"),
		tgbotapi.NewKeyboardButton("VueJS"),
	),
)

var courseSignMap map[int]*finbot.CourseSign

func init() {
	courseSignMap = make(map[int]*finbot.CourseSign)
}


func main_outrps() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}


func main_2() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhookWithCert("https://www.example.com:8443/"+bot.Token, "cert.pem")

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}

func main() {

	var (
		bot        *tgbotapi.BotAPI
		err        error
		updChannel tgbotapi.UpdatesChannel
		update     tgbotapi.Update
		updConfig  tgbotapi.UpdateConfig
		botUser    tgbotapi.User
	)
	bot, err = tgbotapi.NewBotAPI(tgbotapiKey)
	if err != nil {
		panic("bot init error: " + err.Error())
	}

	botUser, err = bot.GetMe()
	if err != nil {
		panic("bot getme error: " + err.Error())
	}

	fmt.Printf("auth ok! bot is: %s\n", botUser.FirstName)

	updConfig.Timeout = 60
	updConfig.Limit = 1
	updConfig.Offset = 0

	updChannel, err = bot.GetUpdatesChan(updConfig)
	if err != nil {
		panic("update channel error: " + err.Error())
	}

	for {

		update = <-updChannel

		if update.Message != nil {

			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				if cmdText == "test" {
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"test cmd")
					bot.Send(msgConfig)
				} else if cmdText == "menu" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Главное меню")
					msg.ReplyMarkup = mainMenu
					bot.Send(msg)
				}
			} else {

				if update.Message.Text == mainMenu.Keyboard[0][1].Text {

					courseSignMap[update.Message.From.ID] = new(finbot.CourseSign)
					courseSignMap[update.Message.From.ID].State = finbot.StateEmail

					fmt.Printf(
						"message: %s\n",
						update.Message.Text)

					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"Введите email:")
					msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msgConfig)
				} else {
					cs, ok := courseSignMap[update.Message.From.ID]
					if ok {
						if cs.State == finbot.StateEmail {
							cs.Email = update.Message.Text
							msgConfig := tgbotapi.NewMessage(
								update.Message.Chat.ID,
								"Введите телефон:")
							bot.Send(msgConfig)
							cs.State = 1
						} else if cs.State == finbot.StateTel {
							cs.Telephone = update.Message.Text
							cs.State = 2
							msgConfig := tgbotapi.NewMessage(
								update.Message.Chat.ID,
								"Введите course:")
							msgConfig.ReplyMarkup = courseMenu
							bot.Send(msgConfig)
						} else if cs.State == finbot.StateCourse {
							cs.Course = update.Message.Text
							msgConfig := tgbotapi.NewMessage(
								update.Message.Chat.ID,
								"ok!")
							msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
							bot.Send(msgConfig)
							delete(courseSignMap, update.Message.From.ID)
							//  post to site!
							err = post.SendPost(cs)
							if err != nil {
								fmt.Printf("send post error: %v\n", err)
							}
						}
						fmt.Printf("state: %+v\n", cs)
					} else {
						// other messages
						msgConfig := tgbotapi.NewMessage(
							update.Message.Chat.ID,
							"ok")
						msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						bot.Send(msgConfig)
					}
				}
			}
		} else {
			fmt.Printf("not message: %+v\n", update)
		}
	}

	bot.StopReceivingUpdates()
}

type SignState int

const (
	StateEmail SignState = iota
	StateTel
	StateCourse
)

type CourseSign struct {
	State     SignState // 0 - email, 1 - tel, 2 - course
	Name      string
	Email     string
	Telephone string
	Course    string
}

const postURL = "http://finalistx.com/email.php"

func SendPost(data *finbot.CourseSign) error {

	params := fmt.Sprintf(
		"name=%s&email=%s&tel=%s&course=%s",
		data.Name,
		data.Email,
		data.Telephone,
		data.Course)

	buf := bytes.NewBufferString(params)
	resp, err := http.Post(
		postURL,
		"application/x-www-form-urlencoded",
		buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ba, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("response: %s\n", ba)

	if resp.StatusCode != 200 {
		return errors.New("not 200 response")
	}
	return nil
}