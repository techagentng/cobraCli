package todo

import "fmt"

func (todos *Todos) Done(value int) bool {
	if len(todos.Tasks) < value {
		fmt.Println("Please enter a valid no")
		return false
	} else {
		todos.Tasks[value-1].Done = true
		todos.addToFile()
		return true
	}
}
