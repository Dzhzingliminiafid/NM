package linearAlgebraicEquations

import "math"

func Map(m Matrix, f Mapper) Matrix {
	n := New(m.Rows, m.Columns, nil)

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			val := m.Data[i][j]
			n.Data[i][j] = f(val, i, j)
		}
	}
	return n
}

// Накапливаем значения в матрице в соответствии с функцией f
// accumulator - накапливаемое значение
func Fold(m Matrix, f Folder, accumulator float64) float64 {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			accumulator = f(accumulator, m.Data[i][j], i, j)
		}
	}

	return accumulator
}

func Sum(m Matrix) float64 {
	return Fold(m, func(accumulator float64, val float64, x int, y int) float64 {
		return accumulator + val
	}, 0)
}

func Divide(m Matrix, a float64) Matrix {
	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] / a
	})
}

func Add(m Matrix, a float64) Matrix {
	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] + a
	})
}

func addMatrix(m, n Matrix) Matrix {
	if m.Rows != n.Rows || m.Columns != n.Columns {
		panic("mismatched different sizes")
	}

	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] + n.Data[x][y]
	})
}

func subMatrix(m, n Matrix) Matrix {
	if m.Rows != n.Rows || m.Columns != n.Columns {
		panic("mismatched different sizes")
	}

	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] - n.Data[x][y]
	})
}

func Multiply(m Matrix, a float64) Matrix {
	return Map(New(m.Rows, m.Columns, nil), func(val float64, x, y int) float64 {
		return m.Data[x][y] * a
	})
}

func multiplyMatrix(m, n Matrix) Matrix {
	if m.Columns != n.Rows {
		panic("mismatched different sizes")
	}

	return Map(New(m.Rows, n.Columns, nil), func(_ float64, x, y int) float64 {
		sum := 0.0
		for i := 0; i < n.Rows; i++ {
			sum += m.Data[x][i] * n.Data[i][y]
		}
		return sum
	})
}

func Det(m Matrix) float64 {
	if m.Rows == 1 && m.Columns == 1 {
		return m.Data[0][0]
	}
	// Remove 1st column and n-th row
	f := func(m Matrix, n int) Matrix {
		data := [][]float64{}

		for i, row := range m.Data {
			if i == 0 {
				continue
			}
			current := []float64{}

			for j, col := range row {
				if j == n {
					continue
				}
				current = append(current, col)
			}
			data = append(data, current)
		}

		return New(m.Rows-1, m.Columns-1, data)
	}
	det := 0.0
	for n, v := range m.Data[0] {
		det += v * math.Pow(-1, float64(n)) * Det(f(m, n))
	}
	return det
}

func Transpose(m Matrix) Matrix {
	return Map(New(m.Columns, m.Rows, nil), func(val float64, x, y int) float64 {
		return m.Data[y][x]
	})
}

func Norm(m Matrix, typeNorm float64) float64 {
	if typeNorm == 1.0 {
		norm := 0.0

		for i := 0; i < m.Columns; i++ {
			columnSum := 0.0
			for j := 0; j < m.Rows; j++ {
				columnSum += math.Abs(m.Data[i][j])
			}
			norm = math.Max(norm, columnSum)
		}
		return norm
	}

	if typeNorm == math.Inf(1) {
		norm := 0.0

		for i := 0; i < m.Rows; i++ {
			rowSum := 0.0
			for j := 0; j < m.Columns; j++ {
				rowSum += math.Abs(m.Data[i][j])
			}
			norm = math.Max(norm, rowSum)
		}
		return norm
	}
	return 0.0
}

func EM(rows int) Matrix {
	m := New(rows, rows, nil)

	for i := 0; i < rows; i++ {
		m.Data[i][i] = 1.0
	}

	return m
}

func diagonalDominance(m Matrix) bool {
	for i := 0; i < m.Rows; i++ {
		sum := 0.0
		for j := 0; j < m.Columns; j++ {
			if i != j {
				sum += math.Abs(m.Data[i][j])
			}
		}
		if math.Abs(m.Data[i][i]) < sum {
			return false
		}
	}
	return true
}
