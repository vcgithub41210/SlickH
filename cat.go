package main

import (
    "fmt"
    "os"
)


func Cat(args []string,target string) {
    var o *os.File
    if target == "" {
	o = os.Stdout
    } else { 
	var err error
	o, err = os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
	    panic(err)
	}
    }
    for _,file := range args {
	content, err := os.ReadFile(file)
	if err != nil {
	    os.Stdout.Write([]byte(fmt.Sprintf("cat: %s: No such file or directory\n",file)))
	} else {
	    o.Write(content)
	}
    }
}
