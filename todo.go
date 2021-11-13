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
	"io/ioutil"
	"os"
	"strings"
)

type TodoList map[int]*item

func main() {
	filename := ".todo.json"

	todoList := TodoList{}
	readFromFile(todoList, filename)

	// Parsing the arguments ...
	h := flag.Bool("h", false, "help")
	v := flag.Bool("v", false, "version")
	l := flag.Bool("l", false, "list all items (un-completed)")
	c := flag.Bool("c", false, "list completed items")
	a := flag.String("a", "", "add new item")
	m := flag.Int("m", -1, "mark as complete")
	d := flag.Int("d", -1, "delete item")
	flag.Parse()

	if *h {
		printHelp()
	} else if *v {
		printVersion()
	} else if *l {
		printInComplete(todoList)
	} else if *c {
		printComplete(todoList)
	} else if *a != "" {
		addItem(&todoList, *a)
	} else if *m != -1 {
		markItem(&todoList, *m)
	} else if *d != -1 {
		deleteItem(&todoList, *d)
	} else {
		//fmt.Println("No arguments given")
		printHelp()
	}

	// Saving the todo list to a file

	todoList.writeToFile(filename)
}

func printHelp() {
	fmt.Println("-----------------\nHelp for using the todo cli app")
	fmt.Println("-----------------")
	fmt.Println("todo -h            # help")
	fmt.Println("todo -v            # version")
	fmt.Println("todo -l            # list all items (un-completed)")
	fmt.Println("todo -c            # list completed items")
	fmt.Println("todo -a \"Buy Milk\"  # add new item")
	fmt.Println("todo -m TODO-ID    # mark as complete")
	fmt.Println("todo -d TODO-ID    # delete item")
}

func printVersion() {
	v := "0.2.0"
	date := "2021.11.14"
	fmt.Println("todo ", v, "\nYou're currently using the version ", v, " released on", date)
	fmt.Println("There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.")
}

func printInComplete(todoList TodoList) {
	for _, item := range todoList {
		if !item.Status {
			fmt.Println(item.toString())
		}
	}
	//printMap(todoList)
}
func printComplete(todoList TodoList) {
	for _, item := range todoList {
		if item.Status {
			fmt.Println(item.toString())
		}
	}
	//printMap(todoList)
}

// TODO length is not correct, when an item is deleted
// then the (*todoList)[length]length may be already exist
func addItem(todoList *TodoList, Title string) int {
	length := len(*todoList)
	i := item{length, Title, "today", false}
	(*todoList)[length] = &i
	return length
}

func deleteItem(todoList *TodoList, Id int) {
	delete(*todoList, Id)
}

func markItem(tdmap *TodoList, Id int) {
	(*tdmap)[Id].Status = true
}

func (td TodoList) writeToFile(filename string) error {
	str := "["
	for _, item := range td {
		str += item.toJSON() + ","
	}
	// we need to remove the last comma
	if len(str) > 1 {
		str = str[:len(str)-1]
	}
	str += "]"
	return ioutil.WriteFile(filename, []byte(str), 0666)

}

func readFromFile(td TodoList, filename string) TodoList {

	// Reading the file
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("Cannot read ", filename)
		return td
	}
	str := string(byteSlice)
	//fmt.Println("\n", str)
	str = strings.Trim(str, "[ ]")
	//fmt.Println("\n", str)
	str = strings.Replace(str, "},{", "}|{", -1)
	//fmt.Println("\n", str)
	arr := strings.Split(str, "|")
	//fmt.Println("\n", arr)

	for _, i := range arr {
		//fmt.Println("\n\n", xx, " ", i)
		it := parseJSON(i)
		if it != nil {
			td[it.Id] = it
		}
	}
	return td
}

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
