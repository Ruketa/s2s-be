package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouterHealthCheck(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/healthcheck", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"msg\":\"ok\"}", w.Body.String())
}
