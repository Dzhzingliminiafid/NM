package linearAlgebraicEquations

import (
	"math"
)

// SIM - Simple Iteration Method
func SIM(A Matrix, b Matrix, eps float64) (Matrix, int) {
	n := A.Rows
	E := EM(n)

	mu := 1 / Norm(A, math.Inf(1))
	B := subMatrix(E, Multiply(A, mu)) // B = E - A*mu

	BNorm := Norm(B, math.Inf(1))
	if BNorm >= 1 { // Если не выполняется д.у.
		T := Transpose(A)
		A = multiplyMatrix(T, A)
		b = multiplyMatrix(T, b)

		mu = 1.0 / Norm(A, math.Inf(1))
		B = subMatrix(E, Multiply(A, mu))

		BNorm = Norm(B, math.Inf(1))
	}

	c := Multiply(b, mu)
	x := c

	k := 0

	for {
		k++

		x_ := addMatrix(multiplyMatrix(B, x), c)

		if BNorm < 1 && BNorm/(1-BNorm)*(x_.subMatrix(x)).Norm(math.Inf(1)) < eps {
			return x_, k
		} else if BNorm >= 1 && multiplyMatrix(A, x_).subMatrix(b).Norm(math.Inf(1)) < eps {
			return x_, k
		}

		x = x_

	}
}
