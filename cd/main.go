package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"linearAlgebraicEquations/pkg"
)

func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("№", "x_", "eps", "МПИ→", "x", "Δx", "k", "Зейдель→", "x", "Δx", "k")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	tests := []struct {
		A linearAlgebraicEquations.Matrix
		b linearAlgebraicEquations.Matrix
		x linearAlgebraicEquations.Matrix
	}{
		{
			A: linearAlgebraicEquations.New(3, 3, [][]float64{
				{0, 2, 3},
				{1, 2, 4},
				{4, 5, 6},
			}),
			b: linearAlgebraicEquations.New(3, 1, [][]float64{
				{13},
				{17},
				{32},
			}),
			x: linearAlgebraicEquations.New(3, 1, [][]float64{
				{1},
				{2},
				{3},
			}),
		},
	}

	eps := 1e-3

	N := 0

	for _, test := range tests {
		A := test.A
		b := test.b
		x_ := test.x
		for i := 0; i < 3; i++ {
			x1, k1 := linearAlgebraicEquations.SIM(A, b, eps)
			x2, k2 := linearAlgebraicEquations.Seidel(A, b, eps)
			for j := 0; j < x1.Rows; j++ {

				if j == x1.Rows/2 && i == 1 {
					tbl.AddRow(N, x_.Data[j][0], eps, " ", fmt.Sprintf("%.6f", x1.Data[j][0]), fmt.Sprintf("%.6e", x1.Data[j][0]-x_.Data[j][0]), k1, " ", fmt.Sprintf("%.6f", x2.Data[j][0]), fmt.Sprintf("%.6e", x2.Data[j][0]-x_.Data[j][0]), k2)
				} else if j == x1.Rows/2 {
					tbl.AddRow(" ", x_.Data[j][0], eps, " ", fmt.Sprintf("%.6f", x1.Data[j][0]), fmt.Sprintf("%.6e", x1.Data[j][0]-x_.Data[j][0]), k1, " ", fmt.Sprintf("%.6f", x2.Data[j][0]), fmt.Sprintf("%.6e", x2.Data[j][0]-x_.Data[j][0]), k2)
				} else {
					tbl.AddRow(" ", x_.Data[j][0], " ", " ", fmt.Sprintf("%.6f", x1.Data[j][0]), fmt.Sprintf("%.6e", x1.Data[j][0]-x_.Data[j][0]), " ", " ", fmt.Sprintf("%.6f", x2.Data[j][0]), fmt.Sprintf("%.6e", x2.Data[j][0]-x_.Data[j][0]), " ")
				}
			}
			tbl.AddRow(" ", " ", " ", " ", " ", " ", " ", " ", " ", " ")
			eps *= 1e-1
		}
		tbl.AddRow(" ", " ", " ", " ", " ", " ", " ", " ", " ", " ")
		N++
	}
	fmt.Println("Table:")
	tbl.Print()
}
