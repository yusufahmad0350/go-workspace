package main

import (
	"fmt"
	"runtime"
)

func echo(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // allow switching between the scheduled routines
		fmt.Println(s)
	}
}

func main() {

	// maximize CPU usage for multi threading
	runtime.GOMAXPROCS(runtime.NumCPU())

	go echo("bye") // schedule this execution on a new thread [second]

	go echo("world") // [third]

	echo("good") // current [first thread]

	fmt.Println("# go routines : ", runtime.NumGoroutine())
}
