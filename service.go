package main

import (
	"errors"
	"fmt"
	"time"
)

func add(description string) error {

	time := time.Now()
	formatTime := time.Format("2006-01-02 15:04:05")
	task, err := loadTask()
	if err != nil {
		return err
	}

	newId := 1
	if len(task) > 0 {
		last := task[len(task)-1]
		newId = last.ID + 1
	}
	newTodo := Todo{
		ID:          newId,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   formatTime,
		UpdateAt:    formatTime,
	}

	task = append(task, newTodo)
	fmt.Println("Succesfully added , ID :", newTodo.ID)
	return writeFile(task)

}

// Get Data By Status TODO
func getDataByStatus(status Status) ([]Todo, error) {
	getData, _ := loadTask()
	for i := range getData {
		if getData[i].Status == status {
			return getData, nil
		}
	}

	return nil, errors.New("Data with Status Todo Is Empty")
}

// Find Data By ID
func findDataById(todos []Todo, id int) (int, error) { // Find Data By ID
	for i := range todos {
		if todos[i].ID == id {
			return i, nil
		}
	}

	return -1, errors.New("ID tidak ada")
}

// Delete Data With ID
func delete(id int) error {
	todo, err := loadTask()

	if err != nil {
		return err
	}

	index, err := findDataById(todo, id)

	if err != nil {
		return err
	}

	todo = append(todo[:index], todo[index+1:]...)
	return writeFile(todo)
}

func changeStatus(id int, status Status) error {

	todo, err := loadTask()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i := range todo {
		if todo[i].ID == id {
			todo[i].Status = status
			todo[i].UpdateAt = time.Now().Format("2006-01-02 15:04:05")
			return writeFile(todo)
		}
	}

	return errors.New("ID NOT FOUND !!")
}

func updateDescription(id int, description string) error {
	todos, err := loadTask()
	if err != nil {
		return err
	}

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Description = description
			todos[i].UpdateAt = time.Now().Format("2006-01-02 15:04:05")
			return writeFile(todos)
		}
	}

	return errors.New("ID NOT FOUND !!")
}
