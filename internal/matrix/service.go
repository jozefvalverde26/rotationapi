package matrix

import "errors"

//Service represents the matrix service
type Service struct{}

//NewService is the instance of matrix service
func NewService() Service {
	return Service{}
}

//Rotate rotate the matrix
func (s Service) Rotate(matrix [][]int) ([][]int, error) {
	nRows := len(matrix)
	for _, colums := range matrix {
		if nRows != len(colums) {
			return nil, errors.New("index out of range, must be NxN matrix")
		}
	}
	for x := 0; x < nRows/2; x++ {
		for y := x; y < nRows-x-1; y++ {
			temp := matrix[x][y]
			matrix[x][y] = matrix[y][nRows-1-x]
			matrix[y][nRows-1-x] = matrix[nRows-1-x][nRows-1-y]
			matrix[nRows-1-x][nRows-1-y] = matrix[nRows-1-y][x]
			matrix[nRows-1-y][x] = temp
		}
	}
	return matrix, nil
}
