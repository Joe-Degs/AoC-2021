package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readNumbers(path string) []int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var input []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, num)
	}
	return input
}

func puzzle2(input []int) int {
	var prev, result int
	for ii, num := range input {
		window := num + input[ii+1] + input[ii+2]

		if ii == 1 {
			prev = window
			continue
		}

		cur := window

		if cur > prev {
			result += 1
		}
		prev = cur

		if len(input)-3 == ii {
			break
		}
	}
	return result
}

func puzzle1(input []int) int {
	var prev, result int
	for ii, num := range input {
		if ii == 1 {
			prev = num
			continue
		}

		cur := num

		if cur > prev {
			result += 1
		}
		prev = cur
	}
	return result
}

func main() {
	input := readNumbers("input.txt")
	fmt.Println(puzzle1(input), puzzle2(input))
}
