package models

import (
	"crypto/md5"
	"errors"
	"fmt"

	r "github.com/dancannon/gorethink"
	. "github.com/dobegor/steamdonations/app/util"
	"github.com/revel/revel"
)

type Payment struct {
	ID      string `gorethink:"id,omitempty"`
	SteamID string
	Amount  float64
	Status  string
}

func (p Payment) GenerateUrl() (string, error) {
	merchantId, found := revel.Config.String("fk.merchantId")
	if !found {
		return "", errors.New("fk.merchantId not found")
	}

	secret, found := revel.Config.String("fk.secret1")
	if !found {
		return "", errors.New("fk.secret1 not found")
	}

	fk_link := "http://www.free-kassa.ru/merchant/cash.php?m="
	amount := fmt.Sprintf("%.2f", p.Amount)
	sum := md5.Sum([]byte(merchantId + ":" + amount + ":" + secret + ":" + p.ID))
	sign := fmt.Sprintf("%x", sum)

	return fk_link + merchantId + "&oa=" + amount + "&o=" + p.ID + "&s=" + sign, nil
}

func CreatePayment(SteamID string, Amount float64) (Payment, error) {
	_, err := GetUser(SteamID)
	if err != nil {
		return Payment{}, err
	}

	if Amount < 0 {
		return Payment{}, errors.New("Amount should be positive")
	}

	payment := Payment{
		SteamID: SteamID,
		Amount:  Amount,
		Status:  "pending",
	}

	cursor, err := r.DB("db").Table("payments").
		Insert(payment).
		RunWrite(DB)
	if err != nil {
		return Payment{}, err
	}

	payment.ID = cursor.GeneratedKeys[0]
	return payment, nil
}

func GetPayment(ID string) (Payment, error) {
	cursor, err := r.DB("db").Table("payments").
		Filter(r.Row.Field("id").Eq(ID)).
		Run(DB)
	defer cursor.Close()
	if err != nil {
		return Payment{}, err
	}

	payment := Payment{}
	if err := cursor.One(&payment); err != nil {
		return Payment{}, err
	}

	return payment, nil
}

func ConfirmPayment(ID string) error {
	cursor, err := r.DB("db").Table("payments").
		Filter(r.Row.Field("id").Eq(ID)).
		Update(map[string]string{"Status": "success"}).
		RunWrite(DB)

	if cursor.Replaced == 0 && err == nil {
		return errors.New("No such pending payment with this ID.")
	}

	if err != nil {
		return err
	}

	go func() {
		payment, _ := GetPayment(ID)
		user, _ := GetUser(payment.SteamID)
		err := user.AddToBalance(payment.Amount)
		if err != nil {
			println(err.Error())
		}
	}()

	return nil
}
