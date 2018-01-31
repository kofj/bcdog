package api

import "github.com/gin-gonic/gin"

// Engine offers API server handlers
var Engine = gin.Default()

// Init 初始化路由
func Init() {
	Engine.Use(RequestIDMiddleware)
}
