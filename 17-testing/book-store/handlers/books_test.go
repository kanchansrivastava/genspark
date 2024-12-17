package handlers

import (
	"book-store/middlewares"
	"book-store/models"
	"book-store/models/mockmodels"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBook(t *testing.T) {
	mockBook := models.Book{
		ID:          1,
		Title:       "Go Programming",
		AuthorEmail: "John@doe.com",
		AuthorName:  "John Doe",
		Description: "Learn Go programming with this book",
		Price:       29.99,
		Stock:       1,
		Category:    "Fiction",
	}

	tt := [...]struct {
		name             string
		body             []byte // Body to send to the request
		expectedStatus   int
		expectedResponse string
		MockStore        func(m *mockmodels.MockService)
	}{
		{
			name: "OK",
			body: []byte(`{
   					 "title": "Go Programming",
   					 "author_name": "John Doe",
   					 "author_email": "John@doe.com",
   					 "description": "Learn Go programming with this book and enhance your skills.",
   					 "category": "Programming",
   					 "price": 29.99,
   					 "stock": 10
				}`),
			expectedStatus:   http.StatusCreated,
			expectedResponse: `{"id":1,"title":"Go Programming","author_Name":"John Doe","author_email":"John@doe.com","description":"Learn Go programming with this book","category":"Fiction","price":29.99}`,
			MockStore: func(m *mockmodels.MockService) {
				// setting the expectations for the mock call
				m.EXPECT().InsertBook(gomock.Any(), gomock.Any()).Return(mockBook, nil).Times(1)
			},
		},
		// input
		// output
	}
	router := gin.New()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockDb := mockmodels.NewMockService(ctrl)

	h := Handler{
		service:  mockDb,
		validate: validator.New(),
	}
	router.POST("/create", h.CreateBook)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// passing the mockDB to MockStore to set expectation for the mocking calls
			tc.MockStore(mockDb)

			// we need a traceId in the create book, creating a fake one and putting it in the context
			traceId := "fake-trace-id"
			ctx := context.WithValue(context.Background(), middlewares.TraceIdKey, traceId)

			//creating the request with context
			req := httptest.NewRequestWithContext(ctx, http.MethodPost, "/create", bytes.NewReader(tc.body))

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			require.Equal(t, tc.expectedStatus, rec.Code)
			require.Equal(t, tc.expectedResponse, rec.Body.String())

		})
	}
}
