package login

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type (
	Result interface{}
)

func Solve(c *gin.Context) (re Result, err error) {
	method := c.PostForm("m")

	switch method {
	case "login":
		re, err = login(c)
	case "register":
		re, err = register(c)
	case "logout":
		re, err = logout(c)
	default:
		err = errors.New("Invalid method...")
	}

	return
}
