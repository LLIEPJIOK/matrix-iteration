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

func Start() error {
	mtr, rhs := matrix.GenerateMatrixAndRHS()

	fmt.Println("Initial matrix A:")
	matrix.Print2DMatrix(mtr)
	fmt.Println()

	fmt.Println("Initial right-hand side vector b:")
	matrix.PrintMatrix(rhs)
	fmt.Println()

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
