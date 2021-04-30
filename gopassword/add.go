package gopassword

import (
	"fmt"
	"os"
)

func addCmd(splt []string) {
	if len(splt) != 4 {
		displayCommands()
	} else {
		service := splt[1]
		username := splt[2]
		password := splt[3]

		line := service + " " + username + " " + password + " \n"
		wline := []byte(line)
		f, err := os.OpenFile("rct_pass_file", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		check(err)
		defer f.Close()
		_, err = f.Write(wline)
		check(err)
		fmt.Println("Added password for " + service + " with username " + username + " \n")
	}
}
