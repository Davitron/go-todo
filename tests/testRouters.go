package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func postRouter(url string, body *bytes.Buffer, handeler gin.HandlerFunc) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	r.POST(url, handeler)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
}
