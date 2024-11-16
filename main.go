package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Options struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
}

type Question struct {
	Question string  `json:"question"`
	Answer   string  `json:"answer"`
	Options  Options `json:"options"`
}

var (
	questions       = map[int]Question{}
	solvedQuestions = map[int]Question{}
	correct         int
)

func (q *Question) New() (Question, bool) {
	for i := 0; i < len(questions); i++ {
		// Check if question[i] is inside solvedQuestions[i]
		if _, ok := solvedQuestions[i]; !ok {
			// Add new question to solvedQuestions
			solvedQuestions[i] = questions[i]
			return questions[i], false
		}
	}
	// Return true if all of the questions are solvedQuestions
	return Question{}, true
}

func main() {
	data, err := os.ReadFile("questions.json")
	if err != nil {
		panic(err)
	}

	var qs []Question
	if err := json.Unmarshal(data, &qs); err != nil {
		panic(err)
	}

	// Add the question slice to the question map with i as key
	for i, v := range qs {
		questions[i] = v
	}

	for {
		var q Question
		q, completed := q.New()
		if completed {
			break
		}

		fmt.Printf(
			"Question: %s\nOptions:\na) %s\nb) %s\nc) %s\nd) %s\n",
			q.Question,
			q.Options.A,
			q.Options.B,
			q.Options.C,
			q.Options.B,
		)

		// Get answer
		var answer string
		fmt.Scanln(&answer)

		// Verify
		if answer == q.Answer {
			correct++ // Increase the amount of correctly answered questions
			fmt.Println("Correct!")
			continue
		}

		fmt.Printf("Wrong!\nCorrect answer was %s\n", q.Answer)
	}

	// Results
	fmt.Printf(
		"Quiz completed!\nTotal correct answers %d out of %d questions\n",
		correct,
		len(solvedQuestions),
	)
}
