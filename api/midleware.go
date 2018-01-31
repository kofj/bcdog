package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestIDMiddleware 给请求标注 ID
func RequestIDMiddleware(c *gin.Context) {
	c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
	c.Next()
}
