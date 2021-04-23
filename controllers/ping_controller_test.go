package controllers

import (
	"fmt"
	"goUnitTest/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type pingServiceMoke struct {
	handlePingFn func() (string, error)
}

func (mock pingServiceMoke) HandlePing() (string, error) {
	return mock.handlePingFn()
}
func TestPingWithError(t *testing.T) {
	services.PingService = pingServiceMoke{
		handlePingFn: func() (string, error) {
			fmt.Println("im the mocking..")
			return "", fmt.Errorf("error")
		},
	}

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	Ping(context)

	if response.Code != http.StatusInternalServerError {
		t.Error("response code should be 500")
	}

	if response.Body.String() != "error" {
		t.Error("response body should say error")
	}
}
func TestPingNoError(t *testing.T) {
	services.PingService = pingServiceMoke{
		handlePingFn: func() (string, error) {
			fmt.Println("im the mocking..")
			return "pong", nil
		},
	}
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	Ping(context)

	if response.Code != http.StatusOK {
		t.Error("response code should be 200")
	}

	if response.Body.String() != "pong" {
		t.Error("response body should say pong")
	}
}
