package web

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/master"
	"github.com/stretchr/testify/assert"
)

func newAPITest(t *testing.T, adminPassword string) (*chi.Mux, *httptest.ResponseRecorder) {
	logger := logger.NewNopLogger()
	m, _ := master.NewMaster(&master.Options{
		Addr:    "0.0.0.0",
		Port:    8080,
		HomeDir: "/tmp",
	}, logger)

	m.SetIsScheduled(false)

	err := m.Start()
	assert.Nil(t, err)
	h := newHandler(m, adminPassword, logger)
	r := h.r

	w := httptest.NewRecorder()

	return r, w
}

func newApp(t *testing.T) *ent.Application {
	r, w := newAPITest(t, "")
	name := "name 1"
	scenario := "scenario 1"
	encScenario := base64.StdEncoding.EncodeToString([]byte(scenario))

	reqBody, _ := json.Marshal(map[string]string{
		"Name":     name,
		"Scenario": encScenario,
	})
	req, _ := http.NewRequest("POST", "/api/applications", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var app ent.Application
	_ = json.Unmarshal(w.Body.Bytes(), &app)

	return &app
}

func TestLogin(t *testing.T) {
	adminUsername := "admin"
	adminPassword := "adminPassword"

	tt := []struct {
		input []string
		want  int
	}{
		{[]string{adminUsername, "not a password"}, 401},
		{[]string{"not an admin", adminPassword}, 401},
		{[]string{adminUsername, adminPassword}, 200},
	}

	for _, tc := range tt {
		r, w := newAPITest(t, adminPassword)

		reqBody, _ := json.Marshal(map[string]string{
			"username": tc.input[0],
			"password": tc.input[1],
		})

		loginReq, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(reqBody))
		loginReq.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, loginReq)
		assert.Equal(t, tc.want, w.Code)
	}
}

func TestAuth401(t *testing.T) {
	r, w := newAPITest(t, "adminPassword")

	req, _ := http.NewRequest("GET", "/api/applications", nil)
	req.Header.Add("Authorization", "Bearer sometoken")
	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestAuth200(t *testing.T) {
	adminPassword := "adminPassword"

	r, w := newAPITest(t, adminPassword)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "admin",
		"password": adminPassword,
	})
	loginReq, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(reqBody))
	loginReq.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, loginReq)

	assert.Equal(t, 200, w.Code)

	act := new(accesstokenResponse)
	json.Unmarshal(w.Body.Bytes(), act)

	req, _ := http.NewRequest("GET", "/api/applications", nil)
	req.Header.Add("Authorization", "Bearer "+act.ID)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestListApplications(t *testing.T) {
	r, w := newAPITest(t, "")
	req, _ := http.NewRequest("GET", "/api/applications", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateApplications(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		r, w := newAPITest(t, "")
		name := "name"
		scenario := "this is the scenario"
		encScenario := base64.StdEncoding.EncodeToString([]byte(scenario))

		reqBody, _ := json.Marshal(map[string]string{
			"Name":     name,
			"Scenario": encScenario,
		})
		req, _ := http.NewRequest("POST", "/api/applications", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)

		var app ent.Application
		json.Unmarshal(w.Body.Bytes(), &app)
		assert.Equal(t, app.Name, name)
		assert.Equal(t, app.Scenario, scenario)
		// assert.Equal(t, app.Status, "pending")
	})

	t.Run("successful request with go module", func(t *testing.T) {
		r, w := newAPITest(t, "")
		name := "name"
		scenario := "this is the scenario"
		gomod := "this is the go.mod"
		gosum := "this is the go.sum"

		reqBody, _ := json.Marshal(map[string]string{
			"Name":     name,
			"Scenario": base64.StdEncoding.EncodeToString([]byte(scenario)),
			"Gomod":    base64.StdEncoding.EncodeToString([]byte(gomod)),
			"Gosum":    base64.StdEncoding.EncodeToString([]byte(gosum)),
		})
		req, _ := http.NewRequest("POST", "/api/applications", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)

		var app ent.Application
		json.Unmarshal(w.Body.Bytes(), &app)
		assert.Equal(t, app.Name, name)
		assert.Equal(t, app.Scenario, scenario)
		assert.Equal(t, app.Gomod, gomod)
		assert.Equal(t, app.Gosum, gosum)
		// assert.Equal(t, app.Status, "pending")

	})

	t.Run("invalid request - without Name", func(t *testing.T) {
		r, w := newAPITest(t, "")
		reqBody, _ := json.Marshal(map[string]string{
			"Scenario": "this is the scenario",
		})
		req, _ := http.NewRequest("POST", "/api/applications", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Contains(
			t,
			w.Body.String(),
			`{"error":{"code":400,"message":"Name required","status":"Invalid Request"}}`,
		)
	})

	t.Run("invalid request - without Scenario", func(t *testing.T) {
		r, w := newAPITest(t, "")
		reqBody, _ := json.Marshal(map[string]string{
			"Name": "name",
		})
		req, _ := http.NewRequest("POST", "/api/applications", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Contains(
			t,
			w.Body.String(),
			`{"error":{"code":400,"message":"Scenario required","status":"Invalid Request"}}`,
		)
	})
}

func TestGetApplication(t *testing.T) {
	t.Run("not found request", func(t *testing.T) {
		r, w := newAPITest(t, "")
		req, _ := http.NewRequest("GET", "/api/applications/not-a-number", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
		assert.Contains(
			t,
			w.Body.String(),
			`{"error":{"code":404,"message":"Request data not found","status":"Model Not Found"}}`,
		)
	})

	t.Run("successful request", func(t *testing.T) {
		app := newApp(t)

		r, w := newAPITest(t, "")
		req, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("/api/applications/%d", app.ID),
			nil,
		)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		var resApp ent.Application
		_ = json.Unmarshal(w.Body.Bytes(), &resApp)
		assert.Equal(t, resApp, *app)
	})
}

func TestCancelApplication(t *testing.T) {
	app := newApp(t)

	r, w := newAPITest(t, "")
	req, _ := http.NewRequest(
		"PUT",
		fmt.Sprintf("/api/applications/%d/cancel", app.ID),
		nil,
	)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var resApp ent.Application
	_ = json.Unmarshal(w.Body.Bytes(), &resApp)
	assert.Equal(t, resApp.Status, "cancel")
}

func TestDeleteApplication(t *testing.T) {
	app := newApp(t)

	r, w := newAPITest(t, "")
	req, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("/api/applications/%d", app.ID),
		nil,
	)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
func TestSetApplicationTag(t *testing.T) {
	app := newApp(t)

	r, w := newAPITest(t, "")

	reqBody, _ := json.Marshal(map[string]string{
		"tags": "test,test2,test3",
	})

	req, _ := http.NewRequest(
		"PATCH",
		fmt.Sprintf("/api/applications/%d/tags", app.ID),
		bytes.NewBuffer(reqBody),
	)

	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetApplicationLogs(t *testing.T) {
	app := newApp(t)

	r, w := newAPITest(t, "")
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("/api/applications/%d/logs/system", app.ID),
		nil,
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	req, _ = http.NewRequest(
		"GET",
		fmt.Sprintf("/api/applications/%d/logs/user", app.ID),
		nil,
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
