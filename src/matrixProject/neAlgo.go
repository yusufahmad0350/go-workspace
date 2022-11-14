package main

import (
	"fmt"
	"video/gonum.org/v1/gonum/mat"
)

// This example shows how simple user types can be constructed to
// implement basic vector functionality within the mat package.
func main() {
	// Perform the cross product of [1 2 3 4] and [1 2 3].
	r := row{1, 2, 3, 4}
	c := column{1, 2, 3}

	var m mat.Dense
	m.Mul(c, r)

	fmt.Println(mat.Formatted(&m))

}

// row is a user-defined row vector.
type row []float64

// Dims, At and T minimally satisfy the mat.Matrix interface.
func (v row) Dims() (r, c int)    { return 1, len(v) }
func (v row) At(_, j int) float64 { return v[j] }
func (v row) T() mat.Matrix       { return column(v) }

// RawVector allows fast path computation with the vector.
func (v row) RawVector() blas64.Vector {
	return blas64.Vector{N: len(v), Data: v, Inc: 1}
}

// column is a user-defined column vector.
type column []float64

// Dims, At and T minimally satisfy the mat.Matrix interface.
func (v column) Dims() (r, c int)    { return len(v), 1 }
func (v column) At(i, _ int) float64 { return v[i] }
func (v column) T() mat.Matrix       { return row(v) }

// RawVector allows fast path computation with the vector.
func (v column) RawVector() blas64.Vector {
	return blas64.Vector{N: len(v), Data: v, Inc: 1}
}
