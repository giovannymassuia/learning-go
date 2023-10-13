package tax

import (
	"github.com/stretchr/testify/assert"
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
