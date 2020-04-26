// Wii Telegram bot.
// Author:
//    Andrea Cervesato <andrea.cervesato@mailbox.org>

package main

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
	"gopkg.in/tucnak/telebot.v2"
)

// Telegram bot definition.
type WiiBot struct {
	Token         string
	ImgurClientID string
	Bot           *telebot.Bot
	ImgurClient   *Imgur
	IRCQuotesFile *IRCQuotes
}

// Reply with a message to a chat.
func (obj *WiiBot) SendMessage(msg *telebot.Message, text string) {
	log.Printf("Message: %s\n", text)
	obj.Bot.Send(msg.Chat, text)
}

// Reply with an IRC quote to a chat.
func (obj *WiiBot) SendIRCQuote(msg *telebot.Message) {
	text := obj.IRCQuotesFile.RandQuote()

	log.Printf("Quote: %s\n", text)
	obj.Bot.Send(msg.Chat, text)
}

// Reply with a sticker to a chat.
func (obj *WiiBot) SendSticker(msg *telebot.Message, location string) {
	sticker := &telebot.Sticker{File: telebot.FromDisk(location)}

	log.Printf("Sticker: %s\n", location)
	obj.Bot.Send(msg.Chat, sticker)
}

// Reply with an image taken from imgur subreddit to a chat.
func (obj *WiiBot) SendSubredditImage(msg *telebot.Message, gallery string) {
	log.Printf("Getting image from gallery '%s'\n", gallery)
	image := obj.ImgurClient.RandSubredditImage(gallery)

	log.Printf("Image: %s\n", image)
	obj.Bot.Send(msg.Chat, image)
}

// Create a new bot.
func NewBot() *WiiBot {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
	}

	// read telegram section from config file
	telegram_section, err := cfg.GetSection("telegram")
	if err != nil {
		log.Fatal(err)
	}

	// read imgur section from config file
	imgur_section, err := cfg.GetSection("imgur")
	if err != nil {
		log.Fatal(err)
	}

	if !telegram_section.HasKey("token") {
		log.Fatal("telegram token is not defined")
	}

	if !imgur_section.HasKey("client_id") {
		log.Fatal("imgur client_id is not defined")
	}

	// create telegram bot
	token := telegram_section.Key("token").String()
	imgur_client_id := imgur_section.Key("client_id").String()

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	obj := WiiBot{
		Token:         token,
		ImgurClientID: imgur_client_id,
		Bot:           bot,
		ImgurClient:   NewImgur(imgur_client_id),
		IRCQuotesFile: NewIRCQuotes("data/quotes.txt"),
	}

	return &obj
}

func main() {
	log.Println("Creating bot")
	wii := NewBot()

	log.Println("Registering messages")
	wii.Bot.Handle("/help", func(msg *telebot.Message) {
		log.Println("/help")
		wii.Bot.Send(msg.Chat,
			"These are not the commands you are looking for..\n\n"+
				"/help: show this message\n"+
				"/irc_quotes: show irc quote\n"+
				"/bycicle: Byyyycicle Byyyycicle\n"+
				"/fap: FAP FAP FAP\n"+
				"/lamerda: everything is La Merda\n"+
				"/ftttt: ftttt ftttt\n"+
				"/russia: random pic from Russia\n"+
				"/startrek: random pic from Star Trek\n"+
				"/cats: random pic a cute cats\n"+
				"/dogs: random pic of a cute dog\n"+
				"/nintendo: random pic of Nintendo stuff\n"+
				"/mario: random pic of Mario\n"+
				"/doge: wow\n")
	})
	wii.Bot.Handle("/irc_quote", func(msg *telebot.Message) {
		wii.SendIRCQuote(msg)
	})
	wii.Bot.Handle("/ftttt", func(msg *telebot.Message) {
		wii.SendMessage(msg, "@valedix https://i.imgur.com/3STgUHv.jpg")
	})
	wii.Bot.Handle("/bycicle", func(msg *telebot.Message) {
		wii.SendSticker(msg, "data/bycicle.webp")
	})
	wii.Bot.Handle("/fap", func(msg *telebot.Message) {
		wii.SendSticker(msg, "data/fap.webp")
	})
	wii.Bot.Handle("/lamerda", func(msg *telebot.Message) {
		wii.SendSticker(msg, "data/lamerda.webp")
	})
	wii.Bot.Handle("/russia", func(msg *telebot.Message) {
		wii.SendSubredditImage(msg, "ANormalDayInRussia")
	})
	wii.Bot.Handle("/startrek", func(msg *telebot.Message) {
		wii.SendSubredditImage(msg, "startrekgifs")
	})
	wii.Bot.Handle("/cats", func(msg *telebot.Message) {
		wii.SendSubredditImage(msg, "catgifs")
	})
	wii.Bot.Handle("/dogs", func(msg *telebot.Message) {
		wii.SendSubredditImage(msg, "doggifs")
	})
	wii.Bot.Handle("/nintendo", func(msg *telebot.Message) {
		wii.SendSubredditImage(msg, "nintendo")
	})
	wii.Bot.Handle("/mario", func(msg *telebot.Message) {
		wii.SendSubredditImage(msg, "mario")
	})
	wii.Bot.Handle("/doge", func(msg *telebot.Message) {
		wii.SendSubredditImage(msg, "doge")
	})

	log.Println("Starting bot")
	wii.Bot.Start()
}
