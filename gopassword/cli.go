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
			println("ADD")
		} else if splt[0] == "remove" {
			println("REMOVE")
		} else if splt[0] == "mod" {
			println("MOD")
		} else {
			println("unknown")
		}
	}
}
