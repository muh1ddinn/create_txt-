package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func writeToFile(filename string, data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {
	ff, _ := os.Open("data.txt") // assuming the numbers are in input.txt
	defer ff.Close()

	s := bufio.NewScanner(ff)
	var sum int

	for s.Scan() {
		// Parse the current line into an integer
		line := s.Text()
		if strings.ContainsAny(line, "1234567890") {
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error parsing line:", err)
				return // Skip to the next line if there's an error
			}

			// Write the number to the output file
			writeToFile("output.txt", fmt.Sprintf("%d\n", num))

			// Accumulate the number to the sum
			sum += num
		}
	}

	// Write the sum to the output file
	writeToFile("worddd.txt", fmt.Sprintf("sum = %d", sum))
}
