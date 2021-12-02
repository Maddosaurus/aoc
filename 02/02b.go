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
	file, err := os.Open("02a.txt") // Technically hasn't changed, but eh.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cmd []string
	var cmd_value []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), " ")
		cmd = append(cmd, splitted[0])
		cast_value, err := strconv.Atoi(splitted[1])

		if err != nil {
			log.Fatal(err)
		}

		cmd_value = append(cmd_value, cast_value)
	}
	fmt.Println("Done")

	var depth, position, aim int = 0, 0, 0
	for i := 0; i < len(cmd); i++ {
		switch cmd[i] {
		case "forward":
			position = position + cmd_value[i]
			depth = depth + (aim * cmd_value[i])
			fmt.Printf("Forward: %d - Position is now %d\n", cmd_value[i], position)
			fmt.Printf("Depth is now %d\n", depth)
		case "down":
			aim = aim + cmd_value[i]
			fmt.Printf("Down: %d - Aim is now %d\n", cmd_value[i], aim)
		case "up":
			aim = aim - cmd_value[i]
			fmt.Printf("Up: %d - Aim is now %d\n", cmd_value[i], aim)
		}
	}
	fmt.Printf("Final Depth: %d\nFinal Position: %d\nAnswer:%d", depth, position, depth*position)
}
