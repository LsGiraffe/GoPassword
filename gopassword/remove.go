package gopassword

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func removeElem(array [][]string, index int) [][]string {
	array[index] = array[len(array)-1]
	array[len(array)-1] = nil
	array = array[:len(array)-1]
	return array
}

func writeArrayInFile(array [][]string) {
	file, err := os.Open("rct_pass_file")
	if err != nil {
		check(err)
	}
	defer file.Close()
	write_string := ""
	for _, line := range array {
		concatString := strings.Join(line, " ")
		write_string = write_string + concatString + "\n"
		byte_write_string := []byte(write_string)
		err := ioutil.WriteFile("rct_pass_file", byte_write_string, 0644)
		check(err)
	}
	fmt.Println("Removal complete")
}

func removeCmd(splt []string) {
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
				file_array = removeElem(file_array, index)
				writeArrayInFile(file_array)
				if err := scanner.Err(); err != nil {
					check(err)
				}
				return
			} else {
				if err := scanner.Err(); err != nil {
					check(err)
				}
				fmt.Println("Service or username not found")
			}
		}
	}
}
