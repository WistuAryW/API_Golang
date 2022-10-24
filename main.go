package main

import(
	"github.com/gin-gonic/gin"
	"errors"
	"net/http"
)

type todo struct {
	Id         	 string  `json:"id"`
	Name         string  `json:"name"`
	Price        int     `json:"price"`
	Quantity     int     `json:"quantity"`
	Status       bool    `json:"status"`
}

var todos = []todo{
	{Id: "1", Name: "Minyak goreng", Price: 12000, Quantity: 20, Status: false},
	{Id: "2", Name: "Gula pasir", Price: 14000, Quantity: 30, Status: false},
	{Id: "3", Name: "Daging", Price: 78000, Quantity: 10, Status: false},
}

func getTodos(context *gin.Context){
    context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context){
    var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil{
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}


func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	
	context.IndentedJSON(http.StatusOK, todo)
}

func statusTodo(context *gin.Context){
    id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Status = !todo.Status
	context.IndentedJSON(http.StatusOK, todo)
}


func getTodoById(id string) (*todo, error){
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}
func main()  {
	router  := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", statusTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}




