package gopassword

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dispElem(array [][]string, index int) {
	fmt.Println("The password of " + array[index][1] + " is : " + array[index][2])
}

func unlockPassword(splt []string) {
	if len(splt) != 3 {
		displayCommands()
	} else {
		file, err := os.Open("rct_pass_file")
		if err != nil {
			check(err)
		}
		defer file.Close()
		file_array := make([][]string, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			splt_file := strings.Split(scanner.Text(), " ")
			file_array = append(file_array, splt_file)
		}
		for index, line := range file_array {
			if line[0] == splt[1] && line[1] == splt[2] {
				dispElem(file_array, index)
				if err := scanner.Err(); err != nil {
					check(err)
				}
				return
			} else {
				if err := scanner.Err(); err != nil {
					check(err)
				}
			}
		}
		fmt.Println("Service or username not found")
	}
}
