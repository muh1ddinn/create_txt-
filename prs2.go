package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open the file
	ff, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer ff.Close()

	// Create a scanner to read the file line by line
	s := bufio.NewScanner(ff)
	var sum int
	first := true // For printing the sum in the desired format
	for s.Scan() {
		// Parse the current line into an integer
		line := s.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue // Skip to the next line if there's an error
		}
		// Accumulate the number to the sum
		sum += num
		// Print the intermediate sum
		if first {
			fmt.Printf("%d", num)
			first = false
		} else {
			fmt.Printf(" + %d", num)
		}
	}
	// Print the total sum
	fmt.Printf(" = %d\n", sum)

	// Check for scanner errors
	if err := s.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}
