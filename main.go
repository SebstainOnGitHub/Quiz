package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func parseCSV(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	problems, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return problems
}

func parseProblems(problems [][]string) map[string]int {
	h := make(map[string]int, 30)
	for i := 0; i < len(problems); i++ {
		for j := 0; j < len(problems[i]); j++ {
			if _, err := strconv.Atoi(problems[i][j]); err != nil { //I can't even remember why this works lol
				resultint, err := strconv.Atoi(problems[i][j+1])
				if err != nil {
					log.Fatal(err)
				}
				h[problems[i][j]] = resultint
			}
		}
	}
	return h
}

func blockUntilEnter(scan *bufio.Scanner) {
	for {
		fmt.Println("Please press enter to start")
		scan.Scan()
		if scan.Text() == "" {
			break
		}
	}
}

func main() {
	answered, score := 0, 0
	outOfTime := false

	var endTime = flag.Int("time", 40, "controls the amount of time you have to complete the quiz (in seconds)")
	var newPtr = flag.Bool("new", false, "creates a new quiz with the same number of questions")

	flag.Parse()

	if *newPtr {
		newQuestions()
	}

	problems := parseCSV("problems.csv")
	problemMap := parseProblems(problems)

	scan := bufio.NewScanner(os.Stdin)

	blockUntilEnter(scan)

	start := time.Now()
	var end time.Duration

	timer := time.NewTimer(time.Duration(*endTime) * time.Second)

	go func() {
		<-timer.C
		end = time.Since(start)
		fmt.Println("\nYou ran out of time!")
		fmt.Println("Press any key to see your stats")
		outOfTime = true
	}()

	for key, value := range problemMap {
		fmt.Print(key + " = ")

		scan.Scan()

		if outOfTime {
			break
		}

		if scan.Text() == "" {
			fmt.Println("Incorrect, your score is still", score)
			continue
		}

		answered++

		num, err := strconv.Atoi(scan.Text())

		if err != nil {
			fmt.Println("Incorrect, your score is still", score)
			continue
		}

		if num == value {
			score++
			fmt.Println("Correct, your score is now", score)
		} else {
			fmt.Println("Incorrect, your score is still", score)
		}

		if outOfTime {
			break
		}
	}

	if answered == 0 {answered = 1} // to prevent division by 0

	fmt.Println("\nYour total accuracy was", math.Round(float64(score)/float64(answered))*100, "%")

	if outOfTime {
		fmt.Println("You answered", answered, "questions in", math.Round(float64(end)), "seconds")
	} else {
		fmt.Println("You answered", answered, "questions in", math.Round(float64(time.Since(start).Seconds())), "seconds")
	}
}
