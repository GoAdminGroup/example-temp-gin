package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// define  no router do what
func noRouterBiz(c *gin.Context) {
	//c.String(http.StatusNotFound, "The incorrect API route.")
	contentType := c.GetHeader("Accept")
	if strings.Contains(contentType, "text/html") {
		//route404 := errdef.NoRoute404("err", c.Request.URL.String())
		//c.Redirect(http.StatusMovedPermanently, route404)
		c.String(http.StatusMovedPermanently, "not found url: "+c.Request.URL.String())
	} else {
		c.String(http.StatusInternalServerError, "The incorrect API route.")
	}
}
