package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("01b.txt") // Technically hasn't changed, but eh.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var readings []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intval, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		readings = append(readings, intval)
	}

	var bigger_count int = 0
	var previous_window = (readings[0] + readings[1] + readings[2])
	for i := 2; i < len(readings)-1; i++ {
		current_window := readings[i-1] + readings[i] + readings[i+1]
		if current_window > previous_window {
			bigger_count++
		}
		previous_window = current_window
	}

	fmt.Printf("Number of window measurements larger than the previous: %d\n", bigger_count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
