package main

import "fmt"

func abc(channel chan string) {
	channel <- "A"
	channel <- "B"
	channel <- "C"

}
func bcd(channel chan string) {
	channel <- "F"

	channel <- "Q"

	channel <- "J"
	//close(channel) //Close the channel
	//channel <- "K"

}
func main() {
	channel1 := make(chan string, 1) //Buffered Channel
	channel2 := make(chan string)
	go abc(channel1)
	go bcd(channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)

}
