package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// A Generic function that responds to the client request.
func RespondClientError(c *gin.Context, r *http.Request, status int, reason string) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, gin.H{
		"error": reason,
	})

}
