package refactored_quiz_test

import (
	"bytes"
	"quiz"
	"reflect"
	"strings"
	"testing"
)

func refactored_quiz_test() {

}

func Test_ImportQuizCSV(t *testing.T) {
	tests := []struct {
		input string
		quiz  []quiz.Q
	}{
		{
			"",
			[]quiz.Q{},
		},
		{
			"2+2, 4",
			[]quiz.Q{
				{Question: "2+2", Answer: "4"},
			},
		},
		{
			"2+2,4\n1+2,3",
			[]quiz.Q{
				{Question: "2+2", Answer: "4"},
				{Question: "1+2", Answer: "3"},
			},
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.input)
		quizes, err := quiz.ImportQuizCSV(r)
		if err != nil {
			t.Errorf("ImportQuizCSV(%q) returned error: %v", tt.input, err)
		}
		if got, want := quizes, tt.quiz; !reflect.DeepEqual(got, want) {
			t.Errorf("ImportQuizCSV(%q) = %v, wnated %v", tt.input, got, want)
		}
	}
}

func Test_start_quiz(t *testing.T) {
	quizes := []quiz.Q{
		{Question: "2+2", Answer: "4"},
		{Question: "1+2", Answer: "3"},
	}
	input := "4\n2"
	output := "> Correct!\n > Incorrect\n You had 1/2 correct answers!\n"

	r := strings.NewReader(input)
	var buf bytes.Buffer
	err := quiz.PlayQuiz(&buf, r, quizes)
	if err != nil {
		t.Errorf("play quiz %v with input %q returned error: %v", quizes, input, err)
	}
	if got, want := buf.String(), output; got != want {
		t.Errorf("play quiz %v with input %q\nhave: %q\nwant: %q", quizes, input, got, want)
	}
}
