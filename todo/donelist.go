package todo

import (
	"os"
	"text/template"
)

func (todo Todos) DoneList() {
	tempString := "{{with .Tasks}}{{range $idx, $list := .}}{{if $list.Done}}"
	tempString += "{{increment $idx}}. {{$list.Content}} - ({{date $list.Date}})\n"
	//tempString += "{{increment $idx}}. {{$list.Content}}\n"
	tempString += "{{end}}{{end}}{{end}}"

	t := template.Must(template.New("UndoneList").Funcs(funcMap).Parse(tempString))
	err := t.Execute(os.Stdout, todo)
	if err != nil {
		if err != nil {
			err.Error()
		}
	}
}
