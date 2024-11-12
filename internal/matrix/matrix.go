package matrix

import (
	"fmt"
	"math"
)

const (
	MatrixSize  = 15
	FirstXValue = 9
)

func RealX() []float64 {
	xs := make([]float64, MatrixSize)
	for i := range MatrixSize {
		xs[i] = float64(FirstXValue + i)
	}

	return xs
}

func GenerateMatrixAndRHS() ([][]float64, []float64) {
	mtr, rhs, xs := make([][]float64, MatrixSize), make([]float64, MatrixSize), RealX()

	for i := range MatrixSize {
		mtr[i] = make([]float64, MatrixSize)
		curRHS := 0.0

		for j := range MatrixSize {
			if i == j {
				mtr[i][j] = 11.0 * math.Sqrt(float64(i+1))
			} else {
				mtr[i][j] = 0.001 * float64(i+1) / float64(j+1)
			}

			curRHS += mtr[i][j] * xs[j]
		}

		rhs[i] = curRHS
	}

	return mtr, rhs
}

func CubicNorm(vector []float64) float64 {
	norm := 0.0

	for _, v := range vector {
		norm = max(norm, math.Abs(v))
	}

	return norm
}

func Diff(first, second []float64) float64 {
	norm := 0.0

	for i := range len(first) {
		norm = max(norm, math.Abs(first[i]-second[i]))
	}

	return norm
}

func CalcError(xs []float64) float64 {
	return Diff(RealX(), xs)
}

func RelativeCalcError(xs []float64) float64 {
	realX := RealX()
	norm := 0.0

	for i := range len(realX) {
		norm = max(norm, math.Abs(realX[i]-xs[i]))
	}

	return norm / CubicNorm(realX)
}

func CopyMatrix(matrix []float64) []float64 {
	matrixCopy := make([]float64, len(matrix))
	copy(matrixCopy, matrix)

	return matrixCopy
}

func Copy2DMatrix(matrix [][]float64) [][]float64 {
	matrixCopy := make([][]float64, len(matrix))

	for i := range len(matrix) {
		matrixCopy[i] = CopyMatrix(matrix[i])
	}

	return matrixCopy
}

func PrintMatrix(matrix []float64) {
	for i := range len(matrix) {
		fmt.Printf("%10.5f ", matrix[i])
	}

	fmt.Println()
}

func Print2DMatrix(matrix [][]float64) {
	for i := range len(matrix) {
		PrintMatrix(matrix[i])
	}
}