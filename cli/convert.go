package cli

import (
	"github.com/boggydigital/nod"
	"net/url"
	"strconv"
)

const (
	// ImperialPoundToSIGrams is a ratio between imperial pound and SI gram
	ImperialPoundToSIGrams = 453.59237
	// ImperialPoundToAvoirdupoisOunce is a ratio between imperial pound and Avoirdupois ounce
	ImperialPoundToAvoirdupoisOunce = 16.0
)

func ConvertHandler(u *url.URL) error {

	gramsStr := u.Query().Get("grams")
	if grams, err := strconv.ParseFloat(gramsStr, 32); err == nil {
		return Convert(grams)
	} else {
		return err
	}
}

// Convert grams to pounds and ounces
func Convert(grams float64) error {

	ca := nod.Begin("%.2f gr =", grams)
	defer ca.Done()

	fractionalPounds := grams / ImperialPoundToSIGrams
	pounds := int(fractionalPounds)
	ounces := (fractionalPounds - float64(pounds)) * ImperialPoundToAvoirdupoisOunce

	// round a bit more if there is almost 16 ounces
	if 16-ounces <= 0.01 {
		pounds++
		ounces -= 16
	}

	ca.EndWithResult("%d lb %.2f oz", pounds, ounces)

	return nil
}
