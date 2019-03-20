package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/grt1st/shortcut/core"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"host": core.Config.Host,
		"protocol": core.Config.Protocol,
		"gkey": core.Config.Gkey,
	})
}

func E404(c *gin.Context) {
	c.HTML(http.StatusOK, "404", gin.H{})
}
