package telegrambotnative

import(
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"bytes"
	"budget_tracker/cmd/structs"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func SetWebhook() {
	botToken := os.Getenv("TELEGRAM_API_KEY")
	url := "https://adverse-zora-nehemiah-e50fbf0b.koyeb.app/bot"

	resp, err := http.Get(
		fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", botToken, url),
	)
	if err != nil {
		log.Println("❌ SetWebhook error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Println("🔗 SetWebhook response:", string(body))
}

func HandleWebhook(c *gin.Context) {
	var rawBody []byte
	rawBody, _ = io.ReadAll(c.Request.Body)

	log.Println("📩 Incoming Telegram update:")
	log.Println(string(rawBody))

	// Restore body so Gin can bind it
	c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	var update structs.Update
	if err := c.ShouldBindJSON(&update); err != nil {
		log.Println("❌ JSON bind error:", err)
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
	if botToken == "" {
		return fmt.Errorf("TOKEN not set")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	payload := map[string]any{
		"chat_id": chatID,
		"text":    text,
	}

	body, _ := json.Marshal(payload)

	log.Printf("📤 Sending message to chat %d: %s\n", chatID, text)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("❌ sendMessage error:", err)
		return err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	log.Printf(
		"📨 Telegram response (status %d): %s\n",
		resp.StatusCode,
		string(respBody),
	)

	return nil
}

func handleMessage(msg *structs.Message) {
	log.Printf(
		"💬 Message from %s (%d): %s\n",
		msg.From.FirstName,
		msg.Chat.ID,
		msg.Text,
	)

	switch msg.Text {
	case "/start":
		_ = sendMessage(
			msg.Chat.ID,
			fmt.Sprintf("Hello %s! Welcome!", msg.From.FirstName),
		)
	default:
		_ = sendMessage(
			msg.Chat.ID,
			fmt.Sprintf("You said: %s", msg.Text),
		)
	}
}
