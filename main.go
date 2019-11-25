package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("csv", "problems.csv", "The csv file that contains the quiz in the order 'question,answer'")

	flag.Parse()

	_ = filename
	fmt.Println(*filename)
}
