package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RedHoodJT1988/todo-cli"
)

// Hardcoding the file name
const todoFileName = ".todo.json"

func main() {
	// Testing out the Usage function
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for a convienent way to create tasks of ToDo Items|n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2021\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}
	// Parsing command line flags
	task := flag.String("task", "", "Task to be included in the ToDo List")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	// Define an item list
	l := &todo.List{}
	
	// Use the Get command to read ToDo items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	case *list:
		// List current to do items
		fmt.Print(l)
	case *complete > 0:
		// Complete the Given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		
	case *task != "":
		// Add the task
		l.Add(*task)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}