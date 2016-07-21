package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

type (
	RegisterRpc struct {
		username string `json:"name" binding:"required,alphanum"`
		passwd string `json:"passwd" binding:"required,alphanum"`
		regIP string ``
		regTime string ``
	}

	RegisterReply struct {

	}
)

func register(c *gin.Context) (re RegisterReply, err error) {
	// TODO - validate
	var form RegisterRpc
	if err := c.Bind(&form); err != nil {
		return err
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		log.Fatal(1)
	}
	service := os.Args[1]

	client, err := jsonrpc.Dial("tcp", service)
	defer client.Close()
	if err != nil {
		return err
	}

	args := RegisterRpc{}
	err = client.Call("Login.Register", args, &re)
	if err != nil {
		return err
	}

	return
}
