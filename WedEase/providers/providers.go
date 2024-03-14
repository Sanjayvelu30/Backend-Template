package providers

import "github.com/gin-gonic/gin"

type Middelware interface {
	AuthenticationMiddleware() gin.HandlerFunc
}
