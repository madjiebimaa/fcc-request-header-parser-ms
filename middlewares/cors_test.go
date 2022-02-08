package middlewares_test

import (
	"fcc-request-header-parser-ms/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c, r := gin.CreateTestContext(rec)

	h := middlewares.CORS()

	r.ServeHTTP(rec, req)

	h(c)

	assert.Equal(t, "*", rec.Header().Get("Access-Control-Allow-Origin"))
}
