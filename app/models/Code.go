package models

import (
	r "github.com/dancannon/gorethink"
	. "github.com/dobegor/steamdonations/app/util"
)

type Item struct {
	ID       int
	Quantity int
}

type Code struct {
	SteamID string `xml:"SteamID,omitempty"`
	Items   []Item `xml:"Items,omitempty>Item,omitempty"`
	Status  bool
}

func CheckCode(SteamID string, code string) (Code, error) {
	cursor, err := r.DB("db").Table("codes").
		Filter(r.Row.Field("Status").Eq(true).
		And(r.Row.Field("id").Match("^" + code)).
		And(r.Row.Field("SteamID").Eq(SteamID))).
		Run(DB)
	if err != nil {
		return Code{}, err
	}
	CodeObj := Code{}

	if err := cursor.One(&CodeObj); err != nil {
		return Code{}, err
	}

	return CodeObj, nil
}

func DisableCode(SteamID string, code string) {
	r.DB("db").Table("codes").
		Filter(r.Row.Field("Status").Eq(true).
		And(r.Row.Field("id").Match("^" + code)).
		And(r.Row.Field("SteamID").Eq(SteamID))).
		Update(map[string]interface{}{"Status": false}).
		RunWrite(DB)
}
