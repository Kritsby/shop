package v1

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bunrouter"
	"net/http"
	"net/http/httptest"
	"shop/internal/model"
	"shop/internal/service"
	mock_service "shop/internal/service/mocks"
	"testing"
)

func TestHandler_stock(t *testing.T) {
	type mockBehavior func(s *mock_service.MockStocker)

	tests := []struct {
		name                 string
		handler              string
		method               string
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "OK",
			handler: "reserve",
			method:  "POST",
			mockBehavior: func(s *mock_service.MockStocker) {
				s.EXPECT().Reserve([]int{1, 2}).Return(nil)
			},
			inputBody:          `{"ids":[1, 2]}`,
			expectedStatusCode: 200,
			expectedResponseBody: `{"data":"product(s) was reserved","params":null,"route":"/reserve"}
`,
		},
		{
			name:    "OK",
			handler: "release",
			method:  "POST",
			mockBehavior: func(s *mock_service.MockStocker) {
				s.EXPECT().ReserveRelease([]int{1, 2}).Return(nil)
			},
			inputBody:          `{"ids":[1, 2]}`,
			expectedStatusCode: 200,
			expectedResponseBody: `{"data":"product(s) was unreserved","params":null,"route":"/release"}
`,
		},
		{
			name:    "OK",
			method:  "GET",
			handler: "amount",
			mockBehavior: func(s *mock_service.MockStocker) {
				s.EXPECT().GetAmount(1).Return([]model.Products{model.Products{
					Storage: 1,
					Product: 1,
					Amount:  2,
				}}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"data":[{"storage_id":1,"product_id":1,"amount":2}],"params":{"storage":"1"},"route":"/amount/:storage"}
`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockStocker(c)
			testCase.mockBehavior(repo)

			services := &service.Service{
				Stocker: repo,
			}
			handlers := New(services)

			r := bunrouter.New()
			var w *httptest.ResponseRecorder
			var req *http.Request
			switch testCase.handler {
			case "amount":
				r.GET("/amount/:storage", handlers.amount)

				w = httptest.NewRecorder()
				req = httptest.NewRequest("GET", "/amount/1",
					nil)
			case "reserve":
				r.POST("/reserve", handlers.reserve)

				w = httptest.NewRecorder()
				req = httptest.NewRequest("POST", "/reserve",
					bytes.NewBufferString(testCase.inputBody))
			case "release":
				r.POST("/release", handlers.reserveRelease)

				w = httptest.NewRecorder()
				req = httptest.NewRequest("POST", "/release",
					bytes.NewBufferString(testCase.inputBody))
			}

			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
