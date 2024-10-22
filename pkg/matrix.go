package linearAlgebraicEquations

type Mapper func(val float64, x int, y int) float64

type Folder func(accumulator float64, val float64, x int, y int) float64

type Matrix struct {
	Rows    int
	Columns int
	Data    [][]float64
}

func New(rows, columns int, data [][]float64) Matrix {
	if data == nil {
		data = make([][]float64, rows)
		for i := range data {
			data[i] = make([]float64, columns)
		}
	}

	if len(data) != rows {
		panic("invalid number of rows")
	}

	for _, row := range data {
		if len(row) != columns {
			panic("invalid number of columns")
		}
	}

	return Matrix{
		Rows:    rows,
		Columns: columns,
		Data:    data,
	}
}

func (m Matrix) Sum() float64 {
	return Sum(m)
}

func (m Matrix) Det() float64 {
	return Det(m)
}

func (m Matrix) Transpose() Matrix {
	return Transpose(m)
}

func (m Matrix) addMatrix(n Matrix) Matrix {
	return addMatrix(m, n)
}

func (m Matrix) subMatrix(n Matrix) Matrix {
	return subMatrix(m, n)

}

func (m Matrix) Add(x float64) Matrix {
	return Add(m, x)
}

func (m Matrix) multiplyMatrix(n Matrix) Matrix {
	return multiplyMatrix(m, n)

}

func (m Matrix) Multiply(x float64) Matrix {
	return Multiply(m, x)
}

func (m Matrix) Divide(x float64) Matrix {
	return Divide(m, x)
}
func (m Matrix) Map(f Mapper) Matrix {
	return Map(m, f)
}

func (m Matrix) Fold(f Folder, accumulator float64) float64 {
	return Fold(m, f, accumulator)
}

func (m Matrix) Norm(typeNorm float64) float64 {
	return Norm(m, typeNorm)
}

func (m Matrix) EM() Matrix {
	return EM(m.Rows)
}

func (m Matrix) diagonalDominance() bool {
	return diagonalDominance(m)
}
