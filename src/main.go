package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var (
	// глобальная переменная в которой храним токен
	telegramBotToken string
)

func init() {
	token := os.Environ()
	fmt.Println(token)
	// принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", "", "Telegram Bot Token")
	flag.Parse()

	// без него не запускаемся
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}

func main() {
	// используя токен создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		// универсальный ответ на любое сообщение
		// reply := "Не знаю что сказать"
		// if update.Message == nil {
		// 	continue
		// }
		reply := ""
		switch update.Message.Text {
		case "привет":
			reply = "Привет. Я телеграм-бот"
		case "hello":
			reply = "world"
		default:
			reply = "отправь мне команду /azino777"
		}
		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// свитч на обработку комманд
		// комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "start":
			reply = "Привет. Я телеграм-бот"
		case "hello":
			reply = "world"

		case "image":
			reply = "world"

		case "azino777":
			azino := "А-а-азино три топора ннначилась игра"
			rand.Seed(time.Now().UnixNano())
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, azino))
			x := strconv.Itoa(rand.Intn(10))
			y := strconv.Itoa(rand.Intn(10))
			z := strconv.Itoa(rand.Intn(10))

			reply = x + " : " + y + " : " + z
			if x == y && y == z {
				vezenie := "повезло повезло"
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, vezenie))
			} else if x == y || y == z || x == z {
				vezenie := "повезло"
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, vezenie))
			}
		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)

	}
}
