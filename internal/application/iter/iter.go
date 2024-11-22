package iter

import (
	"errors"
	"fmt"

	"matrix-iter/internal/iter"
	"matrix-iter/internal/matrix"
)

func runMethod(method func() ([]float64, int, error)) error {
	xs, iterCnt, err := method()
	if err != nil {
		if errors.As(err, &iter.ErrMaxIters{}) {
			fmt.Println("Maximum number of iterations reached!")
		} else {
			return fmt.Errorf("method Jacobi: %w", err)
		}
	}

	fmt.Println("Vector x:")
	matrix.PrintMatrix(xs)
	fmt.Println()

	fmt.Printf("Resulting number of operations: %d\n", iterCnt)
	fmt.Printf("Calculation error: %e\n", matrix.CalcError(xs))
	fmt.Printf("Relative calculation error: %e\n", matrix.RelativeCalcError(xs))

	return nil
}

func runIters(mtr [][]float64, rhs []float64) error {
	fmt.Println("Initial matrix A:")
	matrix.Print2DMatrix(mtr)
	fmt.Println()

	fmt.Println("Initial right-hand side vector b:")
	matrix.PrintMatrix(rhs)
	fmt.Println()

	if ok := matrix.HasDiagonalDominance(mtr); ok {
		fmt.Println("Matrix has diagonal dominance")
	} else {
		fmt.Println("Matrix hasn't diagonal dominance")
	}

	fmt.Println("--- Method Jacobi")
	if err := runMethod(func() ([]float64, int, error) {
		return iter.Jacobi(mtr, rhs)
	}); err != nil {
		return fmt.Errorf("runMethod(Jacobi): %w", err)
	}

	fmt.Println()

	fmt.Println("--- Method GaussSeidel")
	if err := runMethod(func() ([]float64, int, error) {
		return iter.GaussSeidel(mtr, rhs)
	}); err != nil {
		return fmt.Errorf("runMethod(GaussSeidel): %w", err)
	}

	fmt.Println()

	fmt.Println("--- Method of successive over-relaxation with w=0.5")
	if err := runMethod(func() ([]float64, int, error) {
		return iter.SOR(mtr, rhs, 0.5)
	}); err != nil {
		return fmt.Errorf("runMethod(SOR, w=0.5): %w", err)
	}

	fmt.Println()

	fmt.Println("--- Method of successive over-relaxation with w=1.5")
	if err := runMethod(func() ([]float64, int, error) {
		return iter.SOR(mtr, rhs, 1.5)
	}); err != nil {
		return fmt.Errorf("runMethod(SOR, w=1.5): %w", err)
	}

	return nil
}

func Start() error {
	fmt.Println("Iteration for matrix with diagonal dominance")

	mtr, rhs := matrix.GenerateDiagDominanceMatrixAndRHS()
	if err := runIters(mtr, rhs); err != nil {
		return fmt.Errorf("runIters: %w", err)
	}

	fmt.Print("\n\n-----------------------------------------------\n\n")
	fmt.Println("Iteration for matrix without diagonal dominance")

	mtr, rhs = matrix.GenerateNonDiagDominanceMatrixAndRHS()
	if err := runIters(mtr, rhs); err != nil {
		return fmt.Errorf("runIters: %w", err)
	}

	return nil
}
