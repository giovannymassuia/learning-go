package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(500)
	assert.Nil(t, err)
	assert.Equal(t, 5.0, tax)

	tax, err = CalculateTax(0)
	assert.NotNil(t, err)
	assert.Equal(t, 0.0, tax)
	assert.Error(t, err, "amount must be greater than 0")
}

// --- mocks

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &MockRepository{}
	repository.On("Save", 10.0).Return(nil)
	repository.On("Save", 0.0).Return(errors.New("amount must be greater than 0"))
	repository.On("Save", mock.Anything).Return(nil)

	err := CalculateTaxAndSave(1000, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0, repository)
	assert.NotNil(t, err)
	assert.Error(t, err, "amount must be greater than 0")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "Save", 2)
}
