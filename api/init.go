package api

import "github.com/gin-gonic/gin"

// Router offers API server handlers
var Router = gin.Default()

// Init 初始化路由
func Init() {
	Router.Use(RequestIDMiddleware)
}
