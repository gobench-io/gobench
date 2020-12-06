package http

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gobench-io/gobench/executor"
	"github.com/gobench-io/gobench/executor/metrics"
)

type HttpClient struct {
	prefix string
	client *http.Client
}

func NewHttpClient(ctx context.Context, prefix string) (HttpClient, error) {
	group := metrics.Group{
		Name: "HTTP (" + prefix + ")",
		Graphs: []metrics.Graph{
			{
				Title: "HTTP Response",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: prefix + ".http_ok",
						Type:  metrics.Counter,
					},
					{
						Title: prefix + ".http_fail",
						Type:  metrics.Counter,
					},
					{
						Title: prefix + ".http_other_fail",
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: prefix + ".latency",
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

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}

	httpClient := HttpClient{
		prefix: prefix,
		client: client,
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
	otherFail := h.prefix + ".http_other_fail"
	fail := h.prefix + ".http_fail"
	success := h.prefix + ".http_ok"
	latency := h.prefix + ".latency"

	defer func() {
		diff := time.Since(begin)
		executor.Notify(latency, diff.Microseconds())
		if err != nil {
			executor.Notify(otherFail, 1)
			return
		}
		if res.StatusCode >= 300 || res.StatusCode < 200 {
			executor.Notify(fail, 1)
			err = fmt.Errorf("request failed with status code %d", res.StatusCode)
			return
		}
		executor.Notify(success, 1)
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

	res.Body.Close()

	return nil
}

// Get makes http get request and record the metrics
func (h *HttpClient) Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	return h.captureRes("GET", url, nil, headers)
}

// GetIgnoreRes makes http get request, records the metrics, but ignore the
// responding body. Use this when you need high speed traffic generation
func (h *HttpClient) GetIgnoreRes(ctx context.Context, url string, headers map[string]string) error {
	return h.ignoreRes("GET", url, nil, headers)
}

// Post makes http post request and record the metrics
func (h *HttpClient) Post(ctx context.Context, url string, body []byte, headers map[string]string) ([]byte, error) {
	return h.captureRes("POST", url, body, headers)
}

// PostIgnoreRes makes http get request, records the metrics, but ignore the
// responding body. Use this when you need high speed traffic generation
func (h *HttpClient) PostIgnoreRes(ctx context.Context, url string, body []byte, headers map[string]string) error {
	return h.ignoreRes("POST", url, body, headers)
}
