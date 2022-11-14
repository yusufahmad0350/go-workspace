package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	// we are seeding the rand variable with present time
	// so that we would get different output each time
}
func main() {
	var col = 1000
	var row = 1000
	var randMatrixA [][]int
	randMatrixA = genMat(row, col)
	fmt.Println(randMatrixA)
	var randMatrixB [][]int
	randMatrixB = genMat(col, row)
	fmt.Println(randMatrixB)
	fmt.Println("The Go Result of Matrix Multiplication = ")
	start := time.Now()
	c := doCalc(randMatrixA, randMatrixB)
	elapsed := time.Since(start)

	fmt.Println(c)

	fmt.Printf("Time taken to calculate %s \n", elapsed)
}

func genMat(row int, col int) [][]int {
	nM := make([][]int, col)
	for i := 0; i < col; i++ {
		nM[i] = make([]int, row)
		// we are creating a slice which can hold type int
	}
	generateNums(nM)
	return nM
}
func generateNums(randMatrix [][]int) {
	for i, innerArray := range randMatrix {
		for j := range innerArray {
			randMatrix[i][j] = rand.Intn(100)
			//looping over each element of array and assigning it a random variable
		}
	}
}
func rowCount(inM [][]int) int {
	return (len(inM))
}
func colCount(inM [][]int) int {
	return (len(inM[0]))
}
func doCalc(inA [][]int, inB [][]int) [][]int {
	var C1 = make([][]int, len(inA))
	var A1 = make([][]int, len(inA))
	var B1 = make([][]int, len(inB))
	var UpdatedMatA = make([][]int, len(inB))
	var UpdatedMatB = make([][]int, len(inB))
	var C2 = make([][]int, len(inA))

	for i := 0; i < len(inA); i++ {

		if i == 0 {
			C1, B1, A1 = first_operation(inA, inB)
			//First Calculation
			fmt.Println(A1)
			fmt.Println(B1)
			fmt.Println("The first Calculation \n", C1)

		} else {
			UpdatedMatA, UpdatedMatB, C2 = other_step_operation(A1, B1)
			copy(A1, UpdatedMatA)
			copy(B1, UpdatedMatB)

			fmt.Println("The next step \n", C2)
			func(C1, C2 [][]int) {
				for i := 0; i < len(C1); i++ {
					for j := 0; j < len(C1); j++ {
						C1[i][j] = C1[i][j] + C2[i][j]

					}

				}

			}(C1, C2)

		}

	}
	fmt.Print("The Final Result for this Cannon Algorith is \n \n \n ", C1)
	return C1
}

func rotate_first_incremental_matrixA(mat [][]int) [][]int {
	k := 0
	for i := 0; i < len(mat); i++ {
		mat1 := mat[i]
		if i == 0 {
			//fmt.Print("Fist row skipped")
		} else {

			k++ //everytime +1 increment to iterative rows
			k = k % len(mat1)
			shifted_rows := append(mat1[k:], mat1[0:k]...)
			copy(mat1, shifted_rows)

		}

	}
	return mat

}

func rotate_second_fixed_matrixA(mat [][]int) [][]int {

	k := 1
	for i := 0; i < len(mat); i++ {
		mat1 := mat[i]
		k = k % len(mat1)
		shifted_rows := append(mat1[k:], mat1[0:k]...)
		copy(mat1, shifted_rows)

	}
	return mat

}
func first_operation(matx [][]int, maty [][]int) ([][]int, [][]int, [][]int) {
	var final_matrixC1 = make([][]int, len(matx))
	rotate_first_incremental_matrixA(matx)
	matrixD := column_to_row_matrix(maty)
	rotate_first_incremental_matrixA(matrixD)
	matrixE := column_to_row_matrix(matrixD)
	for i := 0; i < len(matx); i++ {
		final_matrixC1[i] = make([]int, len(matx[0]))
		for j := 0; j < len(maty); j++ {
			final_matrixC1[i][j] = matx[i][j] * matrixE[i][j]
		}

	}
	return final_matrixC1, matrixE, matx

}
func other_step_operation(matx [][]int, maty [][]int) ([][]int, [][]int, [][]int) {
	var final_matrixC1 = make([][]int, len(matx))
	rotate_second_fixed_matrixA(matx)
	matrixD := column_to_row_matrix(maty)
	rotate_second_fixed_matrixA(matrixD)
	matrixE := column_to_row_matrix(matrixD)
	for i := 0; i < len(matx); i++ {
		final_matrixC1[i] = make([]int, len(matx[0]))
		for j := 0; j < len(maty); j++ {
			final_matrixC1[i][j] = matx[i][j] * matrixE[i][j]
		}

	}
	return matx, matrixE, final_matrixC1

}

func column_to_row_matrix(mat [][]int) [][]int {

	var column_row_matrix = make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		column_row_matrix[i] = make([]int, len(mat[0]))
		for j := 0; j < len(mat); j++ {
			column_row_matrix[i][j] = mat[j][i]

		}
	}
	return column_row_matrix
}
