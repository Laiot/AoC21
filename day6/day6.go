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

func NextDay(state *[]int) {
	for idx := range *state {
		if (*state)[idx] == 0 {
			(*state)[idx] = 6
			*state = append(*state, 8)
		} else {
			(*state)[idx]--
		}
	}
}

func SkipTime(state *[]int, days int) {
	for i := 0; i < days; i++ {
		NextDay(state)
	}
}

func D6Part1() {
	init_state := HandleIntegerInputFile("day6/in6.txt")
	SkipTime(&init_state, 80)
	fmt.Println(len(init_state))
}

func D6Part2() {

}
