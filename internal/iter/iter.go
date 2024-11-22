package iter

import (
	"fmt"

	"matrix-iter/internal/matrix"
)

const (
	maxIters = 100
	eps      = 1e-5
)

func validate(matrix [][]float64, rhs []float64) error {
	if len(matrix) == 0 {
		return NewErrMatrix("matrix is empty")
	}

	n := len(matrix)

	for i := range n {
		if len(matrix[i]) != n {
			return NewErrMatrix("matrix is not square")
		}
	}

	if len(rhs) != n {
		return NewErrRHS("right-hand side vector must have the same size as matrix")
	}

	return nil
}

func Jacobi(mtr [][]float64, rhs []float64) ([]float64, int, error) {
	mtr, rhs = matrix.Copy2DMatrix(mtr), matrix.CopyMatrix(rhs)

	if err := validate(mtr, rhs); err != nil {
		return nil, 0, fmt.Errorf("invalid input params: %w", err)
	}

	n := len(mtr)

	// Ax=b -> x = A'x+b'
	for i := range n {
		diagEl := mtr[i][i]

		for j := range n {
			if i != j {
				mtr[i][j] /= -diagEl
			} else {
				mtr[i][j] = 0
			}
		}

		rhs[i] /= diagEl
	}

	xi := matrix.CopyMatrix(rhs)
	iterCnt := 0

	// iterations
	for range maxIters {
		nx := make([]float64, n)

		for i := range n {
			for j := range n {
				nx[i] += mtr[i][j] * xi[j]
			}

			nx[i] += rhs[i]
		}

		if matrix.Diff(nx, xi) < eps {
			return xi, iterCnt, nil
		}

		xi = nx
		iterCnt++
	}

	return xi, iterCnt, NewErrMaxIters()
}

func SOR(mtr [][]float64, rhs []float64, w float64) ([]float64, int, error) {
	if err := validate(mtr, rhs); err != nil {
		return nil, 0, fmt.Errorf("invalid input params: %w", err)
	}

	n := len(mtr)

	xi := matrix.CopyMatrix(rhs)
	iterCnt := 0

	// iterations
	for range maxIters {
		nx := make([]float64, n)

		for i := range n {
			sum := 0.0
			for j := range n {
				if j < i {
					sum += mtr[i][j] * nx[j]
				} else if j > i {
					sum += mtr[i][j] * xi[j]
				}
			}

			nx[i] = (1-w)*xi[i] + w/mtr[i][i]*(rhs[i]-sum)
		}

		if matrix.Diff(nx, xi) < eps {
			return xi, iterCnt, nil
		}

		xi = nx
		iterCnt++
	}

	return xi, iterCnt, NewErrMaxIters()
}

const gaussSeidelW = 1.0

func GaussSeidel(mtr [][]float64, rhs []float64) ([]float64, int, error) {
	xs, iterCnt, err := SOR(mtr, rhs, gaussSeidelW)
	if err != nil {
		return xs, iterCnt, fmt.Errorf("SOR with w=%f: %w", gaussSeidelW, err)
	}

	return xs, iterCnt, nil
}
