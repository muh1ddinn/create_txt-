package main

import (
	"fmt"
	"os"
)

func main() {

	// Open the file

	fileee, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("err while open file")

		// Handle the error

		return
	}

	defer fileee.Close()

	datlen, err := fileee.Stat()
	if err != nil {

		fmt.Println("err while datlen")
	}

	fmt.Println(datlen.Size())

	//Create a scanner to read the file
	//
	//annerr := bufio.NewScanner(fileee)
	//
	// Check for errors
	// err := scannerr.Err(); err != nil {
	//
	//	fmt.Println("err while scaning ", err)
	//
	//	// Handle the error
	//	return
	//
	//
	//Read the file line by line
	//
	//x := 0
	//r scannerr.Scan() {
	//
	//	for i := 0; i < len(scannerr); i++ {
	//
	//		max = int(max + line[i])
	//
	//	}
	//
	//	fmt.Println(max)
	//
	//
	//
} //
