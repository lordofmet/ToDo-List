package models

//структура задачи - id, название, описание, статус
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Completed   bool   `json:"completed"`
}

// список задач, вместо БД
var Tasks = []Task{}

// первый индекс, по которому возможна запись
var NextID = 1
