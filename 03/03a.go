package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("03a.txt") // Technically hasn't changed, but eh.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//value_add := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // <-- same as:
	value_add := make([]int, 12)
	line_count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line_count++
		line := scanner.Text()
		for pos, achar := range line {
			aint := int(achar - '0')
			value_add[pos] = value_add[pos] + aint
		}
	}
	var gamma_rate []string
	var epsilon_rate []string

	for _, ipos := range value_add {
		if ipos > line_count/2 {
			gamma_rate = append(gamma_rate, strconv.Itoa(1))
			epsilon_rate = append(epsilon_rate, strconv.Itoa(0))
		} else {
			gamma_rate = append(gamma_rate, strconv.Itoa(0))
			epsilon_rate = append(epsilon_rate, strconv.Itoa(1))
		}
	}
	int_gamma, err := strconv.ParseInt(strings.Join(gamma_rate, ""), 2, 64)
	int_epsilon, err := strconv.ParseInt(strings.Join(epsilon_rate, ""), 2, 64)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d", int_gamma)
	fmt.Printf("\n%d", int_epsilon)
	fmt.Printf("\nFinal power consumption: %d", int_epsilon*int_gamma)
}
