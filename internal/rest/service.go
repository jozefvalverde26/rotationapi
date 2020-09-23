package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jozefvalverde26/rotationapi/internal/domain"
)

//Service represents the rest service
type Service struct {
	matrix domain.Matrix
}

//NewService is the instance of rest service
func NewService(matrix domain.Matrix) Service {
	return Service{matrix}
}

//RetrieveRotatedMatrix retrieve the rotated matrix
func (s Service) RetrieveRotatedMatrix(c *gin.Context) {
	var payload domain.RequestPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid input type"})
		return
	}

	rotatedMatrix, err := s.matrix.Rotate(payload.Matrix)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rotatedMatrix})
}
