package main

import (
    "fmt"
    "os"
    "golang.org/x/term"
)
func HandleCompletion(cmd string) string{
    
}

func ReadInput()(string,error){
    fd := int(os.Stdin.Fd())
    old,_ := term.MakeRaw(fd)
    defer term.Restore(fd,old)
    buf := make([]byte, 1)
    cmd := ""
    for{
	_,err := os.Stdin.Read(buf)
	if err != nil {
	    return "",err
	} else if buf[0] == 9{
	    HandleCompletion(cmd)
	} else if buf[0] == 127{
	    if (len(command) > 0){
		command = command[:len(command)-1]
	    }
	}else if buf[0] == 3{
	    command = command + "^C"
	    break
	}else {
	    command = command + string(buf)
	}
	fmt.Print("\x1b[2K\x1b[1G" + command)
    }
    return cmd,nil
}


