package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type lowpoint struct {
	value int
	row   int
	col   int
}

func PutInputIntoIntArray(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	var in [][]int
	for scanner.Scan() {
		var tmp []int
		for _, number_char := range scanner.Text() {
			tmp = append(tmp, int(number_char-'0'))
		}
		in = append(in, tmp)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return in
}

func HandleIntegerInputFile(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return PutInputIntoIntArray(file)
}

func GetLowPoints(points [][]int) []lowpoint {
	var lowpoints []lowpoint
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			is_lowpoint := true
			if i != 0 && points[i][j] >= points[i-1][j] {
				is_lowpoint = false
			}

			if i != len(points)-1 && points[i][j] >= points[i+1][j] {
				is_lowpoint = false
			}

			if j != 0 && points[i][j] >= points[i][j-1] {
				is_lowpoint = false
			}

			if j != len(points[0])-1 && points[i][j] >= points[i][j+1] {
				is_lowpoint = false
			}

			if is_lowpoint {
				lowpoints = append(lowpoints, lowpoint{value: points[i][j], row: i, col: j})
			}
		}
	}
	return lowpoints
}

func GetRiskLevel(points [][]int) int {
	lowpoints := GetLowPoints(points)
	sum := 0
	for _, lowpoint := range lowpoints {
		sum += lowpoint.value + 1
	}
	return sum
}

func GetBasin(points [][]int, row int, col int) int {
	if points[row][col] == 9 || points[row][col] == -1 {
		return 0
	}
	points[row][col] = -1

	sum := 1

	if row != 0 {
		sum += GetBasin(points, row-1, col)
	}

	if row != len(points)-1 {
		sum += GetBasin(points, row+1, col)
	}

	if col != 0 {
		sum += GetBasin(points, row, col-1)
	}

	if col != len(points[0])-1 {
		sum += GetBasin(points, row, col+1)
	}

	return sum
}

func GetBasins(points [][]int) []int {
	var basins_size []int
	for _, lowpoint := range GetLowPoints(points) {
		basins_size = append(basins_size, GetBasin(points, lowpoint.row, lowpoint.col))
	}
	return basins_size
}

func GetHighestBasinsSizes(points [][]int, num int) {
	basins := GetBasins(points)
	res := 1
	for i := 0; i < num; i++ {
		basin := 1
		ind := 0
		for tind, tbasin := range basins {
			if tbasin > basin {
				basin = tbasin
				ind = tind
			}
		}
		basins = append(basins[:ind], basins[ind+1:]...)
		res *= basin
	}
	fmt.Println(res)
}

func D9Part1() {
	in := HandleIntegerInputFile("day9/in9.txt")
	fmt.Println(GetRiskLevel(in))
}

func D9Part2() {
	in := HandleIntegerInputFile("day9/in9.txt")
	GetHighestBasinsSizes(in, 3)
}
