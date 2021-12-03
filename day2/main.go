package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type customScanner struct{ *bufio.Scanner }

type command struct {
	direction string
	by        int
}

func scanner(s *bufio.Scanner) *customScanner { return &customScanner{s} }

func (i *customScanner) Text() command {
	s := strings.Split(i.Scanner.Text(), " ")
	num, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}
	return command{s[0], num}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	chResult := make(chan int)
	ch := dive1(f, chResult)
	go dive2(ch, chResult)

	for result := range chResult {
		fmt.Println(result)
	}
}

func dive1(r io.Reader, chResult chan<- int) chan command {
	ch := make(chan command)
	go func() {
		sc := scanner(bufio.NewScanner(r))
		var horizontal, depth int
		for sc.Scan() {
			comm := sc.Text()
			ch <- comm
			switch comm.direction {
			case "forward":
				horizontal += comm.by
			case "down":
				depth += comm.by
			case "up":
				depth -= comm.by
			}
		}
		close(ch)
		chResult <- horizontal * depth

		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}

	}()

	return ch
}

func dive2(ch <-chan command, chResult chan<- int) {
	var horizontal, depth, aim int
	for comm := range ch {
		switch comm.direction {
		case "forward":
			horizontal += comm.by
			depth += aim * comm.by
		case "down":
			aim += comm.by
		case "up":
			aim -= comm.by
		}
	}
	chResult <- horizontal * depth
	close(chResult)
	return
}
