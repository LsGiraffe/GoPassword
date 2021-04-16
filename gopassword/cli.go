package gopassword

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func displayCommands() {
	fmt.Println("Commands :")
	fmt.Println("- Add an account -> \"add 'service' 'username' 'password'\"")
	fmt.Println("- Remove an account -> \"remove 'service' 'username'\"")
	fmt.Println("- Modify an account -> \"mod 'service' 'username' 'password'\"")
}

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

func Run() {
	displayCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n-->  ")
		scanner.Scan()
		input := scanner.Text()
		splt := strings.Split(input, " ")
		if input == "exit" {
			return
		} else if splt[0] == "add" {
			addCmd(splt)
		} else if splt[0] == "remove" {
			println("REMOVE")
		} else if splt[0] == "mod" {
			println("MOD")
		} else {
			println("unknown")
		}
	}
}
