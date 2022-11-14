package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const length = 1000

var start time.Time
var rez [length][length]int

// toMultiply will hold details of what the goroutine will be multiplying (one row and one column)
type toMultiply struct {
	rowNo    int
	columnNo int
	row      []int
	column   []int
}

func main() {
	const noOfGoRoutines = 5

	// Build up a matrix of dimensions (length) x (length)
	var a [length][length]int
	var b [length][length]int
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			a[i][j] = rand.Intn(10)
			b[i][j] = rand.Intn(10)
		}
	}

	// Setup completed so start the clock...
	start = time.Now()

	// Start off threadlength go routines to multiply each row/column
	toCalc := make(chan toMultiply)
	var wg sync.WaitGroup
	wg.Add(noOfGoRoutines)
	for i := 0; i < noOfGoRoutines; i++ {
		go func() {
			Calc(toCalc)
			wg.Done()
		}()
	}

	// Begin the multiplication.
	start = time.Now()
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			tm := toMultiply{
				rowNo:    i,
				columnNo: j,
				row:      make([]int, length),
				column:   make([]int, length),
			}

			for k := 0; k < length; k++ {
				tm.row[k] = a[i][j]
				tm.column[k] = b[i][k]
			}
			toCalc <- tm
		}
	}

	// All of the data has been sent to the chanel; now we need to wait for all of the
	// goroutines to complete
	close(toCalc)
	wg.Wait()

	fmt.Println("Binomial took ", time.Since(start))

	// The full result should be in tz
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			//fmt.Print(rez[i][j])
			//fmt.Print(" ")
		}
		//fmt.Println(" ")
	}
}

// Calc - Multiply a row from one matrix with a column from another
func Calc(toCalc <-chan toMultiply) {
	for tc := range toCalc {
		var result int
		for i := 0; i < len(tc.row); i++ {
			result += tc.row[i] * tc.column[i]
		}
		// warning - the below should work in this case but be careful writing to global variables from goroutines
		rez[tc.rowNo][tc.columnNo] = result
	}
}
