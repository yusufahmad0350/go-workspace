package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	second := time.Now().Unix()
	rand.Seed(second)
	randNum := rand.Intn(100)
	count := 0
	for i := 0; i < 10; i++ {
		count++

		fmt.Println("Enter a choice of your number")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		userInput, err := strconv.ParseFloat(input, 64)
		fmt.Println(userInput, randNum)
		if err != nil {
			println("Error found: ", err)
		}
		if int(userInput) == randNum {
			fmt.Println("You guessed Right, Congratulations")
			break
		}

		if int(userInput) < randNum {
			fmt.Println("OOPs, Your guess is low, please try again")
		} else {
			fmt.Println("OOPs, Your guess is High, please try again")
		}
		println("Your Remaining trying capacity is ", 10-count)
		if count == 10 {
			println("Sorry you didnt guessed right, The correct number was", randNum)
			break
		}

	}

}
