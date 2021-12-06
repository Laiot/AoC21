package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func FishProgression(init_state *[]int, days int) int {
	state := make(map[int]int, 8)

	for _, i := range *init_state {
		state[i] += 1
	}

	for day := 0; day < days; day++ {
		pastState := map[int]int{}
		for i, v := range state {
			pastState[i] = v
		}
		for i := 0; i < 9; i++ {
			switch i {
			case 8:
				state[8] = pastState[0]
			case 6:
				state[6] = pastState[7] + pastState[0]
			default:
				state[i] = pastState[i+1]
			}
		}
	}

	sum := 0
	for _, i := range state {
		sum += i
	}
	return sum
}

func D6Part1() {
	init_state := HandleIntegerInputFile("day6/in6.txt")
	fmt.Println(FishProgression(&init_state, 80))

}

func D6Part2() {
	init_state := HandleIntegerInputFile("day6/in6.txt")
	fmt.Println(FishProgression(&init_state, 256))
}
