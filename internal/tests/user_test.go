package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run(
		"user get error if bad formatted uuid is not provided", func(t *testing.T) {
			testPath := "/api/user/123"
			testMethod := "GET"
			response := httptest.NewRecorder()
			_ = httptest.NewRequest(testMethod, testPath, nil)
			assert.Equal(t, http.StatusBadRequest, response.Code, "status code match")
		},
	)

	t.Run(
		"user get error if uuid not in database is provided", func(t *testing.T) {
			testPath := "/api/user/32404d15-a878-4392-8ce3-75d4b7e038ce"
			testMethod := "GET"
			response := httptest.NewRecorder()
			_ = httptest.NewRequest(testMethod, testPath, nil)
			assert.Equal(t, http.StatusInternalServerError, response.Code, "status code match")
		},
	)
}
