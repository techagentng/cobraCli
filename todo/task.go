package todo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"text/template"
	"time"
)

const FILENAME = "file.csv"

type Task struct {
	Content string	`json:"content"`
	Done    bool	`json:"done"`
	Date    int64	`json:"date"`
}

type Todos struct {
	Tasks	 []Task `json:"tasks"`
}

var funcMap = template.FuncMap{
	"increment": func(i int) int {
		return i + 1
	},
	"date": func(timeInt int64) string {
		return time.Unix(timeInt, 0).Format("Mon, 02 Jan 2006 15:04:05")
	},
}

func (todos Todos) addToFile() {
	res, _ := json.MarshalIndent(todos, "", "\t")

	err := ioutil.WriteFile(FILENAME, res, 0666)
	if err != nil {
		log.Println(err)
	}
}