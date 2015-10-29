package controllers

import (
	"net/http"

	"html/template"

	"github.com/denisbakhtin/ginbasic/models"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

// GET /pages/:id route
func PageGet(c *gin.Context) {
	page, err := models.GetPage(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	if !page.Published {
		c.HTML(http.StatusNotFound, "errors/404", nil)
		return
	}
	c.HTML(http.StatusOK, "pages/show", gin.H{
		"Title":       page.Name,
		"Description": template.HTML(string(blackfriday.MarkdownCommon([]byte(page.Description)))),
	})
}
