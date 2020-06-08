package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gobench-io/gobench/server"
	"github.com/stretchr/testify/assert"
)

func TestListApplications(t *testing.T) {
	server, _ := server.New()
	_ = server.Start()
	r := New(server.DB())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
