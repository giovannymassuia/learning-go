package tax

import (
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(tax float64) error {
	args := m.Called(tax) // m.Called(tax) returns a struct with the arguments passed to the method
	return args.Error(0)  // args.Error(0) returns the first argument as an error
}
