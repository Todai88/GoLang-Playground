package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// a test
type q struct {
	question, answer string
}

func (q q) ask() bool {
	fmt.Println(q.question, " equals: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	if scanner.Text() == q.answer {
		return true
	}
	return false
}

func quizLoop(path string, verbose bool) {
	// Loop should:
	// 1. Read records line by line
	// 2. Ask the question (i/o)
	// 3. Keep score.
	file, err := os.Open(path)
	correct, lines := 0, 0

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		q := q{question: record[0], answer: record[1]}
		if q.ask() {
			if verbose {
				fmt.Println("Correct")
			}
			correct++
		} else if verbose {
			fmt.Println("Incorrect")
		}
		lines++
	}
	fmt.Printf("You had %d/%d correct answers!\n", correct, lines)
}

func main() {
	// Setup flags.
	p := flag.String("path", "problems.csv", "Specify the path to the quiz questions.")
	v := flag.Bool("verbose", false, "A boolean value to check if you want the program to be verbose or not.")
	flag.Parse()

	// Invoke loop.
	quizLoop(*p, *v)
}
