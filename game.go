package main

import (
	"errors"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"game/achievement"
	"game/friend"
	"game/gamedata"
	"game/leaderboard"
	"game/login"
	"game/message"
	"game/order"
	"game/reward"
	"game/slot"
	"game/task"
	"game/user"
	"net/http"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func main() {
	// Create a gin router with default middleware(Logger and Recovery).
	router := gin.Default()

	router.Any("/", func(c *gin.Context) {
		re, err := clientSolve(c)
		if err != nil {
			c.JSON(http.StatusOK, re)
		}

		c.JSON(http.StatusOK, gin.H{})
	})

	endless.ListenAndServe(":8080", router)
}

func clientSolve(c *gin.Context) (re interface{}, err error) {
	path := c.PostForm("p")

	switch path {
	case "achievement":
		re, err = achievement.Solve(c)
	case "friend":
		re, err = friend.Solve(c)
	case "gamedata":
		re, err = gamedata.Solve(c)
	case "leaderboard":
		re, err = leaderboard.Solve(c)
	case "login":
		re, err = login.Solve(c)
	case "message":
		re, err = message.Solve(c)
	case "order":
		re, err = order.Solve(c)
	case "reward":
		re, err = reward.Solve(c)
	case "slot":
		re, err = slot.Solve(c)
	case "task":
		re, err = task.Solve(c)
	case "user":
		re, err = user.Solve(c)
	default:
		err = errors.New("Invalid input...")
	}

	return
}
