package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// ImperialPoundToSIGrams is a ratio between imperial pound and SI gram
	ImperialPoundToSIGrams = 453.59237
	// ImperialPoundToAvoirdupoisOunce is a ratio between imperial pound and Avoirdupois ounce
	ImperialPoundToAvoirdupoisOunce = 16.0
)

func main() {
	usage := "glo: grams (g) to pounds (lb) and ounces (oz) converter\n" +
		"usage: glo [grams]"

	if len(os.Args) < 2 {
		fmt.Println(usage)
	} else {
		grams, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Printf("\"%s\" is not a valid value in grams\n", os.Args[1])
			os.Exit(1)
		}
		fractionalPounds := grams / ImperialPoundToSIGrams
		pounds := int(fractionalPounds)
		ounces := (fractionalPounds - float64(pounds)) * ImperialPoundToAvoirdupoisOunce

		// round a bit more if there is almost 16 ounces
		if 16-ounces <= 0.01 {
			pounds++
			ounces -= 16
		}

		fmt.Printf("%vg is about %vlb", grams, pounds)
		if ounces > 0.01 {
			fmt.Printf(" %.2foz", ounces)
		}
		fmt.Println()
	}
}
