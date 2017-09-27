package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulHealthcheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res := httptest.NewRecorder()
	mysqlClient := new(mysqlMock)

	mysqlClient.On("Ping").Return(true, nil)

	handler := http.HandlerFunc(HealthcheckHandler(mysqlClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 200)
	assert.Equal(t, res.Body.String(), "{\"status\":\"WORKING\",\"services\":[{\"working\":true,\"service\":\"MySQL\"}]}\n")
}

func TestFailedHealthcheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res := httptest.NewRecorder()
	mysqlClient := new(mysqlMock)

	mysqlClient.On("Ping").Return(false, nil)

	handler := http.HandlerFunc(HealthcheckHandler(mysqlClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 500)
	assert.Equal(t, res.Body.String(), "{\"status\":\"FAILED\",\"services\":[{\"working\":false,\"service\":\"MySQL\"}]}\n")
}
