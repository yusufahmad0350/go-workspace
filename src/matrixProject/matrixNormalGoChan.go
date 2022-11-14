package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const lenOfMat = 256

var start time.Time
var mat [lenOfMat][lenOfMat]int

// toMultiply will hold details of what the goroutine will be multiplying (one rows and one columns)
type toMultiply struct {
	rowsNumber    int
	columnsNumber int
	rows          []int
	columns       []int
}

func main() {
	const noOfGoRoutines = 5

	// Build up a matrix of dimensions (lenOfMat) x (lenOfMat)
	var a [lenOfMat][lenOfMat]int
	var b [lenOfMat][lenOfMat]int
	for i := 0; i < lenOfMat; i++ {
		for j := 0; j < lenOfMat; j++ {
			a[i][j] = rand.Intn(10)
			b[i][j] = rand.Intn(10)
		}
	}
	//fmt.Println("The first generated matrix is", a)
	//fmt.Println("The Second generated matrix is", b)

	// Setup completed so start the clock...
	start = time.Now()

	// Start off threadlenOfMat go routines to multiply each rows/columns
	toCalc := make(chan toMultiply)
	var wg sync.WaitGroup
	wg.Add(noOfGoRoutines)
	fmt.Printf("Number of CPUs Total is %d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(4) //
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))
	for i := 0; i < noOfGoRoutines; i++ {
		go func() {
			Calc(toCalc)
			wg.Done()
		}()
	}

	// Begin the multiplication.
	start = time.Now()
	for i := 0; i < lenOfMat; i++ {
		for j := 0; j < lenOfMat; j++ {
			tm := toMultiply{
				rowsNumber:    i,
				columnsNumber: j,
				rows:          make([]int, lenOfMat),
				columns:       make([]int, lenOfMat),
			}

			for k := 0; k < lenOfMat; k++ {
				tm.rows[k] = a[i][j]
				tm.columns[k] = b[i][k]

			}
			toCalc <- tm
		}
	}

	// All of the data has been sent to the chanel; now we need to wait for all of the
	// goroutines to complete
	close(toCalc)
	wg.Wait()

	fmt.Println("Total time for completion is : ", time.Since(start))

	// The full result should be in tz
	for i := 0; i < lenOfMat; i++ {
		for j := 0; j < lenOfMat; j++ {
			//fmt.Print(mat[i][j])
			//fmt.Print(" ")
		}
		//fmt.Println(" ")
	}
	//fmt.Print(mat)
}

// Calc - Multiply a rows from one matrix with a columns from another
func Calc(toCalc <-chan toMultiply) {
	for tc := range toCalc {
		var result int
		for i := 0; i < len(tc.rows); i++ {
			result += tc.rows[i] * tc.columns[i]
		}
		// warning - the below should work in this case but be careful writing to global variables from goroutines
		mat[tc.rowsNumber][tc.columnsNumber] = result

	}

}
