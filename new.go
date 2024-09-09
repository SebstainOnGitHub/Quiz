package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

func newQuestions() {
	f, _ := os.OpenFile("problems.csv", os.O_CREATE | os.O_RDWR | os.O_TRUNC, os.FileMode(0644))
	defer f.Close()
	ops := []string{"+", "-", "*"}
	str := ""

	for i := 0; i < 15; i++ {
		num1 := rand.Int() % 12
		num2 := rand.Int() % 12
		op := rand.Int() % 3

		if num1 == 0 {
			num1 = rand.Int() % 12
		}
		
		if num2 == 0 {
			num2 = rand.Int() % 12
		}
		
		if op == 0 {
			str = fmt.Sprintf("%v %s %v,%v", num1, ops[op], num2, num1 + num2)
		} else if op == 1 {
			str = fmt.Sprintf("%v %s %v,%v", num1, ops[op], num2, num1 - num2)
		} else {
			str = fmt.Sprintf("%v %s %v,%v", num1, ops[op], num2, num1 * num2)
		}
		f.WriteString(str + "\n")
	}
}
