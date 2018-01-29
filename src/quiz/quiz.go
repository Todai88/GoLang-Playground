package quiz

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Q struct {
	Question, Answer string
}

func (q Q) ask() bool {
	fmt.Println(q.Question, " equals: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	if scanner.Text() == q.Answer {
		return true
	}
	return false
}

func ImportQuizCSV(r io.Reader) ([]Q, error) {
	var quiz = make([]Q, 0)
	csvr := csv.NewReader(r)
	for {
		record, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		qq := Q{Question: record[0], Answer: record[1]}
		quiz = append(quiz, qq)
	}
	return quiz, nil
}

func PlayQuiz(w io.Writer, r io.Reader, quiz []Q) error {
	correct := 0
	for _, q := range quiz {
		fmt.Printf("%s equals?", q.Question)
		scanner := bufio.NewScanner(r)
		scanner.Scan()
		if scanner.Err() != nil {
			return fmt.Errorf("Scanning %v", scanner.Err())
		}
		input := scanner.Text()

		if input != q.Answer {
			fmt.Fprintln(w, "> Incorrect")
			continue
		}
		fmt.Fprintln(w, "> Correct!")
		correct++
	}
	fmt.Fprintf(w, "You had %d/%d correct answers!\n", correct, len(quiz))
	return nil
}

func run() error {
	var (
		csvFile = flag.String("csv", "problems.csv", "Specify the path to the quiz questions.")
	)
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		return fmt.Errorf("Opening CSV: %v", err)
	}
	defer file.Close()

	quiz, err := ImportQuizCSV(file)
	if err != nil {
		return fmt.Errorf("Importing quiz: %v", err)
	}
	return PlayQuiz(os.Stdout, os.Stdin, quiz)
}

func Play() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
