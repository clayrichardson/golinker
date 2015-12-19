package controllers

import "github.com/revel/revel"

type Link struct {
	*revel.Controller
}

func (c Link) RedirectTo(link string) revel.Result {
	return c.Redirect("http://google.com/")
}
