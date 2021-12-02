package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PutInputIntoArray(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var in []string
	for scanner.Scan() {
		i := scanner.Text()
		in = append(in, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return in
}

func HandleInputFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return PutInputIntoArray(file)
}

func Part1(in []string) int {

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
	return x * y
}

func Part2(in []string) int {

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
	return x * y
}

func main() {
	var in []string = HandleInputFile("in.txt")

	fmt.Println(Part1(in))
	fmt.Println(Part2(in))
}
