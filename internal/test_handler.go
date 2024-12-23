package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           Request
		expectedCode   int
		expectedResult string
		expectedError  string
	}{
		{
			name:           "Valid Expression",
			method:         http.MethodPost,
			body:           Request{Expression: "1 + 1"},
			expectedCode:   http.StatusOK,
			expectedResult: `{"result":"2"}`,
		},
		{
			name:          "Invalid JSON",
			method:        http.MethodPost,
			body:          Request{},
			expectedCode:  http.StatusBadRequest,
			expectedError: "Некорректный JSON",
		},
		{
			name:          "Invalid Expression",
			method:        http.MethodPost,
			body:          Request{Expression: "1 + a"},
			expectedCode:  http.StatusUnprocessableEntity,
			expectedError: `{"error": "Expression is not valid"}`,
		},
		{
			name:          "Division by Zero",
			method:        http.MethodPost,
			body:          Request{Expression: "1 / 0"},
			expectedCode:  http.StatusInternalServerError,
			expectedError: `{"error": "Internal server error"}`,
		},
		{
			name:         "Method Not Allowed",
			method:       http.MethodGet,
			body:         Request{Expression: "1 + 1"},
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:          "Malformed Expression",
			method:        http.MethodPost,
			body:          Request{Expression: "(1 + 2"},
			expectedCode:  http.StatusUnprocessableEntity,
			expectedError: `{"error": "Expression is not valid"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем новый запрос
			var body bytes.Buffer
			if tt.method == http.MethodPost {
				json.NewEncoder(&body).Encode(tt.body)
			}

			req := httptest.NewRequest(tt.method, "/calculate", &body)
			rr := httptest.NewRecorder()

			CalculateHandler(rr, req)

			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedCode)
			}

			if tt.expectedResult != "" {
				if rr.Body.String() != tt.expectedResult {
					t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedResult)
				}
			}

			if tt.expectedError != "" {
				if rr.Body.String() != tt.expectedError {
					t.Errorf("handler returned unexpected error: got %v want %v", rr.Body.String(), tt.expectedError)
				}
			}
		})
	}
}
