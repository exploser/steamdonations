package controllers

import (
	"crypto/md5"
	"fmt"

	r "github.com/dancannon/gorethink"
	"github.com/dobegor/steamdonations/app/models"
	. "github.com/dobegor/steamdonations/app/util"
	"github.com/revel/revel"
)

type Payments struct {
	*revel.Controller
	User models.User
}

func (c Payments) Create(SteamID string, Amount float64) revel.Result {
	if Amount <= 0 {
		return c.RenderText("Error: Amount should be positive")
	}

	payment, err := models.CreatePayment(SteamID, Amount)
	if err != nil {
		return c.RenderText("Error: " + err.Error())
	}

	url, err := payment.GenerateUrl()
	if err != nil {
		return c.RenderText("Error: " + err.Error())
	}

	return c.RenderText(url)
}

func (c Payments) Confirm() revel.Result {
	go r.DB("db").Table("payments_raw").
		Insert(c.Params.Values).
		RunWrite(DB)

	amountStr := c.Params.Get("AMOUNT")

	amount, err := ParseAmount(amountStr)

	if err != nil {
		return c.RenderText("Error: " + err.Error())
	}

	signGot := c.Params.Get("SIGN")
	merchantID := c.Params.Get("MERCHANT_ID")
	orderID := c.Params.Get("MERCHANT_ORDER_ID")
	secret, found := revel.Config.String("fk.secret2")

	if !found {
		return c.RenderText("Error: FK secret code #2 not found")
	}

	signCalculated := fmt.Sprintf("%x", md5.Sum([]byte(merchantID+":"+amountStr+":"+secret+":"+orderID)))
	if signCalculated == signGot {
		go models.ConfirmPayment(orderID)
		return c.RenderText("Payment confirmed, amount: " + fmt.Sprintf("%.2f", amount))
	} else {
		return c.RenderText("Signature mismatch, got: " + signGot + ", calculated: " + signCalculated)
	}
}
