package handlers

import (
	"github.com/stretchr/testify/mock"
)

type mysqlMock struct {
	mock.Mock
}

func (o mysqlMock) Ping() (bool, error) {
	args := o.Called()
	return args.Bool(0), args.Error(1)
}

func (o mysqlMock) GetToken(token string) (string, error) {
	args := o.Called()
	return args.String(0), args.Error(1)
}

func (o mysqlMock) Connect() {
}
