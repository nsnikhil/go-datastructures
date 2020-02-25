package iterator

import "github.com/stretchr/testify/mock"

type MockIterator struct {
	mock.Mock
}

func (mock *MockIterator) HasNext() bool {
	args := mock.Called()
	return args.Bool(0)
}

func (mock *MockIterator) Next() interface{} {
	args := mock.Called()
	return args.Get(0).(interface{})
}
