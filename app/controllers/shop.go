package controllers

import (
	"github.com/dobegor/steamdonations/app/models"
	"github.com/revel/revel"
)

type Shop struct {
	*revel.Controller
	User models.User
}

func (c Shop) Index() revel.Result {
	return c.Render()
}

func (c Shop) CheckUser() revel.Result {
	SteamID := c.Session["steamid"]
	if len(SteamID) > 0 {
		user, err := models.GetUser(SteamID)
		if err != nil {
			c.User = models.User{}
			c.RenderArgs["User"] = c.User
			return nil
		}
		c.User = user
		c.RenderArgs["User"] = user
		return nil
	}
	c.User = models.User{}
	c.RenderArgs["User"] = c.User
	return nil
}
