package tax

import "errors"

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than 0")
	}

	if amount >= 1000 {
		return 10.0, nil
	}

	return 5.0, nil
}

// --- mocks

type Repository interface {
	Save(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTaxV2(amount)
	return repository.Save(tax)
}

func CalculateTaxV2(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}

	if amount >= 1000 {
		return 10.0
	}

	return 5.0
}
