package internal

const (
	// ImperialPoundToSIGrams is a ratio between imperial pound and SI gram
	ImperialPoundToSIGrams = 453.59237
	// ImperialPoundToAvoirdupoisOunce is a ratio between imperial pound and Avoirdupois ounce
	ImperialPoundToAvoirdupoisOunce = 16.0
)

// Convert grams to pounds and ounces
func Convert(grams float64) (int, float64) {
	fractionalPounds := grams / ImperialPoundToSIGrams
	pounds := int(fractionalPounds)
	ounces := (fractionalPounds - float64(pounds)) * ImperialPoundToAvoirdupoisOunce

	// round a bit more if there is almost 16 ounces
	if 16-ounces <= 0.01 {
		pounds++
		ounces -= 16
	}

	return pounds, ounces
}
