package main

import (
	"fmt"
	"time"
)

type todo struct {
	ID     			int64  		`json:"id"`
	title  			string  	`json:"title"`
	Status			bool 			`json:"status"`
	CreatedAt		time.Time	`json:"created_at"` 
}

var sampleTodos = []todo {
	{ID: 1, title: "learn go", Status: false, CreatedAt: time.Now()},
	{ID: 2, title: "devops mocking test review", Status: false, CreatedAt: time.Now()},
}

func main() {
	fmt.Print(sampleTodos[0].title)
}