package testing

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func PerformRequest(handler gin.HandlerFunc, body *bytes.Buffer, method, path string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(recorder)
	r.Handle(method, path, handler)
	c.Request, _ = http.NewRequest(http.MethodPost, path, body)
	r.ServeHTTP(recorder, c.Request)
	return recorder
}
