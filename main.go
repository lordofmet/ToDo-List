package main

import (
	"fmt"
	"net/http"
	"todo-app/controllers"
	"todo-app/models"

	"github.com/go-chi/chi/v5"
)

func main() {

	err := models.LoadData()
	if err != nil {
		fmt.Println("Error loading tasks")
	}

	r := chi.NewRouter()

	r.Get("/api/tasks", controllers.GetTasks)
	r.Post("/api/tasks", controllers.CreateTask)
	r.Delete("/api/tasks/{id}", controllers.DeleteTask)
	r.Patch("/api/tasks/{id}/toggle", controllers.ToggleTask)
	r.Get("/api/tasks/{id}", controllers.GetTaskById)
	r.Put("/api/tasks/{id}", controllers.EditTask)

	//любой файл из папки доступен на локалхосте
	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("view"))))

	http.ListenAndServe(":8080", r)
}
