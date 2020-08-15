package web

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationEventLog(t *testing.T) {
	app := newApp(t)

	r, w := newAPITest(t)
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("/api/applications/%d/logs", app.ID),
		nil,
	)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
