package main

import (
	"fmt"
)

func D1Part1() {
	in := HandleIntegerInputFile("inputs/in1.txt")
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
	in := HandleIntegerInputFile("inputs/in1.txt")

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
