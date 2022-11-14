package main

import "fmt"

var matrixA = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

//var matrixB = [][]int{{56},{1,2,3,4},{1,2,3,4},{1,2,3,4}}

func main() {
	fmt.Print(matrixA)
	rotate_first_incremental_matrix(matrixA)
	//shift_one_row()
	//shift_two_row()
	//shift_three_row()
	fmt.Print(matrixA)
}

func shift_one_row() {

	firstrows := matrixA[1]
	//Perform left shift
	Rotate(firstrows, 1)
	fmt.Println(firstrows)

}
func shift_two_row() {
	second_rows := matrixA[2]
	//Perform left shift
	Rotate(second_rows, 2)
	fmt.Println(second_rows)

}
func shift_three_row() {
	third_rows := matrixA[3]
	//Perform left shift
	Rotate(third_rows, 3)
	fmt.Println(third_rows)
}
func Rotate(mat []int, k int) {
	k = k % len(mat)
	shifted_rows := append(mat[k:], mat[0:k]...)
	copy(mat, shifted_rows) // this actually writes to where nums points to
}

func rotate_first_incremental_matrix(mat [][]int) {
	k := 0
	for i := 0; i < len(mat); i++ {
		//mat i is first row
		mat1 := mat[i]
		if i == 0 {
			fmt.Print("Fist row skipped")
		} else {

			k++
			k = k % len(mat1)
			shifted_rows := append(mat1[k:], mat1[0:k]...)
			copy(mat1, shifted_rows)

		}

	}

}
