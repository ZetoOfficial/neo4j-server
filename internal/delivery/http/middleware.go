package http

import (
	"net/http"

	"github.com/ZetoOfficial/neo4j-server/internal/config"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		cfg := config.LoadConfig()
		if token != cfg.AuthToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}
