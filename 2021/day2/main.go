package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	filepath := "input.txt"

	file, err := os.Open(filepath)

	var direction [2]int
	direction[0] = 0
	direction[1] = 0
	aim := 0

	if err != nil {
		fmt.Println("Trouble Reading File, Adjust File Path in Code.")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for _, line := range lines {
		parts := strings.Fields(line) // splits on whitespace
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}
		value, err := strconv.Atoi(parts[1])
		move := parts[0]
		if err != nil {
			fmt.Println("Invalid number:", parts[1])
			continue
		}

		switch move {
		case "up":
			fmt.Println("Aiming Up")
			aim -= value
			fmt.Printf("Your Aim is: %v \n", aim)
		case "down":
			fmt.Println("Aiming Down")
			aim += value
			fmt.Printf("Your Aim is: %v \n", aim)
		case "forward":
			fmt.Println("Moving")
			direction[0] += value
			direction[1] += value * aim
			fmt.Printf("[%v, %v]\n", direction[0], direction[1])

		}
	}
	fmt.Printf("[%v,%v]\n", direction[0], direction[1])

	var finalvalue int = direction[0] * direction[1]

	fmt.Println(finalvalue)

	// Now I have to factor in aim
	// down X increases your aim by X units.
	// up X decreases your aim by X units.
	// forward X does two things:
	//     It increases your horizontal position by X units.
	//     It increases your depth by your aim multiplied by X.

}
