package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"syscall"
	"time"

	r "github.com/dancannon/gorethink"
	"github.com/dobegor/steamdonations/app/models"

	"github.com/dobegor/steamdonations/app/util"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"github.com/yohcop/openid-go"
)

var nonceStore = &openid.SimpleNonceStore{
	Store: make(map[string][]*openid.Nonce)}
var discoveryCache = &openid.SimpleDiscoveryCache{}

type App struct {
	*revel.Controller
	User models.User
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) SteamLogin() revel.Result {
	url, err := openid.RedirectURL(
		"http://steamcommunity.com/openid",
		revel.Config.StringDefault("app.url", "https://exsdev.ru")+"/steamidcallback",
		revel.Config.StringDefault("app.url", "https://exsdev.ru"))

	if err == nil {
		return c.Redirect(url)
	}

	return c.NotFound(err.Error())
}

func (c App) SteamLoginCallback() revel.Result {
	id, err := openid.Verify(revel.Config.StringDefault("app.url", "https://exsdev.ru")+c.Request.URL.String(), discoveryCache, nonceStore)
	if err != nil {
		return c.NotFound(err.Error())
	}

	id = id[strings.LastIndex(id, "/")+1 : len(id)]
	user, err := models.GetUser(id)
	if err != nil {
		answer := c.SteamIDLookup(id)
		user = answer.GetUser()
		if err := models.SaveUser(user); err != nil {
			return c.NotFound(err.Error())
		}
	}

	c.Session["steamid"] = user.SteamID
	c.Session.SetDefaultExpiration()
	go cache.Set(id, user, 24*time.Hour)
	return c.Redirect(App.Index)

}

func (c App) Logout() revel.Result {
	c.Session["steamid"] = ""
	return c.Redirect(App.Index)
}

func (c App) Page(page string) revel.Result {
	if util.Exists(revel.TemplatePaths[0] + "/Pages/" + page + ".html") {
		if page == "about" {
			static1, _ := models.GetServerStaticInfo(1)
			graph, _ := models.GetServerGraph(1)
			res, _ := json.Marshal(graph)
			static1.Graph = string(res)

			static2, _ := models.GetServerStaticInfo(2)
			graph, _ = models.GetServerGraph(2)
			res, _ = json.Marshal(graph)
			static2.Graph = string(res)
			/*
				static3, _ := models.GetServerStaticInfo(3)
				graph, _ = models.GetServerGraph(3)
				res, _ = json.Marshal(graph)
				static3.Graph = string(res)
			*/
			c.RenderArgs["ServerStatic_1"] = static1
			c.RenderArgs["ServerStatic_2"] = static2
		}
		return c.RenderTemplate("Pages/" + page + ".html")
	}

	return c.NotFound("Sorry, we can't find this page!")
}

func (c App) Shutdown() revel.Result {
	println("Got shutdown command. Bye!")
	pid := syscall.Getpid()
	syscall.Kill(pid, syscall.SIGTERM)
	return nil
}

func (c App) SteamIDLookup(SteamID string) models.SteamIDLookupAnswer {
	url := "http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=F3815214B3A0074C2131939B0C438AF8&steamids="
	res, _ := http.Get(url + SteamID)
	defer res.Body.Close()
	var answer models.SteamIDLookupAnswer
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&answer)
	return answer
}

func (c App) GetUser(SteamID string) revel.Result {
	user, err := models.GetUser(SteamID)
	if err != nil {
		return c.RenderText(err.Error())
	}
	return c.RenderText(user.Name)
}

func (c App) CheckUser() revel.Result {
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

func (c App) RedeemCode(SteamID string, code string) revel.Result {
	code = strings.ToLower(code)
	Code, err := models.CheckCode(SteamID, code)
	if err == nil {
		models.DisableCode(SteamID, code)
	}
	return c.RenderXml(Code)
}

func (c App) TestCodeWrite() revel.Result {
	code := models.Code{
		SteamID: "76561198033710859",
		Items: []models.Item{
			{123, 3},
			{94, 3},
		},
		Status: true,
	}
	r.DB("db").Table("codes").Insert(code).RunWrite(util.DB)
	return c.RenderText("yolo")
}
