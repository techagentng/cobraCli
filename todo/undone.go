package todo

import "fmt"

func (todo *Todos) Undone(value int) bool {
	if len(todo.Tasks) < value {
		fmt.Println("Please enter a valid no")
		return false
	} else {
		todo.Tasks[value-1].Done = false
		todo.addToFile()
		return true
	}
}
