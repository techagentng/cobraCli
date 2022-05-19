package todo

func (todos *Todos) Cleanup() {
	for idx:=0; idx < len(todos.Tasks); idx++{
		if todos.Tasks[idx].Done && idx < len(todos.Tasks)-1 {
			todos.Tasks = append(todos.Tasks[:idx], todos.Tasks[idx+1:]...)
			idx -= 1
		} else if todos.Tasks[idx].Done && idx == len(todos.Tasks)-1 {
			todos.Tasks = todos.Tasks[:idx]
		}
	}
	todos.addToFile()
}
