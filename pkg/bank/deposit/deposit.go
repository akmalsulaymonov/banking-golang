package deposit

// Calculate calculates deposit
func Calculate(amount int, currency string) (min int, max int) {

	minPercent, maxPercent := PercentFor(currency)

	min = amount * minPercent / 100 / 12
	max = amount * maxPercent / 100 / 12

	//return min, max
	return
}

// PercentFor returns min and max percent value
func PercentFor(currency string) (min int, max int) {
	switch currency {
	case "TJS":
		return 4, 6
	case "USD":
		return 3, 5
	default:
		return 0, 0
	}
}
