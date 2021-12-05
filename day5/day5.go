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

func PutInputIntoArray(file *os.File) []line {
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
	return PutInputIntoArray(file)
}

func LineIsStraight(line line) (bool, rune) {
	if line.first.x == line.last.x {
		return true, 'x'
	} else if line.first.y == line.last.y {
		return true, 'y'
	} else {
		return false, 'o'
	}
}

func DrawGroundPart1(lines []line) [1000][1000]int {
	ground := [1000][1000]int{}
	for _, line := range lines {
		is_straight, fixed_coord := LineIsStraight(line)
		if is_straight {
			if fixed_coord == 'x' {
				if line.first.y < line.last.y {
					for y := line.first.y; y <= line.last.y; y++ {
						ground[line.first.x][y] += 1
					}
				} else {
					for y := line.last.y; y <= line.first.y; y++ {
						ground[line.first.x][y] += 1
					}
				}

			} else {
				if line.first.x < line.last.x {
					for x := line.first.x; x <= line.last.x; x++ {
						ground[x][line.first.y] += 1
					}
				} else {
					for x := line.last.x; x <= line.first.x; x++ {
						ground[x][line.first.y] += 1
					}
				}

			}
		}
	}
	return ground
}

func LineIsDiagonal(line line) (bool, rune) {
	if (line.first.x-line.last.x)/(line.first.y-line.last.y) == 1 {
		return true, 'a'
	} else if (line.first.x-line.last.x)/(line.first.y-line.last.y) == -1 {
		return true, 'b'
	} else {
		return false, 'o'
	}
}

func DrawGroundPart2(lines []line) [1000][1000]int {
	ground := [1000][1000]int{}
	for _, line := range lines {
		is_straight, fixed_coord := LineIsStraight(line)
		if is_straight {
			if fixed_coord == 'x' {
				if line.first.y < line.last.y {
					for y := line.first.y; y <= line.last.y; y++ {
						ground[line.first.x][y] += 1
					}
				} else {
					for y := line.last.y; y <= line.first.y; y++ {
						ground[line.first.x][y] += 1
					}
				}

			} else {
				if line.first.x < line.last.x {
					for x := line.first.x; x <= line.last.x; x++ {
						ground[x][line.first.y] += 1
					}
				} else {
					for x := line.last.x; x <= line.first.x; x++ {
						ground[x][line.first.y] += 1
					}
				}

			}
		} else {
			is_diagonal, d_type := LineIsDiagonal(line)
			if is_diagonal {
				if d_type == 'a' {
					if line.first.x < line.last.x {
						for i := line.first.x; i <= line.last.x; i++ {
							ground[i][i] += 1
						}
					} else {
						for i := line.last.x; i <= line.first.x; i++ {
							ground[i][i] += 1
						}
					}
				} else {
					if line.first.x < line.last.x {
						for i, j := line.first.x, line.first.y; i <= line.last.x; i++ {
							ground[i][j] += 1
							j--
						}
					} else {
						for i, j := line.last.x, line.last.y; i <= line.first.x; i++ {
							ground[i][j] += 1
							j--
						}
					}
				}
			}
		}

	}
	return ground
}

func CheckGround(ground [1000][1000]int, limit int) int {
	res := 0
	for i := 0; i < len(ground); i++ {
		for j := 0; j < len(ground); j++ {
			if ground[i][j] > limit {
				res++
			}
		}
	}
	return res
}

func D5Part1() {
	lines := HandleStringInputFile("day5/in5.txt")
	ground := DrawGroundPart1(lines)
	fmt.Println(CheckGround(ground, 1))
}

func D5Part2() {
	lines := HandleStringInputFile("day5/in5.txt")
	ground := DrawGroundPart2(lines)
	fmt.Println(ground)
	fmt.Println(CheckGround(ground, 1))
}
