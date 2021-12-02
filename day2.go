package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func D2Part1() {
	var in []string = HandleStringInputFile("inputs/in2.txt")

	x := 0
	y := 0

	for _, line := range in {
		values := strings.Split(line, " ")
		mod, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		switch values[0] {
		case "forward":
			x += mod

		case "down":
			y += mod

		case "up":
			y -= mod

		}
	}
	fmt.Println(x * y)
}

func D2Part2() {
	var in []string = HandleStringInputFile("inputs/in2.txt")

	x := 0
	y := 0
	aim := 0

	for _, line := range in {
		values := strings.Split(line, " ")
		mod, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		switch values[0] {
		case "forward":
			x += mod
			y += aim * mod

		case "down":
			aim += mod

		case "up":
			aim -= mod
		}
	}
	fmt.Println(x * y)
}
