package main

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file_writer.Check_db_file()
}
