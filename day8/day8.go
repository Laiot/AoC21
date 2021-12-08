package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PutInputIntoStringArray(file *os.File) (inputs []string, outputs []string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")

		inputs = append(inputs, strings.Split(line[0], " ")...)
		outputs = append(outputs, strings.Split(line[1], " ")...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return inputs, outputs
}

func HandleStringInputFile(path string) (inputs []string, outputs []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return PutInputIntoStringArray(file)
}

func PutInputIntoStringArrays(file *os.File) (inputs [][]string, outputs [][]string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")

		inputs = append(inputs, strings.Split(line[0], " "))
		outputs = append(outputs, strings.Split(line[1], " "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return inputs, outputs
}

func HandleStringInputFileBis(path string) (inputs [][]string, outputs [][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return PutInputIntoStringArrays(file)
}

func GetUniqueDigits(outputs []string) int {
	res := 0
	for _, output := range outputs {
		if len(output) == 2 || len(output) == 4 || len(output) == 3 || len(output) == 7 {
			res++
		}
	}
	return res
}

func arrayContains(a []rune, x rune) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func deduceDash(inputdashes []string) map[rune]rune {

	var (
		DASHES map[int][]rune = map[int][]rune{
			0: {'a', 'b', 'c', 'e', 'f', 'g'},
			1: {'c', 'f'},
			2: {'a', 'c', 'd', 'e', 'g'},
			3: {'a', 'c', 'd', 'f', 'g'},
			4: {'b', 'c', 'd', 'f'},
			5: {'a', 'b', 'd', 'f', 'g'},
			6: {'a', 'b', 'd', 'e', 'f', 'g'},
			7: {'a', 'c', 'f'},
			8: {'a', 'b', 'c', 'd', 'e', 'f', 'g'},
			9: {'a', 'b', 'c', 'd', 'f', 'g'},
		}
	)

	dashes := map[rune]rune{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
	}

	possibleDash := map[int][]rune{
		0: {},
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
		9: {},
	}

	// 1, 4, 7, and 8 can be deduced solely by their length
	for _, dash := range inputdashes {
		if len(dash) == len(DASHES[1]) {
			possibleDash[1] = append(possibleDash[1], []rune(dash)...)
		} else if len(dash) == len(DASHES[4]) {
			possibleDash[4] = append(possibleDash[4], []rune(dash)...)
		} else if len(dash) == len(DASHES[7]) {
			possibleDash[7] = append(possibleDash[7], []rune(dash)...)
		} else if len(dash) == len(DASHES[8]) {
			possibleDash[8] = append(possibleDash[8], []rune(dash)...)
		}
	}

	// 'a'
	for _, char := range possibleDash[7] {
		if !arrayContains(possibleDash[1], char) {
			dashes['a'] = char
		}
	}

	// 'c' and 'f'
	for _, dash := range inputdashes {
		if len(dash) == len(DASHES[6]) {
			// We found 'c' and therefore identify 'f' as well
			if !arrayContains([]rune(dash), possibleDash[1][0]) {
				possibleDash[6] = append(possibleDash[6], []rune(dash)...)
				dashes['c'] = possibleDash[1][0]
				dashes['f'] = possibleDash[1][1]
			} else if !arrayContains([]rune(dash), possibleDash[1][1]) {
				possibleDash[6] = append(possibleDash[6], []rune(dash)...)
				dashes['c'] = possibleDash[1][1]
				dashes['f'] = possibleDash[1][0]
			}
		}
	}

	for _, dash := range inputdashes {
		if len(dash) == len(DASHES[2]) {
			if arrayContains([]rune(dash), dashes['c']) && arrayContains([]rune(dash), dashes['f']) {
				// We found 3
				possibleDash[3] = append(possibleDash[3], []rune(dash)...)
			} else if !arrayContains([]rune(dash), dashes['f']) {
				// We found 2
				possibleDash[2] = append(possibleDash[2], []rune(dash)...)
			} else if !arrayContains([]rune(dash), dashes['c']) {
				// We found 5
				possibleDash[5] = append(possibleDash[5], []rune(dash)...)
			}
		}
	}

	// 'b' and 'e'
	for _, r := range possibleDash[2] {
		if r != dashes['c'] {
			if !arrayContains(possibleDash[5], r) {
				dashes['e'] = r
			}
		}
	}
	for _, r := range possibleDash[5] {
		if r != dashes['f'] {
			if !arrayContains(possibleDash[2], r) {
				dashes['b'] = r
			}
		}
	}

	for _, dash := range inputdashes {
		if len(dash) == len(DASHES[0]) {
			if !arrayContains([]rune(dash), dashes['e']) {
				possibleDash[9] = append(possibleDash[9], []rune(dash)...)
			} else if arrayContains([]rune(dash), dashes['c']) {
				possibleDash[0] = append(possibleDash[0], []rune(dash)...)
			}
		}
	}

	// Solve 'd'
	for r := 'a'; r <= 'g'; r++ {
		if !arrayContains(possibleDash[0], r) {
			dashes['d'] = r
			break
		}
	}

	// Solve 'g'
	for r := 'a'; r <= 'g'; r++ {
		if _, ok := dashes[r]; !ok {
			dashes['g'] = r
		}
	}

	return dashes
}

func translateDigit(dashes map[rune]rune, digit string) int {

	var (
		DASHES map[int][]rune = map[int][]rune{
			0: {'a', 'b', 'c', 'e', 'f', 'g'},
			1: {'c', 'f'},
			2: {'a', 'c', 'd', 'e', 'g'},
			3: {'a', 'c', 'd', 'f', 'g'},
			4: {'b', 'c', 'd', 'f'},
			5: {'a', 'b', 'd', 'f', 'g'},
			6: {'a', 'b', 'd', 'e', 'f', 'g'},
			7: {'a', 'c', 'f'},
			8: {'a', 'b', 'c', 'd', 'e', 'f', 'g'},
			9: {'a', 'b', 'c', 'd', 'f', 'g'},
		}
	)

	if len(digit) == len(DASHES[1]) {
		return 1
	} else if len(digit) == len(DASHES[4]) {
		return 4
	} else if len(digit) == len(DASHES[7]) {
		return 7
	} else if len(digit) == len(DASHES[8]) {
		return 8
	}

	// 0, 6, 9
	if len(digit) == len(DASHES[0]) {
		if !arrayContains([]rune(digit), dashes['d']) {
			return 0
		} else if !arrayContains([]rune(digit), dashes['c']) {
			return 6
		} else {
			return 9
		}
	}

	// 2, 3, 5
	if len(digit) == len(DASHES[2]) {
		if arrayContains([]rune(digit), dashes['b']) {
			return 5
		} else if arrayContains([]rune(digit), dashes['e']) {
			return 2
		} else {
			return 3
		}
	}

	return 0
}

func getNumber(dashes map[rune]rune, output []string) int {
	answer := 0
	for _, digit := range output {
		answer *= 10
		answer += translateDigit(dashes, digit)
	}
	return answer
}

func D8Part1() {
	_, outputs := HandleStringInputFile("day8/in8.txt")
	fmt.Println(GetUniqueDigits(outputs))
}

func D8Part2() {
	inputs, outputs := HandleStringInputFileBis("day8/in8.txt")
	sum := 0

	for i := 0; i < len(inputs); i++ {
		sum += getNumber(deduceDash(inputs[i]), outputs[i])
	}

	fmt.Println(sum)
}
