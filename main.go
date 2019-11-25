package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("csv", "problems.csv", "The csv file that contains the quiz in the order 'question,answer'")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		fmt.Printf("Cannot open the file %s\n", *filename)
		os.Exit(1)
	}
	_ = file
}
