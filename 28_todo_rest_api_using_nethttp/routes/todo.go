package routes

import (
	"net/http"
	"todo-rest-api/controllers"
)

func RegisterTodoRoutes(){
	http.HandleFunc("GET /todos",controllers.GetTodos)
	http.HandleFunc("GET /todos/{id}",controllers.GetTodoById)
	http.HandleFunc("POST /todos",controllers.CreateTodo)
	http.HandleFunc("PATCH /todos/{id}",controllers.UpdateTodo)
	http.HandleFunc("DELETE /todos/{id}",controllers.DeleteTodo)
}