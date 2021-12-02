package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	quitflag := false

	fmt.Println("Hello there!\nThis is the menu for the Advent of Code 2021, made entirely in Go Language by me, Carmelo Sarta.")

	for !quitflag {
		fmt.Println("\nChoose a day to get the solution o 'q' to quit:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		switch text {
		case "1":
			fmt.Println("\nThe solution for Part 1 is: ")
			D1Part1()
			fmt.Println("\nThe solution for Part 2 is: ")
			D1Part2()
		case "2":
			fmt.Println("\nThe solution for Part 1 is: ")
			D2Part1()
			fmt.Println("\nThe solution for Part 2 is: ")
			D2Part2()
		case "q":
			quitflag = true

		default:
			fmt.Println("This problem hasn't been completed yet!")
		}
	}

}