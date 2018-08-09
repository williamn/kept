package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/*url", func(c *gin.Context) {
		u := c.Param("url")
		u = strings.TrimPrefix(u, "/")
		u = path.Clean(u)

		if u == "." {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Main website",
			})
		} else {
			validURL, err := url.ParseRequestURI(u)
			switch err != nil {
			case true:
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			case false:
				c.JSON(http.StatusOK, gin.H{
					"url": validURL,
				})
			}
		}
	})
	r.Run()
}
