package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminGet(c *gin.Context) {
	user, _ := c.Get("User")
	uri, _ := c.Get("Uri")
	c.HTML(http.StatusOK, "admin/home/show", gin.H{
		"Title": "Admin dashboard",
		"User":  user,
		"Uri":   uri,
	})
}
