package main

import (
	"github.com/DSurfer/goredis/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/get_key", handlers.UsersHandlerGetKey)
	router.POST("/set_key", handlers.UsersHandlerSetKey)
	router.POST("/del_key", handlers.UsersHandlerDelKey)
	router.Run("localhost:8089")
}
