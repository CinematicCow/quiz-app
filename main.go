package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func Questionire(problems []problem, timeLimit *int) {

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctAnswers := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s =\n", i+1, p.question)
		answerChannel := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerChannel <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("Sorry, you ran out of time. You scored %d out of %d.\n", correctAnswers, len(problems))
			return

		case answer := <-answerChannel:
			if answer == p.answer {
				correctAnswers++
			}

		}
	}
	fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(problems))

}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' ")

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		fmt.Printf("Error: Could not open the file: %s\n", *csvFileName)
		os.Exit(1)
	}

	csvFile := csv.NewReader(file)

	lines, err := csvFile.ReadAll()

	if err != nil {
		fmt.Printf("Error: Failed to parse the CSV file.")
		os.Exit(1)
	}

	problems := parseLines(lines)

	Questionire(problems, timeLimit)
}
