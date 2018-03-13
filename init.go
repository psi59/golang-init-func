package main

import (
	"fmt"

	"github.com/tkddlf59/init_func/exam"
)

var v = variableInit()

func main() {
	fmt.Println("called main")
	exam.Exam()
}

func init() {
	fmt.Println("called init")
}

func variableInit() int {
	fmt.Println("called variableInit")
	return 0
}
