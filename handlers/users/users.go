package users

import "github.com/gin-gonic/gin"

type User struct {
	Username string
	Password string
}

var users = []User{
	{Username: "Michail", Password: "1234"},
}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": users,
	})
}
