package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := ReadFile("./testData/2024-06-13.md")
	todos := MakeTodos(data)

	fmt.Printf("Todos: \n%v", todos)
}

type Todo struct {
	task string
	info []string
	done bool
}

func NewTodo(todo []string) Todo {
	if len(todo) == 0 {
		return Todo{}
	}

	var info []string
	info = nil

	if len(todo) > 1 {
		for i := 1; i < len(todo); i++ {
			info = append(info, todo[i])
		}
	}

	return Todo{
		task: todo[0],
		info: info,
		done: false,
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) []string {
	myfile, err := os.Open(filename) //open the file
	defer myfile.Close()
	check(err)

	scanner := bufio.NewScanner(myfile) //scan the contents of a file and print line by line
	var dat []string

	for scanner.Scan() {
		dat = append(dat, scanner.Text())
	}

	return dat
}

func MakeTodos(data []string) []Todo {
	var todos []Todo
	todo := []string{}
	copyTodo := false

	for _, line := range data {
		if strings.HasPrefix(line, "- [ ] ") {
			if len(todo) > 0 {
				todos = append(todos, NewTodo(todo))
			}
			todo = []string{}
			copyTodo = true
		}
		if copyTodo {
			todo = append(todo, line)
		}
	}

	todos = append(todos, NewTodo(todo))

	return todos
}
