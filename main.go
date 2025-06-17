package main

import (
    "strings"
    "fmt"
    "os"
)

var _ = fmt.Fprint

func main() {
    for {
	input , err := GetUserCommand()
	if err != nil {
	    fmt.Println("error getting user command")
	}
	tokens := strings.Split(input, " ")
	switch tokens[0] {
	case "exit":
	    os.Exit(0)
	case "echo":
	    fmt.Println(strings.Join(tokens[1:]," "))
	case "type":
	    fmt.Println(FindCmd(tokens[1]))
	default:
	   fmt.Println(SearchExec(tokens))
	}
    }

}

