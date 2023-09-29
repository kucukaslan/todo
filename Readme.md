[![Go Report](https://goreportcard.com/badge/github.com/kucukaslan/todo)](https://goreportcard.com/report/github.com/kucukaslan/todo) ![Github Actions Status](https://github.com/kucukaslan/todo/actions/workflows/go.yml/badge.svg)

# `todo` cli
the todo is a command line application that allows you to add, delete, and mark items as complete.

## Features and Flags 
- The data is stored in a json file to persist between sessions.   
- Each item is associated with the date added to the list. The `-l` flag also display how much time has been passed since the item added to the list


The following is the list of the flags and their functionalities.
|   command (with flag)	|  description 	| 
|:--------------------|:----------	|
|`todo -h`            | help  |
|`todo -v`            | version  |
|`todo -l`            | list all items (un-completed)   |
|`todo -c`            | list completed items  |
|`todo -a "Buy Milk"` | add new item  |
|`todo -m TODO-ID   ` | mark as complete  |
|`todo -um TODO-ID  ` | mark as incomplete (unmark)  |
|`todo -d TODO-ID   ` | delete item  |


## Installation
### Windows
Download the source code.  
Open the directory.  
Install the application with `go install todo`.  

The data is stored in the `%AppData%/todo/.todo.json`, 
which is obtained by [UserConfigDir](https://pkg.go.dev/os#UserConfigDir)

### GNU/Linux
Download the source code.  
Open the directory.  
Install the application with `go install todo`.

If the command is not installed try building executable using the `go build .` command and copy the executable to `/usr/local/bin/` directory with the `cp todo /usr/local/bin/` command.

The data is stored in the `$XDG_CONFIG_HOME/todo/.todo.json`, 
which is, also, obtained by [UserConfigDir](https://pkg.go.dev/os#UserConfigDir)
