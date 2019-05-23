package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

type AboutController struct {
	beego.Controller
}

type CaseController struct {
	beego.Controller
}

type NewsController struct {
	beego.Controller
}

type NewsDeyailController struct {
	beego.Controller
}

type ProductController struct {
	beego.Controller
}

type EditorController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.TplName = "index.html"
}

func (c *AboutController) Get() {
	c.TplName = "about.html"
}

func (c *CaseController) Get() {
	c.TplName = "case.html"
}

func (c *NewsController) Get() {
	c.TplName = "news.html"
}

func (c *NewsDeyailController) Get() {
	c.TplName = "newsDetail.html"
}

func (c *ProductController) Get() {
	c.TplName = "product.html"
}

func (e *EditorController) Get() {
	e.TplName = "editor.html"
}
