package models

import (
	"errors"
	"time"

	r "github.com/dancannon/gorethink"
	. "github.com/dobegor/steamdonations/app/util"
	"github.com/revel/revel/cache"
)

type User struct {
	SteamID      string
	Name         string
	ProfileURL   string
	Avatar       string
	AvatarMedium string
	AvatarFull   string
	ID           string `gorethink:"id,omitempty"`
}

func (u User) Balance() (float64, error) {
	cursor, err := r.DB("db").Table("users").
		Get(u.ID).
		Field("Money").
		Run(DB)

	if err != nil {
		return -1, err
	}
	defer cursor.Close()

	var balance float64
	if err := cursor.One(&balance); err != nil {
		return -1, err
	}

	return balance, nil
}

func (u User) SetBalance(balance float64) error {
	if balance < 0 {
		return errors.New("Balance should be positive")
	}

	_, err := r.DB("db").Table("users").
		Get(u.ID).
		Update(map[string]float64{"Money": balance}).
		RunWrite(DB)

	if err != nil {
		return err
	}

	return nil
}

func (u User) AddToBalance(sum float64) error {
	if sum < 0 {
		return errors.New("Amount should be positive")
	}

	_, err := r.DB("db").Table("users").
		Get(u.ID).
		Update(map[string]interface{}{"Money": r.Row.Field("Money").Add(sum)}).
		RunWrite(DB)

	if err != nil {
		return err
	}

	return nil
}

func (u User) SubstractFromBalance(sum float64) error {
	if sum < 0 {
		return errors.New("Amount should be positive")
	}

	if balance, _ := u.Balance(); balance < sum {
		return errors.New("User doesn't have enough money")
	}

	_, err := r.DB("db").Table("users").
		Filter(r.Row.Field("id").
		Eq(u.ID)).
		Update(map[string]interface{}{"Money": r.Row.Field("Money").Add(sum)}).
		RunWrite(DB)

	if err != nil {
		return err
	}

	return nil
}

func GetUser(SteamID string) (User, error) {
	user := User{}
	if err := cache.Get(SteamID, &user); err == nil {
		return user, nil
	}

	cursor, err := r.DB("db").Table("users").Filter(r.Row.Field("SteamID").Eq(SteamID)).Run(DB)
	if err != nil {
		return User{}, err
	}

	if err := cursor.One(&user); err != nil {
		return User{}, errors.New("No such user")
	}

	go cache.Set(SteamID, user, 24*time.Hour)
	return user, nil
}

func SaveUser(user User) error {
	SteamID := user.SteamID
	cursor, err := r.DB("db").Table("users").Filter(r.Row.Field("SteamID").Eq(SteamID)).Count().Run(DB)
	if err != nil {
		return err
	}

	var count int
	cursor.One(&count)
	if count > 0 {
		return errors.New("User with this SteamID already exists.")
	}

	if _, err := r.DB("db").Table("users").
		Insert(
		struct {
			User
			Money int
		}{
			user,
			0,
		}).RunWrite(DB); err != nil {
		return err
	}

	return nil
}
