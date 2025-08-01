package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

var commands = []string{"exit", "echo"}


func HandleCompletion(cmd string) string {
	for _, c := range commands {
		if strings.HasPrefix(c, cmd) {

			return c
		}
	}
	return ""
}

func ReadInput() (string, error) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return "", err
	}
	defer term.Restore(fd, oldState)

	buf := make([]byte, 1)
	cmd := ""

	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			return "", err
		}
		if buf[0] == 9{
			e := HandleCompletion(cmd)
			if e != ""{
				cmd = e
			}
		}else if buf[0] == 127 {
			if len(cmd) > 0 {
				cmd = cmd[:len(cmd)-1]
			}
		}else if buf[0] == 13 {
			fmt.Println("\r")
			return cmd, nil
		}else {
			cmd += string(buf[0])
		}

		// Clear and redraw line
		fmt.Print("\x1b[2K\r$ " + cmd)
	}
}

