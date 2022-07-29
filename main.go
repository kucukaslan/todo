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

func main() {
	filename := ".todo.json"

	todoList := TodoList{}

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
		addItem(&todoList, *a)
	} else if *m != -1 {
		markItem(&todoList, *m)
	} else if *um != -1 {
		unMarkItem(&todoList, *um)
	} else if *d != -1 {
		deleteItem(&todoList, *d)
	}

	// Saving the todo list to a file

	todoList.writeToFile(filename)
}

/*
func printHelp() {
	fmt.Println("-----------------")
	fmt.Println("the todo is a command line application that allows you to add, delete, and mark items as complete.")
	fmt.Println("The list of flags in the todo cli app and their explanation:")
	fmt.Println("-----------------")
	fmt.Println("todo -h            # help")
	fmt.Println("todo -v            # version")
	fmt.Println("todo -l            # list all items (un-completed)")
	fmt.Println("todo -c            # list completed items")
	fmt.Println("todo -a \"Buy Milk\" # add new item")
	fmt.Println("todo -m TODO-ID    # mark as complete")
	fmt.Println("todo -um TODO-ID   # mark as incomplete (unmark)")
	fmt.Println("todo -d TODO-ID    # delete item")
}
*/
func printVersion() {
	v := "0.3.0"
	date := "2021.11.14"
	fmt.Println("todo ", v, "(by Muhammed Can Küçükaslan https://github.com/MuhammedCanKucukaslan)\nYou're currently using the version ", v, " released on", date)
	fmt.Println("There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.")
}

func addItem(td *TodoList, Title string) int {
	return td.addItem(Title)

}

func deleteItem(td *TodoList, Id int) {
	td.deleteItem(Id)
}

func markItem(td *TodoList, Id int) {
	td.markItem(Id)
}
func unMarkItem(td *TodoList, Id int) {
	td.unMarkItem(Id)
}

/*
func debugPrintArgs() {
	fmt.Println("-----------------\nList of the Args")
	fmt.Println("-----------------")
	for i, a := range os.Args {
		fmt.Println(i, " ", a)
	}
	fmt.Println("-----------------")
}

func printMap(m TodoList) {
	for i, item := range m {
		fmt.Println(i, " ", item.toString())
	}
}
*/
