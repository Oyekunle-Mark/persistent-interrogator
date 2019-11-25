package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("csv", "problems.csv", "The csv file that contains the quiz in the order 'question,answer'")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		exit(fmt.Sprintf("Cannot open the file %s\n", *filename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Cannot reading the csv file")
	}

	parsed := parseLines(lines)
	fmt.Println(parsed)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problemSlice := make([]problem, len(lines))

	for index, line := range lines {
		problemSlice[index] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return problemSlice
}

func exit(errorMsg string) {
	fmt.Println(errorMsg)
	os.Exit(1)
}
