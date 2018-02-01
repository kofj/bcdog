package api

import "github.com/gin-gonic/gin"

// Router offers API server handlers
var Router = gin.Default()

// Init 初始化路由
func Init() {
	Router.Use(RequestIDMiddleware)
}

// BinFsHandler 返回文件 handler
func BinFsHandler(name string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var bytes, err = Asset(name)
		if err != nil {
			c.String(404, "Page not found")
			return
		}
		c.Header("Content-Type", "text/html")
		c.String(200, string(bytes))
	}
}
