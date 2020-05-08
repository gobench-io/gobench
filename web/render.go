package web

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/gobench-io/gobench/ent"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInternalServer(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal Server Error.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
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
