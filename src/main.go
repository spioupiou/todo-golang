package main

import (
	Model "golang-todo/src/model"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", Model.GetTodos)

	router.Run("localhost:8080")
}