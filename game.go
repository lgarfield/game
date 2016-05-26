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
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func main() {
	// Create a gin router with default middleware(Logger and Recovery).
	router := gin.Default()

	router.Any("/", func(c *gin.Context) {
		err := clientSolve(c)
		if err != nil {

		}
	})

	endless.ListenAndServe(":8080", router)
}

func clientSolve(c *gin.Context) (err error) {
	path := c.PostForm("p")

	switch path {
	case "achievement":
		err = achievement.Solve(c)
	case "friend":
		err = friend.Solve(c)
	case "gamedata":
		err = gamedata.Solve(c)
	case "leaderboard":
		err = leaderboard.Solve(c)
	case "login":
		err = login.Solve(c)
	case "message":
		err = message.Solve(c)
	case "order":
		err = order.Solve(c)
	case "reward":
		err = reward.Solve(c)
	case "slot":
		err = slot.Solve(c)
	case "task":
		err = task.Solve(c)
	case "user":
		err = user.Solve(c)
	default:
		err = errors.New("Invalid input...")
	}

	return err
}
