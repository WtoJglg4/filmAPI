package handler

import (
	"bytes"
	"errors"
	filmapi "github/film-lib"
	"github/film-lib/pkg/service"
	mock_service "github/film-lib/pkg/service/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user filmapi.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           filmapi.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username": "Test", "password": "qwerty"}`,
			inputUser: filmapi.User{
				Username: "Test",
				Password: "qwerty",
				Role:     "default",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user filmapi.User) {
				s.EXPECT().CreateUser(user).Return(2, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":2}`,
		},
		{
			name:                "Empty Fields",
			inputBody:           `{"username": "", "password": ""}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, user filmapi.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"Message":"error: passed structure does not have all the required fields"}`,
		},
		{
			name:      "Admin",
			inputBody: `{"username": "admin", "password": "admin"}`,
			inputUser: filmapi.User{
				Username: "admin",
				Password: "admin",
				Role:     "default",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user filmapi.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"username": "Test", "password": "qwerty"}`,
			inputUser: filmapi.User{
				Username: "Test",
				Password: "qwerty",
				Role:     "default",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user filmapi.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"Message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// init deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			auth := mock_service.NewMockAuthorization(ctrl)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			// Test Server
			mux := http.NewServeMux()
			mux.HandleFunc("/auth/sign-up/", handler.signUp)

			// Test Request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/auth/sign-up/", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			mux.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, strings.Trim(w.Body.String(), "\n"))
		})
	}
}
