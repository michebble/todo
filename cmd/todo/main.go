package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/michebble/todo"
)

// Hardcoding the file name
const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	// Use the Get method to read ToDo items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	//For no extra arguments, print the list
	case len(os.Args) == 1:
		// List current ToDo items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	// Concatenate all provided arguments with a space
	// add to the list as an item
	default:
		item := strings.Join(os.Args[1:], " ")

		l.Add(item)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
