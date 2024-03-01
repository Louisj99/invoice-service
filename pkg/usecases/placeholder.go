package usecases

import (
	"context"
	"github.com/gin-gonic/gin"
)

type PlaceholderRequest struct {
	Placeholder string `json:"placeholder"`
}
type PlaceholderResponse struct {
	Message string `json:"message"`
}

type PlaceholderInterfacePostgres interface {
	Placeholder(ctx context.Context, placeholder string) error
}

func Placeholder(placeholderPostgres PlaceholderInterfacePostgres, placeholder string) gin.HandlerFunc {
	return func(c *gin.Context) {

		err := placeholderPostgres.Placeholder(c, placeholder)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Placeholder"})
	}
}
