package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/server"
	"github.com/stretchr/testify/assert"
)

func newAPITest() (*chi.Mux, *httptest.ResponseRecorder) {
	server, _ := server.New()
	_ = server.Start()
	r := New(server.DB())

	w := httptest.NewRecorder()

	return r, w

}

func TestListApplications(t *testing.T) {
	r, w := newAPITest()
	req, _ := http.NewRequest("GET", "/api/applications", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}

func TestCreateApplications(t *testing.T) {
	r, w := newAPITest()

	reqBody, _ := json.Marshal(map[string]string{
		"Name":     "name",
		"Scenario": "this is the scenario",
	})
	req, _ := http.NewRequest("POST", "/api/applications", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var app ent.Application
	json.Unmarshal(w.Body.Bytes(), &app)
	assert.Equal(t, app.Name, "name")
	assert.Equal(t, app.Scenario, "this is the scenario")
	assert.Equal(t, app.Status, "init")
}
