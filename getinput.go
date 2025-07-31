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
		fmt.Println(c)
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

		switch buf[0] {
		case 9: // Tab key
		fmt.Println("working")
			completed := HandleCompletion(cmd)
			if completed != "" {
				cmd = completed
			}
		case 127: // Backspace
			if len(cmd) > 0 {
				cmd = cmd[:len(cmd)-1]
			}
		case 3: // Ctrl+C
			fmt.Println("^C")
			return "", fmt.Errorf("interrupted")
		case 13: // Enter
			fmt.Println()
			return cmd, nil
		default:
			cmd += string(buf[0])
		}

		// Clear and redraw line
		fmt.Print("\x1b[2K\r" + cmd)
	}
}

