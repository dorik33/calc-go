package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dorik33/calc-go/internal/handler"
	"github.com/dorik33/calc-go/internal/model"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name           string
		reqBody        string
		expectedCode   int
		expectedError  string
		expectedResult string 
	}{
		{
			name:           "Success: 1+1",
			reqBody:        `{"expression": "1+1"}`,
			expectedCode:   http.StatusOK,
			expectedResult: "2.0000", 
		},
		{
			name:           "Success: (2+2)*2",
			reqBody:        `{"expression": "(2+2)*2"}`,
			expectedCode:   http.StatusOK,
			expectedResult: "8.0000", 
		},
		{
			name:          "Error: Division by zero",
			reqBody:       `{"expression": "10/0"}`,
			expectedCode:  http.StatusUnprocessableEntity,
			expectedError: "Division by zero",
		},
		{
			name:          "Error: Invalid input",
			reqBody:       `{"expression": ""}`,
			expectedCode:  http.StatusUnprocessableEntity,
			expectedError: "Invalid input",
		},
		{
			name:          "Error: Invalid JSON",
			reqBody:       `{invalid json}`, 
			expectedCode:  http.StatusInternalServerError,
			expectedError: "Error parsing JSON", 
		},
		{
			name:          "Error: Empty Request Body",
			reqBody:       ``, 
			expectedCode:  http.StatusInternalServerError,
			expectedError: "Error parsing JSON", 
		},
		{
			name:          "Error: Unknown operator",
			reqBody:       `{"expression": "2=2"}`, 
			expectedCode:  http.StatusUnprocessableEntity,
			expectedError: "Invalid input", 
		},
		{
			name:          "Error: Parenthesis error",
			reqBody:       `{"expression": "2+2)"}`, 
			expectedCode:  http.StatusUnprocessableEntity,
			expectedError: "Invalid input", 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/calculate", bytes.NewBufferString(tt.reqBody))
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handler.CalculateHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedCode {
				t.Errorf("Handler returned wrong status code: got %v want %v", rr.Code, tt.expectedCode)
			}

			var response model.Response
			err = json.Unmarshal(rr.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("Error unmarshalling response: %v", err)
			}

			if tt.expectedError != "" {
				if !strings.HasPrefix(response.Error, tt.expectedError) {
					t.Errorf("Handler returned wrong error: got %v want %v", response.Error, tt.expectedError)
				}
			} else {
				if response.Result != tt.expectedResult {
					t.Errorf("Handler returned wrong result: got %v want %v", response.Result, tt.expectedResult)
				}
			}
		})
	}
}
