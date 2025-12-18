package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"budget_tracker/cmd/structs"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
)
func parseTelegramRequest(c *gin.Context) (*structs.Update, error) {
	var update structs.Update

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	return &update, nil
}
func TelegramWebhook(handler *telegohandler.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var update telego.Update

		if err := c.ShouldBindJSON(&update); err != nil {
			c.AbortWithStatus(400)
			return
		}

		handler.HandleUpdate(update)
		c.Status(200)
	}
}

