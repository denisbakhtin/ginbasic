package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeGet(c *gin.Context) {
	user, _ := c.Get("User")
	c.HTML(http.StatusOK, "home/show", gin.H{
		"Title": "Welcome to basic GIN web-site",
		"User":  user,
	})
}
