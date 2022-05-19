package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"todocli/todo"
)

var Lst = todo.Todos{}

func init() {
	LoadTodoList()
}

func LoadTodoList() {
	file, err := ioutil.ReadFile("file.csv")
	if err != nil {
		log.Println(err)
	}
	_ = json.Unmarshal(file, &Lst)
}
