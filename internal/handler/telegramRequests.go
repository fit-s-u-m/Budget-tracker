package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"budget_tracker/cmd/structs"
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
func TelegramWebhook(c *gin.Context) {
	update, err := parseTelegramRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid telegram payload",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
