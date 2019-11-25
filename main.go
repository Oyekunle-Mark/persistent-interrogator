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
	line, err := r.ReadAll()

	if err != nil {
		exit("Cannot reading the csv file")
	}
	fmt.Println(line)
}

func exit(errorMsg string) {
	fmt.Println(errorMsg)
	os.Exit(1)
}
