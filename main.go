package main

import (
    "strings"
    "fmt"
    "os"
)

var _ = fmt.Fprint

func main() {
    for {
	user_input , err := GetUserCommand()
	if err != nil {
	    fmt.Println("error getting user command")
	}
	command, args, outf, errf, mods := ParseCommand(user_input) 
	switch command {
	case "cd":
	    if len(args) > 1{
		fmt.Println("cd: too many arguements")
	    } else if len(args) == 0{
		ChangeDirectory(os.Getenv("HOME"))
	    } else { 
		ChangeDirectory(args[0])
	    }
	case "exit":
	    os.Exit(0)
	case "echo":
		WriteToTarget(fmt.Sprintf("%s\n",strings.Join(args," ")),outf,mods&2)
		WriteToTarget("",errf,mods&1)
	case "type":
	    if len(args) != 0 {
		fmt.Println(FindCmd(args[0]))
	    }
	case "pwd":
	    currentWorkingDirectory,_ := os.Getwd();
	    fmt.Println(currentWorkingDirectory)
	default:
	    err := Execute(command,args,outf,errf,mods)
	    if err != nil {
		WriteToTarget(err.Error(),errf,mods & 1)
	    }
	}
    }

}

