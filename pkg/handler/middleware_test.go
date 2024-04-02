package handler

import (
	"encoding/json"
	"errors"
	"github/film-lib/pkg/service"
	mock_service "github/film-lib/pkg/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(2, "default", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"userId":"2","userRole":"default"}`,
		},
		{
			name:                 "Empty header",
			headerName:           "",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"Message":"empty auth header"}`,
		},
		{
			name:                 "Invalid Bearer",
			headerName:           "Authorization",
			headerValue:          "Bearr token",
			token:                "token",
			mockBehavior:         func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"Message":"invalid auth header"}`,
		},
		{
			name:                 "Invalid Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			mockBehavior:         func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"Message":"token is empty"}`,
		},
		{
			name:        "Service Failure",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, "default", errors.New("failed to parse token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"Message":"failed to parse token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// init deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			auth := mock_service.NewMockAuthorization(ctrl)
			testCase.mockBehavior(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			// Test Server
			mux := http.NewServeMux()
			mux.HandleFunc("/auth/protected/", func(w http.ResponseWriter, r *http.Request) {
				err := handler.userIdentity(w, r)
				if err == nil {
					json, _ := json.Marshal(map[string]interface{}{
						"userId":   w.Header().Get("userId"),
						"userRole": w.Header().Get("userRole"),
					})
					w.Write([]byte(json))
				}
				// w.Write([]byte(strings.Trim(fmt.Sprintf("%v %v", w.Header().Get("userId"), w.Header().Get("userRole")), "\n")))
			})

			// Test Request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/auth/protected/", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			// Perform Request
			mux.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.JSONEq(t, testCase.expectedResponseBody, w.Body.String())
			// assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
