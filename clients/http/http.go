package http

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gobench-io/gobench/v2/executor"
	"github.com/gobench-io/gobench/v2/executor/metrics"
)

type HttpClient struct {
	prefix string
	client *http.Client

	title struct {
		success   string
		fail      string
		otherFail string
		latency   string
	}
}

func NewHttpClient(ctx context.Context, prefix string) (HttpClient, error) {
	httpClient := HttpClient{}

	httpClient.prefix = prefix
	httpClient.title.success = prefix + ".http_ok"
	httpClient.title.otherFail = prefix + ".http_other_fail"
	httpClient.title.fail = prefix + ".http_fail"
	httpClient.title.latency = prefix + ".latency"

	group := metrics.Group{
		Name: "HTTP (" + prefix + ")",
		Graphs: []metrics.Graph{
			{
				Title: "HTTP Response",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: httpClient.title.success,
						Type:  metrics.Counter,
					},
					{
						Title: httpClient.title.fail,
						Type:  metrics.Counter,
					},
					{
						Title: httpClient.title.otherFail,
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: httpClient.title.latency,
						Type:  metrics.Histogram,
					},
				},
			},
		},
	}
	groups := []metrics.Group{
		group,
	}

	tr := &http.Transport{
		MaxIdleConnsPerHost: 300,
	}

	httpClient.client = &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}

	if err := executor.Setup(groups); err != nil {
		return httpClient, err
	}

	return httpClient, nil
}

func (h *HttpClient) do(method, url string, body []byte, headers map[string]string) (
	res *http.Response, err error,
) {
	begin := time.Now()

	defer func() {
		diff := time.Since(begin)
		executor.Notify(h.title.latency, diff.Microseconds())
		if err != nil {
			executor.Notify(h.title.otherFail, 1)
			return
		}
		if res.StatusCode >= 300 || res.StatusCode < 200 {
			executor.Notify(h.title.fail, 1)
			return
		}
		executor.Notify(h.title.success, 1)
	}()

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	// add headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err = h.client.Do(req)

	return
}

func (h *HttpClient) captureRes(verb string, url string, body []byte, headers map[string]string) ([]byte, error) {
	res, err := h.do(verb, url, body, headers)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)

	return buf, err
}

func (h *HttpClient) ignoreRes(verb string, url string, body []byte, headers map[string]string) error {
	res, err := h.do(verb, url, body, headers)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	io.Copy(ioutil.Discard, res.Body)

	return nil
}

// Get makes http get request and record the metrics
func (h *HttpClient) Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	return h.captureRes(http.MethodGet, url, nil, headers)
}

// GetIgnoreRes makes http get request, records the metrics, but ignore the
// responding body. Use this when you need high speed traffic generation
func (h *HttpClient) GetIgnoreRes(ctx context.Context, url string, headers map[string]string) error {
	return h.ignoreRes(http.MethodGet, url, nil, headers)
}

// Post makes http post request and record the metrics
func (h *HttpClient) Post(ctx context.Context, url string, body []byte, headers map[string]string) ([]byte, error) {
	return h.captureRes(http.MethodPost, url, body, headers)
}

// PostIgnoreRes makes http get request, records the metrics, but ignore the
// responding body. Use this when you need high speed traffic generation
func (h *HttpClient) PostIgnoreRes(ctx context.Context, url string, body []byte, headers map[string]string) error {
	return h.ignoreRes(http.MethodPost, url, body, headers)
}

// Put makes http put request and record the metrics
func (h *HttpClient) Put(ctx context.Context, url string, body []byte, headers map[string]string) ([]byte, error) {
	return h.captureRes(http.MethodPut, url, body, headers)
}

// Patch makes http patch request and record the metrics
func (h *HttpClient) Patch(ctx context.Context, url string, body []byte, headers map[string]string) ([]byte, error) {
	return h.captureRes(http.MethodPatch, url, body, headers)
}
