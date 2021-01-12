package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testingUserError string

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func (e testingUserError) Error() string {
	return string(e)
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func TestErrWrapper(t *testing.T) {
	tests := []struct {
		h       appHandler
		code    int
		message string
	}{
		{errPanic, 500, "Internal Server Error"},
		{errUserError, 400, "user error"},
	}

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "https://www.google.com/", nil)
		f(response, request)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")

		if response.Code != tt.code ||
			body != tt.message {
			t.Errorf("expect (%d,%s); "+"got (%d,%s)", tt.code, tt.message, response.Code, body)
		}
	}

}
