package gbGrpc

import (
	"context"
	"io"
	"time"

	"github.com/gobench-io/gobench/v2/executor"
	"github.com/gobench-io/gobench/v2/executor/metrics"
	"google.golang.org/grpc"
)

type GbClientConn struct {
	*grpc.ClientConn
	methodGraphsMap map[string][]metrics.Graph
	target          string
}

func Dial(target string, opts ...grpc.DialOption) (*GbClientConn, error) {
	return DialContext(context.Background(), target, opts...)
}

func DialContext(ctx context.Context, target string, opts ...grpc.DialOption) (
	conn *GbClientConn, err error,
) {
	conn = &GbClientConn{
		methodGraphsMap: make(map[string][]metrics.Graph),
		target:          target,
	}

	conn.ClientConn, err = grpc.DialContext(ctx, target, opts...)

	return
}

func (cc *GbClientConn) setupMethod(method string) ([]metrics.Graph, error) {
	if graphs, ok := cc.methodGraphsMap[method]; ok {
		return graphs, nil
	}

	graphs := []metrics.Graph{
		{
			Title: "gRPC Response",
			Unit:  "N",
			Metrics: []metrics.Metric{
				{
					Title: method + ".grpc_ok", // success
					Type:  metrics.Counter,
				},
				{
					Title: method + ".grpc_fail", // fail
					Type:  metrics.Counter,
				},
			},
		},
		{
			Title: "Latency",
			Unit:  "Microsecond",
			Metrics: []metrics.Metric{
				{
					Title: method + ".latency", // latency
					Type:  metrics.Histogram,
				},
			},
		},
	}

	cc.methodGraphsMap[method] = graphs

	group := metrics.Group{
		Name:   "gRPC (" + cc.target + ")",
		Graphs: graphs,
	}

	groups := []metrics.Group{
		group,
	}

	err := executor.Setup(groups)

	return graphs, err
}

func (cc *GbClientConn) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	graphs, err := cc.setupMethod(method)
	if err != nil {
		return err
	}

	begin := time.Now()
	err = cc.ClientConn.Invoke(ctx, method, args, reply, opts...)
	diff := time.Since(begin)

	latencyTitle := graphs[1].Metrics[0].Title
	countTitle := graphs[0].Metrics[0].Title
	if err != nil {
		countTitle = graphs[0].Metrics[1].Title
	}

	executor.Notify(latencyTitle, diff.Microseconds())
	executor.Notify(countTitle, 1)

	return err
}

type GbClientStream struct {
	grpc.ClientStream
	target string
	method string
	graphs []metrics.Graph
}

func (cs *GbClientStream) setupMethod(target string, method string) (
	[]metrics.Graph, error,
) {
	cs.target = target
	cs.method = method

	graphs := []metrics.Graph{
		{
			Title: "New Stream",
			Unit:  "N",
			Metrics: []metrics.Metric{
				{
					Title: method + ".new_stream_ok", // success
					Type:  metrics.Counter,
				},
				{
					Title: method + ".new_stream_fail", // fail
					Type:  metrics.Counter,
				},
			},
		},
		{
			Title: "New Stream Latency",
			Unit:  "Microsecond",
			Metrics: []metrics.Metric{
				{
					Title: method + ".new_stream_latency", // latency
					Type:  metrics.Histogram,
				},
			},
		},
		{
			Title: "Send Message",
			Unit:  "N",
			Metrics: []metrics.Metric{
				{
					Title: method + ".send_msg_ok", // success
					Type:  metrics.Counter,
				},
				{
					Title: method + ".send_msg_fail", // fail
					Type:  metrics.Counter,
				},
			},
		},
		{
			Title: "Send Message Latency",
			Unit:  "Microsecond",
			Metrics: []metrics.Metric{
				{
					Title: method + ".send_msg_latency", // latency
					Type:  metrics.Histogram,
				},
			},
		},
		{
			Title: "Receive Message",
			Unit:  "N",
			Metrics: []metrics.Metric{
				{
					Title: method + ".recv_msg_ok", // success
					Type:  metrics.Counter,
				},
				{
					Title: method + ".recv_msg_fail", // fail
					Type:  metrics.Counter,
				},
			},
		},
		{
			Title: "Receive Message Latency",
			Unit:  "Microsecond",
			Metrics: []metrics.Metric{
				{
					Title: method + ".recv_msg_latency", // latency
					Type:  metrics.Histogram,
				},
			},
		},
	}

	group := metrics.Group{
		Name:   "gRPC Stream (" + cs.target + ")",
		Graphs: graphs,
	}

	groups := []metrics.Group{
		group,
	}

	cs.graphs = graphs

	err := executor.Setup(groups)

	return graphs, err
}

func (cc *GbClientConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (
	grpc.ClientStream, error,
) {
	gcn := &GbClientStream{}

	var err error

	graphs, err := gcn.setupMethod(cc.target, method)
	if err != nil {
		return gcn, err
	}

	begin := time.Now()
	gcn.ClientStream, err = cc.ClientConn.NewStream(ctx, desc, method, opts...)
	diff := time.Since(begin)

	latencyTitle := graphs[1].Metrics[0].Title
	countTitle := graphs[0].Metrics[0].Title
	if err != nil {
		countTitle = graphs[0].Metrics[1].Title
	}

	executor.Notify(latencyTitle, diff.Microseconds())
	executor.Notify(countTitle, 1)

	return gcn, err
}

func (cs *GbClientStream) SendMsg(m interface{}) error {
	begin := time.Now()
	err := cs.ClientStream.SendMsg(m)
	diff := time.Since(begin)

	latencyTitle := cs.graphs[3].Metrics[0].Title
	countTitle := cs.graphs[2].Metrics[0].Title
	if err != nil {
		countTitle = cs.graphs[2].Metrics[1].Title
	}

	executor.Notify(latencyTitle, diff.Microseconds())
	executor.Notify(countTitle, 1)

	return err
}

func (cs *GbClientStream) RecvMsg(m interface{}) error {
	begin := time.Now()
	err := cs.ClientStream.RecvMsg(m)
	diff := time.Since(begin)

	latencyTitle := cs.graphs[5].Metrics[0].Title
	countTitle := cs.graphs[4].Metrics[0].Title
	// could be EOF when the stream ended which is not an error
	if err != nil && err != io.EOF {
		countTitle = cs.graphs[4].Metrics[1].Title
	}

	executor.Notify(latencyTitle, diff.Microseconds())
	executor.Notify(countTitle, 1)

	return err
}
