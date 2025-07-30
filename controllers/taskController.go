package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-app/models"

	"github.com/go-chi/chi/v5"
)

// получить все задачи
func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Tasks)
}

// получить задачу по id
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	//получаем из URL HTTP запроса id
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	//ищем задачу с заданным id
	for i, t := range models.Tasks {
		if t.ID == id {
			//запись в JSON и ответ фронтенду
			json.NewEncoder(w).Encode(models.Tasks[i])
			return
		}
	}
	//возврат ошибки 404
	http.NotFound(w, r)
}

// создание задачи
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	//декодируем из JSON и пишем по адресу t
	json.NewDecoder(r.Body).Decode(&t)
	t.ID = models.NextID
	models.NextID++
	models.Tasks = append(models.Tasks, t)
	//сохранение в файл
	models.SaveData()
	json.NewEncoder(w).Encode(t)
}

// удаление задачи
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	for i, t := range models.Tasks {
		if t.ID == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			models.SaveData()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

// инвертировать статус задачи
func ToggleTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	for i, t := range models.Tasks {
		if t.ID == id {
			models.Tasks[i].Completed = !models.Tasks[i].Completed
			models.SaveData()
			json.NewEncoder(w).Encode(models.Tasks[i])
			return
		}
	}
	http.NotFound(w, r)
}

// изменить задачу
func EditTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	for i, t := range models.Tasks {
		if t.ID == id {
			var t_new models.Task
			json.NewDecoder(r.Body).Decode(&t_new)
			models.Tasks[i].Title = t_new.Title
			models.Tasks[i].Description = t_new.Description
			models.SaveData()
			json.NewEncoder(w).Encode(models.Tasks[i])
			return
		}
	}
	http.NotFound(w, r)
}
