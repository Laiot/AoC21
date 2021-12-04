package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type board struct {
	values  [][]int
	checked bool
}

func PutInputIntoIntArray(file *os.File) ([]int, []board) {
	scanner := bufio.NewScanner(file)
	fline_bool := true
	var draws []int
	var boards []board
	var tmp2 [][]int
	for scanner.Scan() {
		if fline_bool {
			for _, elem := range strings.Split(scanner.Text(), ",") {
				i, err := strconv.Atoi(elem)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				draws = append(draws, i)
			}
			fline_bool = false
		} else {
			if len(scanner.Text()) != 0 {
				var tmp1 []int
				for _, elem := range strings.Fields(scanner.Text()) {
					intVar, err := strconv.Atoi(elem)
					if err != nil {
						fmt.Println(err)
						os.Exit(2)
					}
					tmp1 = append(tmp1, intVar)
				}
				tmp2 = append(tmp2, tmp1)
			}
			if len(tmp2) == 5 {
				boards = append(boards, board{tmp2, false})
				tmp2 = nil
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return draws, boards
}

func HandleIntegerInputFile(path string) ([]int, []board) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return PutInputIntoIntArray(file)
}

func CheckBingo(board board) bool {
	for i := 0; i < len(board.values); i++ {
		res1 := true
		res2 := true
		for j := 0; j < len(board.values); j++ {
			if board.values[i][j] != -1 {
				res1 = false
				break
			}
		}
		for j := 0; j < len(board.values); j++ {
			if board.values[j][i] != -1 {
				res2 = false
				break
			}
		}
		if res1 || res2 {
			return res1 || res2
		}
	}

	return false
}

func SumUnchecked(board board) int {
	sum := 0
	for i := 0; i < len(board.values); i++ {
		for j := 0; j < len(board.values); j++ {
			if board.values[i][j] != -1 {
				sum += board.values[i][j]
			}
		}
	}
	return sum
}

func CheckBoards(draws []int, boards []board) int {
	for _, draw := range draws {
		for _, board := range boards {
			for i := 0; i < len(board.values); i++ {
				for j := 0; j < len(board.values); j++ {
					if board.values[i][j] == draw {
						board.values[i][j] = -1
						if CheckBingo(board) {
							return SumUnchecked(board) * draw
						}
					}
				}
			}
		}

	}
	return -1
}

func FindLastBoard(draws []int, boards []board) int {
	for _, draw := range draws {
		for k := 0; k < len(boards); k++ {
			for i := 0; i < len(boards[k].values); i++ {
				for j := 0; j < len(boards[k].values); j++ {
					if boards[k].values[i][j] == draw {
						boards[k].values[i][j] = -1
						if !boards[k].checked && CheckBingo(boards[k]) {
							boards[k].checked = true
							counter := 0
							for _, b := range boards {
								if !b.checked {
									counter++
								}
							}
							if counter == 0 || draw == draws[len(draws)-1] {
								return SumUnchecked(boards[k]) * draw
							}
						}
					}
				}
			}
		}
	}
	return -1
}

func D4Part1() {
	draws, boards := HandleIntegerInputFile("day4/in4.txt")
	fmt.Println(CheckBoards(draws, boards))
}

func D4Part2() {
	draws, boards := HandleIntegerInputFile("day4/in4.txt")
	fmt.Println(FindLastBoard(draws, boards))
}
