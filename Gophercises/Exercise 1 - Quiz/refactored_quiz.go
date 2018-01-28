import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

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

func importQuizCSV(r io.Reader) ([]q, error) {
	var quiz = make([]q, 0)
	csvr := csv.NewReader(r)
	for {
		record, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		qq := q{question: record[0], answer: record[1]}
		quiz = append(quiz, qq)
	}
	return quiz, nil
}

func playQuiz(w io.Writer, r io.Reader, quiz []q) error {
	correct := 0
	for _, q := range quiz {
		fmt.Printf("%s equals?", q.question)
		scanner := bufio.NewScanner(r)
		scanner.Scan()
		if scanner.Err() != nil {
			return fmt.Errorf("Scanning %v", scanner.Err())
		}
		input := scanner.Text()

		if input != q.answer {
			fmt.Fprintln(w, "> Incorrect")
			continue
		}
		fmt.Fprintln(w, "> Correct!")
		correct++
	}
	fmt.Fprintf(w, "You had %d/%d correct answers!\n", correct, len(quiz))
	return nil
}
 
func run () error {
	var {
		csvFile = flag.String("csv", "problems.csv", "Specify the path to the quiz questions.")
	}
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		return fmt.Errorf("Opening CSV: %v", err)
	}
	defer file.Close()

	quiz, err := importQuizCSV(file)
	if err != nil {
		return fmt.Errorf("Importing quiz: %v", err)
	}
	return playQuiz(os.Stdout, os.Stdin, quiz)
}

func init() {
	if err := run(); err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }
}