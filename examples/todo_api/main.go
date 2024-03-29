package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID 			string 	`json:"id"`
	Item 		string 	`json:"item"`
	Completed 	bool	`json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Write some code", Completed: false},
	{ID: "2", Item: "Review PRs", Completed: false},
	{ID: "3", Item: "Have a break", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(id string) (*todo, error){
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func addTodo(context *gin.Context){
	var todo todo
	if err := context.BindJSON(&todo); err != nil {
		return
	}

	todos = append(todos, todo)
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func main(){
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9000")
}