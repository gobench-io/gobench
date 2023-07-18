package web

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/gobench-io/gobench/v2/ent"
	"github.com/gobench-io/gobench/v2/logger"
	"github.com/gobench-io/gobench/v2/master"
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

func cleanApplication(t *testing.T) {
	r, w := newAPITest(t, "")
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/applications?limit=%d", ^uint(0)), nil)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	assert.Less(t, w.Code, 400)

	var apps []ent.Application
	err := json.Unmarshal(w.Body.Bytes(), &apps)
	assert.Equal(t, err, nil)
	for _, app := range apps {
		r, w := newAPITest(t, "")
		req, _ := http.NewRequest(
			"DELETE",
			fmt.Sprintf("/api/applications/%d", app.ID),
			nil,
		)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	}
}
func newAppTag(t *testing.T, appID int, name string) *ent.Tag {
	r, w := newAPITest(t, "")
	reqBody, _ := json.Marshal(map[string]string{
		"Name": name,
	})
	req, _ := http.NewRequest("PUT",
		fmt.Sprintf("/api/applications/%d/tags", appID),
		bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var tag ent.Tag
	_ = json.Unmarshal(w.Body.Bytes(), &tag)

	return &tag
}

func newApp(t *testing.T, name string, scenario string) *ent.Application {
	r, w := newAPITest(t, "")
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

func Test_handler_countApplications(t *testing.T) {
	// init test
	cleanApplication(t)
	_ = newApp(t, "foo", "foo")
	_ = newApp(t, "foo1", "foo1")

	type args struct {
		keyword string
	}
	type wantData struct {
		Count int
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantData wantData
	}{
		{
			name:     "should return success when have no params",
			args:     args{},
			wantErr:  false,
			wantData: wantData{Count: 2},
		},
		{
			name:     "should return success when have a keyword match",
			args:     args{keyword: "foo"},
			wantErr:  false,
			wantData: wantData{Count: 2},
		},
		{
			name:     "should return success when have a keyword match again",
			args:     args{keyword: "foo1"},
			wantErr:  false,
			wantData: wantData{Count: 1},
		},
		{
			name:     "should return success when have a keyword does not match",
			args:     args{keyword: "bar"},
			wantErr:  false,
			wantData: wantData{Count: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w := newAPITest(t, "")
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/applications/count?keyword=%s", tt.args.keyword), nil)
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)
			assert.Equal(t, tt.wantErr, w.Code >= 400)
			if !tt.wantErr {
				assert.Equal(t, w.Code, 200)
				result := wantData{}
				json.Unmarshal(w.Body.Bytes(), &result)
				assert.Equal(t, tt.wantData.Count, result.Count)
			}
		})
	}
}
func Test_handler_listApplications(t *testing.T) {
	// init test
	cleanApplication(t)
	_ = newApp(t, "foo", "foo")
	_ = newApp(t, "foo1", "foo1")

	type args struct {
		keyword string
		limit   int
		offset  int
		order   string
		isAsc   bool
	}
	type wantData []struct {
		Name     string
		Scenario string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantData wantData
	}{
		{
			name:    "should return success when have no params",
			args:    args{},
			wantErr: false,
			wantData: wantData{
				{
					Name:     "foo1",
					Scenario: "foo1",
				},
				{
					Name:     "foo",
					Scenario: "foo",
				},
			},
		},
		{
			name:    "should return success when have a keyword match",
			args:    args{keyword: "foo"},
			wantErr: false,
			wantData: wantData{
				{
					Name:     "foo1",
					Scenario: "foo1",
				},
				{
					Name:     "foo",
					Scenario: "foo",
				},
			},
		},
		{
			name:    "should return success when have a keyword match again",
			args:    args{keyword: "foo1"},
			wantErr: false,
			wantData: wantData{
				{
					Name:     "foo1",
					Scenario: "foo1",
				},
			},
		},
		{
			name:     "should return success when have a keyword does not match",
			args:     args{keyword: "foo1111"},
			wantErr:  false,
			wantData: wantData{},
		},
		{
			name:    "should return success when have limit is set 1",
			args:    args{limit: 1},
			wantErr: false,
			wantData: wantData{
				{
					Name:     "foo1",
					Scenario: "foo1",
				},
			},
		},
		{
			name:    "should return success when have offset is set 1",
			args:    args{offset: 1},
			wantErr: false,
			wantData: wantData{
				{
					Name:     "foo",
					Scenario: "foo",
				},
			},
		},
		{
			name:    "should return success when have isAsc is set true",
			args:    args{isAsc: true},
			wantErr: false,
			wantData: wantData{
				{
					Name:     "foo",
					Scenario: "foo",
				},
				{
					Name:     "foo1",
					Scenario: "foo1",
				},
			},
		},
		{
			name:    "should return success when have order is set name and isAsc is true",
			args:    args{order: "name", isAsc: true},
			wantErr: false,
			wantData: wantData{
				{
					Name:     "foo",
					Scenario: "foo",
				},
				{
					Name:     "foo1",
					Scenario: "foo1",
				},
			},
		},
		{
			name:    "should return success when have order is set name and isAsc is false",
			args:    args{order: "name", isAsc: false},
			wantErr: false,
			wantData: wantData{
				{
					Name:     "foo1",
					Scenario: "foo1",
				},
				{
					Name:     "foo",
					Scenario: "foo",
				},
			},
		},
		{
			name:    "should return success when have order is set name and isAsc is false, limit is 1 and offset is 1",
			args:    args{order: "name", isAsc: false, limit: 1, offset: 1},
			wantErr: false,
			wantData: wantData{
				{
					Name:     "foo",
					Scenario: "foo",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w := newAPITest(t, "")
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/applications?keyword=%s&limit=%d&offset=%d&order=%s&isAsc=%v", tt.args.keyword, tt.args.limit, tt.args.offset, tt.args.order, tt.args.isAsc), nil)
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)
			assert.Equal(t, tt.wantErr, w.Code >= 400)
			if !tt.wantErr {
				assert.Equal(t, w.Code, 200)
				result := wantData{}
				json.Unmarshal(w.Body.Bytes(), &result)
				assert.Equal(t, len(tt.wantData), len(result))
				for k, v := range tt.wantData {
					assert.Equal(t, tt.wantData[k].Name, v.Name)
					assert.Equal(t, tt.wantData[k].Scenario, v.Scenario)
				}
			}
		})
	}
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
		app := newApp(t, "name 1", "scenario 1")

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
	app := newApp(t, "name 1", "scenario 1")

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
	app := newApp(t, "name 1", "scenario 1")

	r, w := newAPITest(t, "")
	req, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("/api/applications/%d", app.ID),
		nil,
	)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestAddApplicationTag(t *testing.T) {
	tagName := "foo"
	tag := &ent.Tag{}
	app := newApp(t, "name 1", "scenario 1")

	r, w := newAPITest(t, "")

	reqBody, _ := json.Marshal(map[string]string{
		"name": tagName,
	})

	req, _ := http.NewRequest(
		"PUT",
		fmt.Sprintf("/api/applications/%d/tags", app.ID),
		bytes.NewBuffer(reqBody),
	)

	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	_ = json.Unmarshal(w.Body.Bytes(), &tag)
	assert.Equal(t, tagName, tag.Name)
	assert.Equal(t, 200, w.Code)
}

func TestAddApplicationTagAgain(t *testing.T) {
	tagName := "foo"
	tag := &ent.Tag{}
	app := newApp(t, "name 1", "scenario 1")
	_ = newAppTag(t, app.ID, tagName)

	r, w := newAPITest(t, "")

	reqBody, _ := json.Marshal(map[string]string{
		"name": tagName,
	})

	req, _ := http.NewRequest(
		"PUT",
		fmt.Sprintf("/api/applications/%d/tags", app.ID),
		bytes.NewBuffer(reqBody),
	)

	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	err := json.Unmarshal(w.Body.Bytes(), &tag)
	assert.Nil(t, err)
	assert.Equal(t, tagName, tag.Name)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteApplicationTag(t *testing.T) {
	tagName := "foo"
	app := newApp(t, "name 1", "scenario 1")
	fooTag := newAppTag(t, app.ID, tagName)

	r, w := newAPITest(t, "")

	req, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("/api/applications/%d/tags/%d", app.ID, fooTag.ID),
		bytes.NewBuffer(nil),
	)

	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetApplicationLogs(t *testing.T) {
	app := newApp(t, "name 1", "scenario 1")

	// create 2 logs
	fd := fmt.Sprintf("/tmp/applications/%d", app.ID)
	ul := fmt.Sprintf("/tmp/applications/%d/user.log", app.ID)
	sl := fmt.Sprintf("/tmp/applications/%d/system.log", app.ID)
	err := exec.Command("mkdir", fd).Run()
	assert.Nil(t, err)
	err = exec.Command("sh", "-c", "echo user >"+ul).Run()
	assert.Nil(t, err)
	err = exec.Command("sh", "-c", "echo system >"+sl).Run()
	assert.Nil(t, err)

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

func TestGetHeathz(t *testing.T) {
	r, w := newAPITest(t, "adminPassword")

	healthzReq, _ := http.NewRequest("GET", "/healthz", nil)
	healthzReq.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, healthzReq)
	assert.Equal(t, 200, w.Code)
}

func TestGetVarz(t *testing.T) {
	r, w := newAPITest(t, "adminPassword")

	varzReq, _ := http.NewRequest("GET", "/varz", nil)
	varzReq.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, varzReq)
	assert.Equal(t, 200, w.Code)

	var vr varzResponse
	err := json.Unmarshal(w.Body.Bytes(), &vr)
	assert.Nil(t, err)

	// Do some sanity checks on values
	if time.Since(vr.Start) > 10*time.Second {
		assert.Fail(t, "Expected start time to be within 10 seconds.")
	}
	if vr.ID == "" {
		assert.Fail(t, "Expect server_id to be valid")
	}
	// if vr.Version == "" {
	// 	assert.Fail(t, "Expect version to be valid")
	// }
	if vr.GoVersion == "" {
		assert.Fail(t, "Expect Go version to be valid")
	}
	if vr.Uptime == "" {
		assert.Fail(t, "Expect uptime to be valid")
	}
	if vr.Mem == 0 {
		assert.Fail(t, "Expect mem usage to be valid")
	}
	if vr.Cores == 0 {
		assert.Fail(t, "Expect cores to be valid")
	}
	if vr.MaxProcs == 0 {
		assert.Fail(t, "Expect gomaxprocs to be valid")
	}
}
