package todo

import (
	"fmt"
	"os"
	"text/template"
)

func (todos Todos) All() {

	if len(todos.Tasks) == 0 {
		fmt.Println("Empty Task, Please Add Content")
	}


	todoTmpl := "{{with .Tasks}}{{range $idx, $todo := .}}"

	todoTmpl += "{{if $todo.Done}}{{else}}"

	todoTmpl += "{{increment $idx }}. {{ $todo.Content}} - ({{date $todo.Date}})\n"

	todoTmpl += "{{end}}{{end}}{{end}}"

	tmpl := template.Must(template.New("todo").Funcs(funcMap).Parse(todoTmpl))

	_ = tmpl.Execute(os.Stdout, todos)

}
