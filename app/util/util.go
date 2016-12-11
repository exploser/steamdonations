package util

import (
	"os"
	"strconv"

	r "github.com/dancannon/gorethink"
)

var DB *r.Session

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ParseAmount(amount string) (float64, error) {
	amountInt, err := strconv.Atoi(amount)

	if err != nil {
		amountFloat, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			return -1, err
		}
		return amountFloat, nil
	}

	return float64(amountInt), nil
}

func TruncateFloat(f float64) float64 {
	return float64(int(f*100)) / 100
}
