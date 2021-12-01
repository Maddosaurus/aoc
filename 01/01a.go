package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("01/01a.txt")
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
	for i := 1; i < len(readings); i++ {
		if readings[i] > readings[i-1] {
			bigger_count++
		}
	}

	fmt.Printf("Number of measurements larger than the previous: %d\n", bigger_count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
