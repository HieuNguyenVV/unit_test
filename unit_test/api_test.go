package unit_test

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"unit_test/api"
	"unit_test/models"
)

func TestCreateUser(t *testing.T) {
	e := echo.New()
	e.POST("/create", api.CreateUser)

	type (
		testCaseInput models.CreateUserRequest
	)

	testCases := []struct {
		name  string
		input testCaseInput
		code  int
	}{
		{
			name: "1: Bad request",
			input: testCaseInput{
				ID:       "",
				Username: "",
			},
			code: 400,
		},
		{
			name: "2: Server error",
			input: testCaseInput{
				ID:       "123",
				Username: "Tes",
			},
			code: 500,
		},
		{
			name: "3: Success",
			input: testCaseInput{
				ID:       "123456",
				Username: "Nothing's gonna change my love for you",
			},
			code: 200,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			jsonInput, err := json.Marshal(&v.input)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}
			req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(jsonInput))
			if err != nil {
				t.Fatalf("Error: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			e.ServeHTTP(res, req)

			if ok := assert.Equal(t, v.code, res.Code); !ok {
				t.Fatalf("Test fail, exp: %v, got: %v", v.code, res.Code)
			}
			t.Logf("Test case: %v passed", v.name)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	e := echo.New()
	e.POST("/create", api.UpdateUser)

	type (
		testCaseInput models.CreateUserRequest
	)

	testCases := []struct {
		name  string
		input testCaseInput
		code  int
	}{
		{
			name: "1: Bad request",
			input: testCaseInput{
				ID:       "",
				Username: "",
			},
			code: 400,
		},
		{
			name: "2: Server error",
			input: testCaseInput{
				ID:       "123",
				Username: "Tes",
			},
			code: 500,
		},
		{
			name: "3: Success",
			input: testCaseInput{
				ID:       "123456",
				Username: "Nothing's gonna change my love for you",
			},
			code: 200,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			jsonInput, err := json.Marshal(&v.input)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}
			req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(jsonInput))
			if err != nil {
				t.Fatalf("Error: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			e.ServeHTTP(res, req)

			if ok := assert.Equal(t, v.code, res.Code); !ok {
				t.Fatalf("Test fail, exp: %v, got: %v", v.code, res.Code)
			}
			t.Logf("Test case: %v passed", v.name)
		})
	}
}
