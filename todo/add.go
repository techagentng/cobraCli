package todo

import "time"

func (todos *Todos) AddTask(content string) {
	todo := Task{Content: content, Date: time.Now().Unix()}
	todos.Tasks = append(todos.Tasks, todo)

	todos.addToFile()
}
