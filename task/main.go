package task

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func Solve(c *gin.Context) error {
	err := errors.New("good")

	return err
}
