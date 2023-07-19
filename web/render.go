package web

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/master"
)

// Err is the error struct that compatible to Google API recommendation
// https://cloud.google.com/apis/design/errors#error_model
type Err struct {
	Code    int    `json:"code,omitempty"`    // application-specific error code
	Message string `json:"message,omitempty"` // application-level error message, for debugging
	Status  string `json:"status"`            // user-level status message
}

// ErrResponse is the error struct that compatible to Google API recommendation
// https://cloud.google.com/apis/design/errors#error_model
type ErrResponse struct {
	Error Err `json:"error"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Error.Code)
	return nil
}

func ErrInternalServer(err error) render.Renderer {
	return &ErrResponse{
		Error: Err{
			Code:    500,
			Message: err.Error(),
			Status:  "Internal Server Error",
		},
	}
}

func ErrUnauthenticated(err error) render.Renderer {
	return &ErrResponse{
		Error: Err{
			Code:    401,
			Message: err.Error(),
			Status:  "Unauthenticated",
		},
	}
}

func ErrAppIsFinished(err error) render.Renderer {
	return &ErrResponse{
		Error: Err{
			Code:    400,
			Message: err.Error(),
			Status:  "Application Finished",
		},
	}
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Error: Err{
			Code:    400,
			Message: err.Error(),
			Status:  "Invalid Request",
		},
	}
}

func ErrNotFoundRequest(err error) render.Renderer {
	return &ErrResponse{
		Error: Err{
			Code:    404,
			Message: "Request data not found",
			Status:  "Model Not Found",
		},
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Error: Err{
			Code:    422,
			Message: err.Error(),
			Status:  "Error Rendering Response.",
		},
	}
}

// application response
type applicationRequest struct {
	*ent.Application
	ProtectedID int `json:"id"`
}

func (a *applicationRequest) Bind(r *http.Request) (err error) {
	return nil
}

type applicationResponse struct {
	*ent.Application
	Edges *struct{} `json:"edges,omitempty"`
}

type countApplicationResponse struct {
	Count int `json:"count"`
}

func (ar *applicationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ar *countApplicationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newCountApplicationResponse(count int) *countApplicationResponse {
	return &countApplicationResponse{
		Count: count,
	}
}

func newApplicationResponse(a *ent.Application) *applicationResponse {
	return &applicationResponse{
		a,
		nil,
	}
}

func newApplicationListResponse(aps []*ent.Application) []render.Renderer {
	list := []render.Renderer{}
	for _, ap := range aps {
		list = append(list, newApplicationResponse(ap))
	}
	return list
}

// group response
type groupResponse struct {
	*ent.Group
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *groupResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newGroupResponse(g *ent.Group) *groupResponse {
	return &groupResponse{
		g,
		nil,
	}
}

func newGroupListResponse(gs []*ent.Group) []render.Renderer {
	list := []render.Renderer{}
	for _, g := range gs {
		list = append(list, newGroupResponse(g))
	}
	return list
}

// graph response
type graphResponse struct {
	*ent.Graph
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *graphResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newGraphResponse(g *ent.Graph) *graphResponse {
	return &graphResponse{
		g,
		nil,
	}
}

func newGraphListResponse(gs []*ent.Graph) []render.Renderer {
	list := []render.Renderer{}
	for _, g := range gs {
		list = append(list, newGraphResponse(g))
	}
	return list
}

// metric response
type metricResponse struct {
	*ent.Metric
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *metricResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newMetricResponse(m *ent.Metric) *metricResponse {
	return &metricResponse{
		m,
		nil,
	}
}

func newMetricListResponse(ms []*ent.Metric) []render.Renderer {
	list := []render.Renderer{}
	for _, m := range ms {
		list = append(list, newMetricResponse(m))
	}
	return list
}

// counter response
type counterResponse struct {
	*ent.Counter
	Count int64     `json:"count"`
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *counterResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newCounterResponse(e *ent.Counter) *counterResponse {
	return &counterResponse{
		e,
		e.Count,
		nil,
	}
}

func newCounterListResponse(es []*ent.Counter) []render.Renderer {
	list := []render.Renderer{}
	for _, e := range es {
		list = append(list, newCounterResponse(e))
	}
	return list
}

// histogram response
type histogramResponse struct {
	*ent.Histogram
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *histogramResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newHistogramResponse(e *ent.Histogram) *histogramResponse {
	return &histogramResponse{
		e,
		nil,
	}
}

func newHistogramListResponse(es []*ent.Histogram) []render.Renderer {
	list := []render.Renderer{}
	for _, e := range es {
		list = append(list, newHistogramResponse(e))
	}
	return list
}

// gauge response
type gaugeResponse struct {
	*ent.Gauge
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *gaugeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newGaugeResponse(e *ent.Gauge) *gaugeResponse {
	return &gaugeResponse{
		e,
		nil,
	}
}

func newGaugeListResponse(es []*ent.Gauge) []render.Renderer {
	list := []render.Renderer{}
	for _, e := range es {
		list = append(list, newGaugeResponse(e))
	}
	return list
}

type accesstokenRequest struct {
	Username string
	Password string
}

func (act *accesstokenRequest) Bind(r *http.Request) (err error) {
	return nil
}

// access token response
type accesstokenResponse struct {
	ID string `json:"id"`
}

func (act *accesstokenResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type tagRequest struct {
	*ent.Tag
	ProtectedID int `json:"id"`
}

func (a *tagRequest) Bind(r *http.Request) (err error) {
	return nil
}

type tagResponse struct {
	*ent.Tag
	Edges *struct{} `json:"edges,omitempty"`
}

func newTagResponse(t *ent.Tag) *tagResponse {
	return &tagResponse{
		t,
		nil,
	}
}

func newTagListResponse(tags []*ent.Tag) []render.Renderer {
	list := []render.Renderer{}
	for _, tag := range tags {
		list = append(list, newTagResponse(tag))
	}
	return list
}

func (ar *tagResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// varz response
type varzResponse struct {
	master.Varz
}

func (vr *varzResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
