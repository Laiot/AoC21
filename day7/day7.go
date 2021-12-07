package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func PutInputIntoIntArray(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	var in []int
	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")

	for _, val := range line {
		i, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		in = append(in, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return in
}

func HandleIntegerInputFile(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return PutInputIntoIntArray(file)
}

func GetMedian(ar []int) int {
	sort.Ints(ar)
	if len(ar)%2 == 0 {
		return (ar[len(ar)/2] + ar[(len(ar)/2)-1]) / 2
	} else {
		return ar[(len(ar)-1)/2]
	}
}

func GetAverage(ar []int) int {
	total := 0
	for _, number := range ar {
		total = total + number
	}
	return total / len(ar)
}

func PriceProgression(price int) int {
	return price * (price + 1) / 2
}

func FindMinimumCost(init_state []int) int {
	least_expensive_position := GetMedian(init_state)
	sum := 0
	for _, val := range init_state {
		if val-least_expensive_position < 0 {
			sum += least_expensive_position - val
		} else {
			sum += val - least_expensive_position
		}
	}
	return sum
}

func FindNewMinimumCost(init_state []int) int {
	least_expensive_guess := GetAverage(init_state)
	min := -1
	for i := least_expensive_guess - 1; i < least_expensive_guess+2; i++ {
		sum := 0
		for _, val := range init_state {
			if val-least_expensive_guess < 0 {
				sum += PriceProgression(least_expensive_guess - val)
			} else {
				sum += PriceProgression(val - least_expensive_guess)
			}
		}
		if sum < min || min == -1 {
			min = sum
		}
	}
	return min
}

func D7Part1() {
	init_state := HandleIntegerInputFile("day7/in7.txt")
	fmt.Println(FindMinimumCost(init_state))
}

func D7Part2() {
	init_state := HandleIntegerInputFile("day7/in7.txt")
	fmt.Println(FindNewMinimumCost(init_state))
}
