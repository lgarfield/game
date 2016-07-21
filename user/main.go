package user

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func Solve(c *gin.Context) error {
	var err error

	method := c.PostForm("m")
	switch method {
	case "userbasicinfo":
		err = userbasicinfo(c)
	default:
		err = errors.New("Invalid method...")
	}

	return err
}

func userbasicinfo(c *gin.Context) error {
	var err error
	return err
}
