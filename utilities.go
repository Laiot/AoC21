package main

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
