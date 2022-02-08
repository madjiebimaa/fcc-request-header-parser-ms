package handlers_test

import (
	"encoding/json"
	"fcc-request-header-parser-ms/handlers"
	"fcc-request-header-parser-ms/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestWhoAmI(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ipAddress := "192.168.43.192"
	acceptLang := "en-US,en;q=0.9"
	userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36"

	sucResBody, err := json.Marshal(models.WhoAmI{
		IPAddress: ipAddress,
		Language:  acceptLang,
		Software:  userAgent,
	})
	assert.NoError(t, err)

	// ipNotFoundResBody, err := json.Marshal(gin.H{
	// 	"message": "can't get ip address",
	// })
	// assert.NoError(t, err)

	langNotFoundResBody, err := json.Marshal(gin.H{
		"message": "can't get language from the header",
	})
	assert.NoError(t, err)

	userNotFoundResBody, err := json.Marshal(gin.H{
		"message": "can't get software that used by user from the header",
	})
	assert.NoError(t, err)

	testCases := []struct {
		name       string
		reqHead    func(req *http.Request)
		statusCode int
		resBody    []byte
	}{
		{
			name: "success",
			reqHead: func(req *http.Request) {
				req.Header.Set("Accept-Language", acceptLang)
				req.Header.Set("User-Agent", userAgent)
			},
			statusCode: http.StatusOK,
			resBody:    sucResBody,
		},
		// {
		// 	name: "fail ip address not found",
		// 	reqHead: func(req *http.Request) {
		// 		req.Header.Set("Accept-Language", acceptLang)
		// 		req.Header.Set("User-Agent", userAgent)
		// 	},
		// 	statusCode: http.StatusBadRequest,
		// 	resBody:    ipNotFoundResBody,
		// },
		{
			name: "fail language header not found",
			reqHead: func(req *http.Request) {
				req.Header.Set("User-Agent", userAgent)
			},
			statusCode: http.StatusBadRequest,
			resBody:    langNotFoundResBody,
		},
		{
			name: "fail user agent header not found",
			reqHead: func(req *http.Request) {
				req.Header.Set("Accept-Language", acceptLang)
			},
			statusCode: http.StatusBadRequest,
			resBody:    userNotFoundResBody,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/api/whoami", nil)
			assert.NoError(t, err)

			tt.reqHead(req)

			rec := httptest.NewRecorder()
			_, r := gin.CreateTestContext(rec)
			r.GET("/api/whoami", handlers.WhoAmIHandler)

			r.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
			assert.Equal(t, tt.resBody, rec.Body.Bytes())
		})
	}
}
