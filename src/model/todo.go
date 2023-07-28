package model

import (
	// "net/http"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     			int64  		`json:"id"`
	Title  			string  	`json:"title"`
	Status			bool 			`json:"status"`
	CreatedAt		time.Time	`json:"created_at"` 
}

var sampleTodos = []Todo {
	{ID: 1, Title: "learn go", Status: false, CreatedAt: time.Now()},
	{ID: 2, Title: "devops mocking test review", Status: false, CreatedAt: time.Now()},
}

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sampleTodos)
}