package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckHandler is a handler that returns the status of the service.
func CheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
