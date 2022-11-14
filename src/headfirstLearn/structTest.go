package main

import "fmt"

type studentDetails struct {
	name  string
	roll  int
	marks float64
}

func main() {
	var detail studentDetails
	detail.name = "Yusuf"
	detail.roll = 16
	detail.marks = 88.90
	showStudentDetails(detail)
	a := partdetail("Danish Ahmad")
	fmt.Printf("The name is %v and his roll is %v and marks %v", a.name, a.roll, a.marks)
}
func showStudentDetails(a studentDetails) {
	fmt.Printf("The name of the student is %v and roll number is %v and marks is %v", a.name, a.roll, a.marks)
}
func partdetail(nameData string) studentDetails {
	var nameDetail studentDetails
	nameDetail.name = "Danish Ahmad"
	nameDetail.roll = 36
	nameDetail.marks = 44.7
	return nameDetail
}


