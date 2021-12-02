package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func PutInputIntoIntArray(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	var in []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
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

func D1Part1() {
	in := HandleIntegerInputFile("day1/in1.txt")
	res := 0

	for i, tmp := 0, in[0]; i < len(in); i++ {
		if in[i] > tmp {
			res += 1
		}
		tmp = in[i]
	}
	fmt.Println(res)
}

func D1Part2() {
	in := HandleIntegerInputFile("day1/in1.txt")

	res := 0
	tmp := in[0] + in[1] + in[2]

	for i := 0; i < len(in)-2; i++ {
		sum := in[i] + in[i+1] + in[i+2]

		if sum > tmp {
			res += 1
		}
		tmp = sum
	}
	fmt.Println(res)
}
