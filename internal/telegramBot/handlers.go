package telegrambot

import (
	"fmt"
 "github.com/mymmrac/telego"
  th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleAll (ctx *th.Context, update telego.Update) error{

	 _, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
	))

 return nil
}
// func (bot telego.Bot) HandleMessage (ctx context.Context,update telego.Update) {
// 		if update.Message != nil {
// 			bot.SendMessage(telego.SendMessageParams{
// 				ChatID: telego.ChatID{ID: update.Message.Chat.ID},
// 				Text:   "Hello via webhook 🚀",
// 			})
// 		}
// }
