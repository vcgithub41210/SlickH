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
	command, args, target, redir_type := ParseCommand(user_input) 
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
	    if redir_type == 0{
		WriteToTarget(fmt.Sprintf("%s\n",strings.Join(args," ")),target)
	    } else {
		WriteToTarget(fmt.Sprintf("%s\n",strings.Join(args," ")),"")
		WriteToTarget("",target)
	    }
	case "cat":
	    Cat(args,target,redir_type)
	case "type":
	    fmt.Println(FindCmd(args[0]))
	case "pwd":
	    currentWorkingDirectory,_ := os.Getwd();
	    fmt.Println(currentWorkingDirectory)
	default:
	    err := Execute(command,args,target,redir_type)
	    if err != nil {
		if redir_type == 0{
		    WriteToTarget(err.Error(),"")
		} else {
		    WriteToTarget(err.Error(),target)
		}
	    }
	}
    }

}

