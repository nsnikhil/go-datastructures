package operator

import "github.com/stretchr/testify/mock"

type MockUnaryOperator struct {
	mock.Mock
}

func (mock *MockUnaryOperator) Apply(e interface{}) interface{} {
	args := mock.Called(e)
	return args.Get(0).(interface{})
}
