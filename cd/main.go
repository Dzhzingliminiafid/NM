package main

import (
	"fmt"
	"linearAlgebraicEquations/pkg"
)

func main() {

	A := linearAlgebraicEquations.New(3, 3, [][]float64{
		{3, 1, 1},
		{1, 5, 1},
		{1, 1, 7},
	})
	b := linearAlgebraicEquations.New(3, 1, [][]float64{
		{5}, {7}, {9},
	})

	eps := 1e-3

	x, iterations := linearAlgebraicEquations.SIM(A, b, eps)
	fmt.Println("Решение:", x)
	fmt.Println("Количество итераций:", iterations)

	x, iterations = linearAlgebraicEquations.Seidel(A, b, eps)
	fmt.Println("Решение:", x)
	fmt.Println("Количество итераций:", iterations)
}
