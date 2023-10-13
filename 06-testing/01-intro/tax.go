package tax

import "time"

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}

	if amount >= 1000 {
		return 10.0
	}

	return 5.0
}

func CalculateTaxWithSleep(amount float64) float64 {

	// sleep for 1 ms
	time.Sleep(1 * time.Millisecond)

	if amount >= 1000 {
		return 10.0
	}

	return 5.0
}
