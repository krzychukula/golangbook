package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		//handle err
		return
	}
	defer file.Close()

	file.WriteString("test")

}
