package main

/*
the todo is a command line application that allows you to add, delete, and mark items as complete.
todo -h            # help
todo -v            # version
todo -l            # list all items (un-completed)
todo -c            # list completed items
todo -a"Buy Milk"  # add new item
todo -m TODO-ID    # mark as complete
todo -d TODO-ID    # delete item
*/

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	h := flag.Bool("h", false, "help")
	v := flag.Bool("v", false, "version")
	l := flag.Bool("l", false, "list all items (un-completed)")
	c := flag.Bool("c", false, "list completed items")
	a := flag.String("a", "", "add new item")
	m := flag.String("m", "", "mark as complete")
	d := flag.String("d", "", "delete item")
	flag.Parse()

	if *h {
		printHelp()
	} else if *v {
		printVersion()
	} else if *l {
		printlist()
	}

	fmt.Println(*c)
	fmt.Println(*a)
	fmt.Println(*m)
	fmt.Println(*d)
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
	fmt.Println("todo 0.0.2\nYou're currently using the version 0.02 released on 2021.11.13")
	fmt.Println("There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.")
}

func printlist() {
	fmt.Println("nList of the todo cli app")
}

func debugPrintArgs() {
	fmt.Println("-----------------\nList of the Args")
	fmt.Println("-----------------")
	for i, a := range os.Args {
		fmt.Println(i, " ", a)
	}
	fmt.Println("-----------------")
}
