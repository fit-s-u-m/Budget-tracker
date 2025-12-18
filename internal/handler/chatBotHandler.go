package handler

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/mistral"
)
func ChatBot(c *gin.Context) {
  apiKey := os.Getenv("MISTRAL_API_KEY")
 	if apiKey == "" {
 		c.JSON(http.StatusInternalServerError, gin.H{
 			"error": "GEMINI_API_KEY not set",
 		})
 		return
   }
   // Create a context
   ctx := context.Background()

		llm, err := mistral.New (
				mistral.WithAPIKey(apiKey),
				mistral.WithModel("mistral-tiny"),
		)

    if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
    }

    // Send a message to the LLM
    response, err := llms.GenerateFromSinglePrompt(
        ctx,
        llm,
        "Hello! How can you help me today?",
    )
    if err != nil {
				c.JSON(http.StatusTooManyRequests, gin.H{
					"error": err.Error(),
				})
				return
    }

	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}

func ChatBotFunc(message string) (response string, err error) {
  apiKey := os.Getenv("MISTRAL_API_KEY")
	if apiKey == "" {
			return "", errors.New("Api key is not found")
	}
   // Create a context
   ctx := context.Background()

		llm, err := mistral.New (
				mistral.WithAPIKey(apiKey),
				mistral.WithModel("mistral-tiny"),
		)

    if err != nil {
			  return "", err
    }

    // Send a message to the LLM
    response, err = llms.GenerateFromSinglePrompt(
        ctx,
        llm,
        "Hello! How can you help me today?",
    )
    if err != nil {
				return "", err
    }

		return response, nil
}
