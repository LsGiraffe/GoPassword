package main

import "github.com/LsGiraffe/GoPassword/gopassword"

func main() {
	gopassword.Check_db_file()
	gopassword.Run()
	return
}
