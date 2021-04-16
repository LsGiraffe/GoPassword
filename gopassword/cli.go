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
			addCmd(splt)
		} else if splt[0] == "remove" {
			removeCmd(splt)
		} else if splt[0] == "mod" {
			println("MOD")
		} else {
			println("Unknown command")
		}
	}
}
