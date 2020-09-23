package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	tests := map[string]struct {
		matrix                [][]int
		expectedRotatedMatrix [][]int
	}{
		"should return rotated matrix when given a valid matrix": {
			matrix:                [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expectedRotatedMatrix: [][]int{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}},
		},
		"with return error when given an invalid matrix": {
			matrix:                [][]int{{1, 2, 3}, {4, 6}, {7, 8, 9}},
			expectedRotatedMatrix: nil,
		},
	}

	for name, test := range tests {
		matrixService := NewService()
		t.Run(name, func(t *testing.T) {
			rotatedMatrix, err := matrixService.Rotate(test.matrix)
			if err != nil {
				assert.Equal(t, err.Error(), "index out of range, must be NxN matrix")
			} else {
				assert.Equal(t, rotatedMatrix, test.expectedRotatedMatrix)
			}
		})
	}
}
