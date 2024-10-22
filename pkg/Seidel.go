package linearAlgebraicEquations

import "math"

func Seidel(A Matrix, b Matrix, eps float64) (Matrix, int) {
	if diagonalDominance(A) == false {
		panic("Матрица не c диаганальным преобладанием")
	}
	C := New(A.Rows, A.Columns, nil)
	d := New(b.Rows, b.Columns, nil)
	for i := 0; i < A.Rows; i++ {
		for j := 0; j < A.Columns; j++ {
			if i == j {
				C.Data[i][j] = 0
			} else {
				C.Data[i][j] = -(A.Data[i][j] / A.Data[i][i])
			}
		}
		d.Data[i][0] = b.Data[i][0] / A.Data[i][i]
	}

	x0 := d

	k := 0

	for {
		k++
		x_ := New(x0.Rows, x0.Columns, nil)
		for i := 0; i < C.Rows; i++ {
			x_.Data[i][0] = 0
			for j := 0; j < C.Columns; j++ {
				if j < i {
					x_.Data[i][0] += C.Data[i][j] * x_.Data[j][0]
				} else if i != j {
					x_.Data[i][0] += C.Data[i][j] * x0.Data[j][0]
				}
			}
			x_.Data[i][0] += d.Data[i][0]
		}
		if multiplyMatrix(A, x_).subMatrix(b).Norm(math.Inf(1)) < eps {
			return x_, k
		}
		x0 = x_
	}

}
