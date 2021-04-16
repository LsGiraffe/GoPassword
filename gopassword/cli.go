package gopassword

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Run() {
	_, err := os.Open("rct_pass_file") // For read access.
	if err != nil {
		panic(errors.New("Unable to open the file"))
	}
	fmt.Println("Commands :")
	fmt.Println("- Add an account -> \"add 'service' 'username' 'password'\"")
	fmt.Println("- Remove an account -> \"remove 'service' 'username'\"")
	fmt.Println("- Modify an account -> \"mod 'service' 'username' 'password'\"")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n-->  ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			return
		}
	}
}
