package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	//"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type (
	LoginRpc struct {
		Username string `json:`
		Passwd string `json:`
	}

	LoginReply struct {
	}
)

func login(c *gin.Context) (re LoginReply, err error) {
	// TODO - validate


	if len(os.Args) != 2 {
		// TODO error
		fmt.Println("usage: ", os.Args[0], "server:port")
		log.Fatal(1)
	}
	service := os.Args[1]

	client, err := jsonrpc.Dial("tcp", service)
	defer client.Close()
	if err != nil {
		return err
	}

	args := LoginRpc{}
	err = client.Call("Login.Login", args, &re)
	if err != nil {
		return err
	}

	return
}
