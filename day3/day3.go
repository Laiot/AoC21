package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FromStringToBinary(s string) int {
	val, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(val)
}

func PutInputIntoStringArray(file *os.File) []string {
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

func HandleStringInputFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return PutInputIntoStringArray(file)
}

func MostCommon(lines []string) string {
	result := make([]byte, len(lines[0]))
	for i := range lines[0] {
		counter := make(map[byte]int)
		for _, line := range lines {
			c := line[i]
			v := counter[c]
			counter[c] = v + 1
		}
		if counter['0'] > counter['1'] {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return string(result)
}

func LeastCommon(lines []string) string {
	mostCommon := MostCommon(lines)
	result := ""
	for _, v := range mostCommon {
		if v == '1' {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func GetAir(input []string, patternFn func([]string) string) string {
	validLines := make(map[string]bool)
	for _, l := range input {
		validLines[l] = true
	}
	for i := 0; ; i++ {
		currentLines := make([]string, len(validLines))
		j := 0
		for k := range validLines {
			currentLines[j] = string(k[i])
			j++
		}
		pattern := patternFn(currentLines)
		for l := range validLines {
			if l[i] != pattern[0] {
				delete(validLines, l)
			}
			if len(validLines) == 1 {
				return l
			}
		}
	}
}

func D3Part1() {
	lines := HandleStringInputFile("day3/in3.txt")
	gammaStr := MostCommon(lines)
	gamma := FromStringToBinary(gammaStr)
	epsilon := gamma ^ FromStringToBinary(strings.Repeat("1", len(gammaStr)))
	fmt.Println(gamma * epsilon)
}

func D3Part2() {
	lines := HandleStringInputFile("day3/in3.txt")
	oxygen := GetAir(lines, MostCommon)
	co2 := GetAir(lines, LeastCommon)
	fmt.Println(FromStringToBinary(oxygen) * FromStringToBinary(co2))
}
