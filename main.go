package main

import (
    "strings"
    "fmt"
    "os"
)

var _ = fmt.Fprint

func main() {
    for {
	cmd , err := GetUserCommand()
	if err != nil {
	    fmt.Println("error getting user command")
	}
	tokens := strings.Split(cmd, " ")
	switch tokens[0] {
	case "cd":
	    err := os.Chdir(tokens[1])
	    if err != nil {
		fmt.Println("cd: " + tokens[1] + ": No such file or directory")
	    }
	case "exit":
	    os.Exit(0)
	case "echo":
	    fmt.Println(strings.Join(tokens[1:]," "))
	case "type":
	    fmt.Println(FindCmd(tokens[1]))
	case "pwd":
	    currentWorkingDirectory,_ := os.Getwd();
	    fmt.Println(currentWorkingDirectory)
	default:
	   fmt.Println(SearchExec(tokens))
	}
    }

}
