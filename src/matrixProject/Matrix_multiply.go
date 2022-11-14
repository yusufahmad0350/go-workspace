package main

import (
	"fmt"
)

var matrixA = [][]int{{2, 3}, {9, 8}}
var matrixB = [][]int{{3, 5}, {2, 7}}

func main() {
	//n:=len(matrixA)
	//m:=len(matrixB)

	var C1 = make([][]int, len(matrixA))
	var A1 = make([][]int, len(matrixA))
	var B1 = make([][]int, len(matrixB))
	var UpdatedMatA = make([][]int, len(matrixB))
	var UpdatedMatB = make([][]int, len(matrixB))
	var C2 = make([][]int, len(matrixA))

	for i := 0; i < len(matrixA); i++ {

		if i == 0 {
			C1, B1, A1 = first_operation(matrixA, matrixB)
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

}
func rotate_first_incremental_matrixA(mat [][]int) [][]int {
	k := 0
	for i := 0; i < len(mat); i++ {
		mat1 := mat[i]
		//fmt.Print(mat1)
		if i != 0 {
			k++ //everytime +1 increment to iterative rows
			k = k % len(mat1)
			shifted_rows := append(mat1[k:], mat1[0:k]...)
			copy(mat1, shifted_rows)
			fmt.Print("mat is", mat1)

		}

	}
	fmt.Print("mat is", mat)
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
