package exam

import "fmt"

var exam = variableInit()

func init() {
	fmt.Println("called init in 'exam/exam.go'")
}

func Exam() {
	fmt.Println("called first_exam in 'exam/exam.go'")
}

func variableInit() int {
	fmt.Println("called variableInit in 'exam/exam.go'")
	return 0
}
