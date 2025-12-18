package telegrambotnative

import(
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"bytes"
	"budget_tracker/cmd/structs"
	"github.com/gin-gonic/gin"
)

func SetWebhook() {
	botToken := os.Getenv("TELEGRAM_API_KEY")
	url := "https://adverse-zora-nehemiah-e50fbf0b.koyeb.app:8443/bot"

	http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", botToken, url))
}
func HandleWebhook(c *gin.Context){
		var update structs.Update
		if err := c.ShouldBindJSON(&update); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if update.Message != nil {
			handleMessage(update.Message)
		}

		c.Status(http.StatusOK)
}


func sendMessage(chatID int, text string) error {
	botToken := os.Getenv("TELEGRAM_API_KEY")
	if botToken == "" { return fmt.Errorf("TOKEN not set")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	payload := map[string] any {
		"chat_id": chatID,
		"text":    text,
	}

	body, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
func handleMessage(msg *structs.Message) {
	text := msg.Text
	chatID := msg.Chat.ID

	switch text {
		case "/start":
			sendMessage(chatID, fmt.Sprintf("Hello %s! Welcome!", msg.From.FirstName))
		default:
			sendMessage(chatID, fmt.Sprintf("You said: %s", text))
		}
}

