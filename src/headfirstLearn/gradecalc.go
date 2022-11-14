package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter your Grade")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') //ReadString return 2 value in golang
	fmt.Println(input)
	marks := (strings.TrimSpace(input))
	grade, err := strconv.ParseFloat(marks, 64)
	if err != nil {
		log.Fatal(err)
	} else if grade == 100 {
		fmt.Println("Topper")
	} else if grade >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

}
