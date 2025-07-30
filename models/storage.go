package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// относительный путь к файлу
const path = "data.json"

// загрузка списка задач из файла
func LoadData() error {
	//попытка открыть файл
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	//отложенное закрытие
	defer file.Close()

	//декодируем в уже существующую в task.go переменную
	err = json.NewDecoder(file).Decode(&Tasks)
	if err != nil {
		fmt.Println("JSON Decoding error")
		return err
	}

	//новый счетчик, т.к. в прошлом сеансе могли быть удалены задачи
	maxID := 0
	for _, t := range Tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	//обновляем текущий индекс для записи
	NextID = maxID + 1
	return nil
}

// сохранение списка задач в файл
func SaveData() error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Couldn't create a save file")
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(Tasks)
}
