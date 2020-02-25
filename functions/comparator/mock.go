package comparator

import "github.com/stretchr/testify/mock"

type MockComparator struct {
	mock.Mock
}

func (mock *MockComparator) Compare(one interface{}, two interface{}) (int, error) {
	args := mock.Called(one, two)
	return args.Int(0), args.Error(1)
}
