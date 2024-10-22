package linearAlgebraicEquations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranpose(t *testing.T) {
	m := New(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	trans := New(3, 3, [][]float64{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	})

	assert.Equal(t, trans, m.Transpose())
}
