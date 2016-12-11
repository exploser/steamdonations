// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tUpload struct {}
var Upload tUpload


func (_ tUpload) Upload(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Upload.Upload", args).Url
}

func (_ tUpload) HandleUpload(
		bundle []byte,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "bundle", bundle)
	return revel.MainRouter.Reverse("Upload.HandleUpload", args).Url
}


type tShop struct {}
var Shop tShop


func (_ tShop) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Shop.Index", args).Url
}

func (_ tShop) CheckUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Shop.CheckUser", args).Url
}


type tPayments struct {}
var Payments tPayments


func (_ tPayments) Create(
		SteamID string,
		Amount float64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "SteamID", SteamID)
	revel.Unbind(args, "Amount", Amount)
	return revel.MainRouter.Reverse("Payments.Create", args).Url
}

func (_ tPayments) Confirm(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Payments.Confirm", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) SteamLogin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.SteamLogin", args).Url
}

func (_ tApp) SteamLoginCallback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.SteamLoginCallback", args).Url
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).Url
}

func (_ tApp) Page(
		page string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("App.Page", args).Url
}

func (_ tApp) Shutdown(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Shutdown", args).Url
}

func (_ tApp) GetUser(
		SteamID string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "SteamID", SteamID)
	return revel.MainRouter.Reverse("App.GetUser", args).Url
}

func (_ tApp) CheckUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.CheckUser", args).Url
}

func (_ tApp) RedeemCode(
		SteamID string,
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "SteamID", SteamID)
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("App.RedeemCode", args).Url
}

func (_ tApp) TestCodeWrite(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.TestCodeWrite", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


