package file_writer

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Check_db_file() (bool) {
	if not os.path.exists("rct_pass_file") {
		file, err := os.Create("rct_pass_file")
		check(err)
		fmt.Println("Password file created")
		defer file.Close()
		d2 := []byte("GoPassword storage file\n")
		n2, err := file.Write(d2)
		check(err)
		fmt.Printf("wrote %d bytes\n", n2)
		return false
	}
	else {
		fmt.Println("File found")
		return true
	}
}