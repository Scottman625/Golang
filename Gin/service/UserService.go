package service

import (
	"golangAPI/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userList = []pojo.User{}

func FindAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
	return
}

func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "error :"+err.Error())
		return
	}
	userList = append(userList, user)
	c.JSON(http.StatusOK, "Successfully Posted")
}
