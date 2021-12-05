package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	first point
	last  point
}

func PutInputIntoStringArray(file *os.File) []line {
	scanner := bufio.NewScanner(file)
	var lines []line
	for scanner.Scan() {
		coord := strings.Split(scanner.Text(), " -> ")
		first_coord := strings.Split(coord[0], ",")
		last_coord := strings.Split(coord[1], ",")

		x1, errx1 := strconv.Atoi(first_coord[0])
		y1, erry1 := strconv.Atoi(first_coord[1])
		x2, errx2 := strconv.Atoi(last_coord[0])
		y2, erry2 := strconv.Atoi(last_coord[1])

		if errx1 != nil || errx2 != nil || erry1 != nil || erry2 != nil {
			os.Exit(2)
		}

		first_point := point{x: x1, y: y1}
		last_point := point{x: x2, y: y2}
		lines = append(lines, line{first: first_point, last: last_point})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return lines
}

func HandleStringInputFile(path string) []line {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return PutInputIntoStringArray(file)
}

func D5Part1() {
	fmt.Println(HandleStringInputFile("day5/in5.txt"))
}
