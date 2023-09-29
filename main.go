package main

/*
the todo is a command line application that allows you to add, delete, and mark items as complete.
todo -h            # help
todo -v            # version
todo -l            # list all items (un-completed)
todo -c            # list completed items
todo -a"Buy Milk"  # add new item
todo -m TODO-Id    # mark as complete
todo -d TODO-Id    # delete item
*/

import (
	"flag"
	"fmt"
)

var (
	// Version represents the git tag of a particular release.
	Version = "v0.0.0"

	// GitCommit represents git commit hash of a particular release.
	GitCommit = "dev"
)

func main() {
	filename := ".todo.json"

	todoList := TodoList{}

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "NAME:\n	todo - simple todo application that allows you to add, delete, and mark items as complete.")
		fmt.Fprintln(flag.CommandLine.Output(), "USAGE:\n	todo [OPTIONS]")
		fmt.Fprintln(flag.CommandLine.Output(), "OPTIONS:")
		flag.PrintDefaults()
	}
	// Parsing the arguments ...
	//h := flag.Bool("h", false, "help")
	v := flag.Bool("v", false, "version")
	l := flag.Bool("l", false, "list all items (un-completed)")
	c := flag.Bool("c", false, "list completed items")
	a := flag.String("a", "", "add new item")
	m := flag.Int("m", -1, "mark as complete")
	um := flag.Int("um", -1, "mark as incomplete (unmark)")
	d := flag.Int("d", -1, "delete item")

	flag.Parse()

	readFromFile(todoList, filename)

	if *v {
		printVersion()
	} else if *l {
		todoList.printInComplete()
	} else if *c {
		todoList.printComplete()
	} else if *a != "" {
		todoList.addItem(*a)
	} else if *m != -1 {
		todoList.markItem(*m)
	} else if *um != -1 {
		todoList.unMarkItem(*um)
	} else if *d != -1 {
		todoList.deleteItem(*d)
	}

	// Saving the todo list to a file

	todoList.writeToFile(filename)
}

func printVersion() {
	v := Version + "-" + GitCommit
	fmt.Println("todo ", v, " (by Muhammed Can Küçükaslan https://github.com/Kucukaslan)\nYou're currently using the version ", v)
	fmt.Println("There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.")
}
