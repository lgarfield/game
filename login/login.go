package login

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func Solve(c *gin.Context) error {
	var err error
	method := c.PostForm("m")
	switch method {
	case "login":
		err = login(c)
	case "register":
		err = register(c)
	case "logout":
		err = logout(c)
	default:
		err = errors.New("Invalid method...")
	}

	return err
}

func login(c *gin.Context) error {
	return errors.New("Invalid method...")
}

func register(c *gin.Context) error {
	return errors.New("Invalid method...")
}

func logout(c *gin.Context) error {
	return errors.New("Invalid method...")
}
