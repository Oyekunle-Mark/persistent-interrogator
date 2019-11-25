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

	correctCount := 0

	for index, problem := range parsed {
		fmt.Printf("Question #%d: %s = ", index+1, problem.q)

		var response string
		fmt.Scanf("%s", &response)

		if response == problem.a {
			correctCount++
		}

	}

	fmt.Printf("You scored %d, out of %d questions\n", correctCount, len(parsed))
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
