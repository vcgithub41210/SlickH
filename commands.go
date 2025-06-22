package main

import (
    "os/exec"
    "strings"
    "os"
    "bufio"
    "fmt"
)

var builtins = []string{"type", "echo", "exit", "pwd"}

func ParseCommand(command string) (string,[]string,string){
    // set of variables required for parsing
    var (
	arguements []string
	curr string
	isEscape bool
	inSingleQuote bool
	inDoubleQuote bool
	isHomeReference bool
    )
    redirect_idx := -1
    idx := 0
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
	} else if ch == '>' && (curr == "" || curr == "1") && !inDoubleQuote && !inSingleQuote{
	    redirect_idx = idx
	    curr = ""
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
		idx++
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
	idx++
    }
    if len(arguements) == 0 {
	return "",arguements,""
    }
    if redirect_idx != -1{
	if redirect_idx == idx{
	    return arguements[0],arguements[1:idx],""
	}
	return arguements[0] , arguements[1:redirect_idx], arguements[redirect_idx]
    }
    return arguements[0],arguements[1:],""
}

func ChangeDirectory(path string) {
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

func SearchExec(cmd string,tokens []string) (string,error){
    for _,path := range( strings.Split(os.Getenv("PATH"),":")){
	file := path + "/" + cmd
	if _, err := os.Stat(file); err == nil {
	    command := exec.Command(file, tokens...)
	    output, err := command.CombinedOutput()
	    if err != nil {
		return "",err
	    }
	    file = string(output)
	    return file, nil
	}
    }
    return fmt.Sprintf("%s: command not found",cmd),nil
}

func WriteToTarget(output string,file string) error{
    var (
	f *os.File
	err error
    )
    if file == "" {
	os.Stdout.Write([]byte(output))
	return nil
    }
    err = nil
    f, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    _,err = f.Write([]byte(output))
    err = f.Close()
    return err
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
