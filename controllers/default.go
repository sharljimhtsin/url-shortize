package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"regexp"
	"strings"
	"website/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Welcome() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "welcome.tpl"
}

func (c *MainController) Index() {
	c.TplName = "index.tpl"
}

func (c *MainController) Add() {
	url := c.GetString("url")
	if len(strings.TrimSpace(url)) == 0 {
		c.CustomAbort(http.StatusBadRequest, "URL NULL")
	} else {
		isUrl, err := regexp.MatchString("[a-zA-z]+://[^\\s]*", url)
		if err != nil || !isUrl {
			c.CustomAbort(http.StatusBadRequest, "BAD URL")
		} else {
			isOK, res := models.AddUrl(url)
			if isOK {
				c.Data["Hash"] = res.B
				c.Data["Url"] = res.A
				c.Data["Host"] = c.Ctx.Request.Host
				c.TplName = "added.tpl"
			} else {
				c.CustomAbort(http.StatusBadRequest, "ERROR")
			}
		}
	}
}

func (c *MainController) Query() {
	hash := c.GetString(":hash")
	data := models.GetUrlByHash(hash)
	if len(strings.TrimSpace(data.A)) == 0 {
		c.CustomAbort(http.StatusNotFound, "NO FOUND")
	} else {
		c.Redirect(data.A, http.StatusMovedPermanently)
	}
}

func (c *MainController) List() {
	data := models.GetAllUrls()
	c.Data["json"] = data
	_ = c.ServeJSON()
}
