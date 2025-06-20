package main

import (
    "os/exec"
    "strings"
    "os"
    "bufio"
    "fmt"
)

var builtins = []string{"type", "echo", "exit", "pwd"}

func ParseCommand(command string) []string{
    // set of variables required for parsing
    var (
	arguements []string
	curr string
	isEscape bool
	inSingleQuote bool
	inDoubleQuote bool
	isHomeReference bool
    )
    for i:= 0; i < len(command);i++{
	ch := command[i]
	if isHomeReference {
	    if ch == ' ' || ch == '/' {
		curr += os.Getenv("HOME")
	    } else {
		curr += "~"
	    }
	    isHomeReference = false
	}
	if isEscape{
	    if inDoubleQuote && !(ch == '$' || ch == '\\' || ch == '"'){
		curr += "\\"
	    }
	    curr = curr + string(ch)
	    isEscape = false
	} else if ch == '~' && !inDoubleQuote && !inSingleQuote && curr == "" {
	    isHomeReference = true
	} else if ch == '\\' && !inSingleQuote{
	    isEscape = true
	} else if ch == '\'' && !inDoubleQuote {
	    inSingleQuote = !inSingleQuote
	} else if ch == '"' && !inSingleQuote {
	    inDoubleQuote = !inDoubleQuote
	} else if ch == ' ' && !inSingleQuote && !inDoubleQuote{
	    if curr != ""{
		arguements = append(arguements,curr)
		curr = ""
	    }
	} else {
	    curr = curr + string(ch)
	}
    }
    if curr != "" || isHomeReference{
	if isHomeReference {
	    curr = os.Getenv("HOME")
	}
	arguements = append(arguements,curr)
    }
    return arguements
}
func ChangeDirectory(path string) {
    if path[0] == '~'{
	path = os.Getenv("HOME") + path[1:]
    }
    err := os.Chdir(path)
    if err != nil{
	fmt.Println("cd: " +  path + ": No such file or directory")
    }
}

func GetUserCommand() (string, error){
    fmt.Fprint(os.Stdout, "$ ")
    command, err := bufio.NewReader(os.Stdin).ReadString('\n')
    if err != nil {
	return "", err
    }
    command = command[:len(command)-1]
    return command, nil
}

func SearchExec(tokens []string) string{
    name := tokens[0]
    for _,path := range( strings.Split(os.Getenv("PATH"),":")){
	file := path + "/" + name
	if _, err := os.Stat(file); err == nil {
	    cmd := exec.Command(file, tokens[1:]...)
	    output, err := cmd.CombinedOutput()
	    if err != nil {
		panic(err)
	    }
	    file = string(output)
	    return file[:len(file)-1]
	}
    }
    return fmt.Sprintf("%s: command not found",name)
}

func FindCmd(cmd string) string {
    for _,builtin := range builtins {
	if builtin == cmd {
	    return fmt.Sprintf("%s is a shell builtin",cmd)
	}
    }
    for _,path := range (strings.Split(os.Getenv("PATH"), ":")) {
	file := path + "/" + cmd
	if _, err := os.Stat(file); err == nil{
	    return fmt.Sprintf("%s is %s",cmd,file)
	}
    }
    return fmt.Sprintf("%s: not found",cmd)
}
