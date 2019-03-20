package handles

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grt1st/shortcut/core"
	"github.com/grt1st/shortcut/models"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func NewShortcut(c *gin.Context) {
	uri := c.PostForm("url")
	captcha := c.PostForm("captcha")
	if captcha == "" || uri == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"msg":    html.EscapeString("please input content"),
		})
		return
	}
	err := checkGCaptcha(captcha)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"msg":    html.EscapeString(err.Error()),
		})
		return
	}
	if !strings.HasPrefix(uri, "http://") && !strings.HasPrefix(uri, "https://") {
		uri = fmt.Sprintf("http://%s", uri)
	}
	s, e := models.NewShortcut(uri)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"msg":    html.EscapeString(e.Error()),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"msg":    s.Short,
		})
	}
}

func GetShortcut(c *gin.Context) {
	var uri string
	name := c.Param("shortcut")
	if name == "" {
		uri = "/"
	} else {
		uri = core.GetFromRedis(fmt.Sprintf("shortcut:%s", name))
		if uri == "" {
			s := models.GetShortcutByShort(name)
			if s.Value == "" {
				c.Redirect(http.StatusNotFound, "/404")
				return
			} else {
				uri = s.Value
				core.SetToRedis(fmt.Sprintf("shortcut:%s", name), uri)
			}
		}
	}
	c.Redirect(http.StatusMovedPermanently, uri)
}

func checkGCaptcha(captcha string) error {

	data := make(url.Values)
	data["secret"] = []string{core.Config.Gsecret}
	data["response"] = []string{captcha}

	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", data)
	if err != nil {
		return errors.New(fmt.Sprintf("request server error: %s", err.Error()))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("parse html from server error : %s", err.Error()))
	}

	var rdata GObject
	err = json.Unmarshal(body, &rdata)
	if err != nil {
		return errors.New(fmt.Sprintf("parse json from server error : %s", err.Error()))
	}

	if rdata.Success == true {
		return nil
	} else {
		return errors.New(fmt.Sprintf("verify failed : %v", rdata.ErrorC))
	}
}

type GObject struct {
	Success  bool     `json:"success"`
	ErrorC   []string `json:"error-codes"`
	HostName string   `json:"hostname"`
}
