package telegrambot

import (
	"context"
	"fmt"
	"os"
	"github.com/mymmrac/telego"
)

func Instantiate() (*telego.Bot) {
	botToken := os.Getenv("TELEGRAM_API_KEY")
	if botToken == "" {
		panic("TELEGRAM_API_KEY not set")
	}

	bot, err := telego.NewBot(
		botToken,
		telego.WithWarnings(),
	)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	me, err := bot.GetMe(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bot connected: @%s (ID: %d)\n", me.Username, me.ID)
	return bot
}
