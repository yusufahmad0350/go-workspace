package main

import (
	"fmt"
	"mpi"
	"os"
)

func main() {
	mpi.Init(&os.Args)
	worldSize, _ := mpi.Comm_size(mpi.COMM_WORLD)
	rank, _ := mpi.Comm_rank(mpi.COMM_WORLD)
	procName, _ := mpi.Get_processor_name()
	fmt.Printf("Hello world from processor %s, rank %d out of %d processors\n", procName, rank, worldSize)
	mpi.Finalize()

	var matrixA = [][]int{{1, 2}, {3, 4}}
	var matrixB = [][]int{{1, 2}, {3, 4}}
	var matrixC = make([][]int, 2)

	for i := 0; i < 2; i++ {
		matrixC[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				matrixC[i][j] += matrixA[i][k] * matrixB[k][j]
			}

		}

	}
	fmt.Print(matrixC)
}
