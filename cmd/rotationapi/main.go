package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jozefvalverde26/rotationapi/internal/matrix"
	"github.com/jozefvalverde26/rotationapi/internal/rest"
)

func main() {
	matrixService := matrix.NewService()
	restService := rest.NewService(matrixService)
	r := gin.Default()
	r.POST("/matrix/rotate", restService.RetrieveRotatedMatrix)
	r.Run(":5000")
}
