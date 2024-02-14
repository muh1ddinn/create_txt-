package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var word []string

	for s.Scan() {
		// Parse the current line into an integer
		line := s.Text()
		if strings.ContainsAny(line, "1234567890") {
			num, err := strconv.Atoi(line)

			if err != nil {
				fmt.Println("Error parsing line:", err)
				return // Skip to the next line if there's an error
			}
			fmt.Println(num)
			// Accumulate the number to the sum
			sum += num
		} else {

			for i := 0; i < len(line); i++ {

				if line[i] == 'A' || line[i] == 'O' || line[i] == 'a' || line[i] == 'o' {

					word = append(word, line)
					//fmt.Println(word)
				}

			}

		}
	}
	// Print the total sum
	fmt.Printf(" = %d\n", sum)

	writeresufile("wordd.txt", word)

}

func writeresufile(filename string, data string) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("err while creat write file ")

		return
	}

	defer file.Close()

	_, err = file.WriteString(strings.Join(result, "\n"))

	if err != nil {
		fmt.Println("erre while writing result")
	}

}
