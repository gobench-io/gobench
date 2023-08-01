package nats

import (
	"context"
	"log"
	"time"

	"github.com/gobench-io/gobench/v2/executor"
	"github.com/gobench-io/gobench/v2/executor/metrics"
	"github.com/nats-io/nats.go"
)

const conTotal string = "nats.connection.current_total"
const conError string = "nats.connection.connect.errors"
const conReconnect string = "nats.connection.reconnects"
const conLatency string = "nats.connection.connack.latency"

const pubLatency string = "nats.publisher.puback.latency"
const pubTotal string = "nats.publisher.puback.in.total"
const pubWaiting string = "nats.publisher.puback.waiting"

const subLatency string = "nats.subscriber.suback.latency"
const unsubLatency string = "nats.subscriber.unsuback.latency"
const subTotal string = "nats.subscriber.current_total"
const subError string = "nats.subscriber.suback.error"
const msgSubTotal string = "nats.message.consumed.total"

const reqLatency string = "nats.req.latency"
const reqOk string = "nats.req.ok"
const reqErr string = "nats.req.err"

func groups() []metrics.Group {
	conGroup := metrics.Group{
		Name: "NAT Connections",
		Graphs: []metrics.Graph{
			{
				Title: "Connections",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: conTotal,
						Type:  metrics.Counter,
					},
					{
						Title: conError,
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Reconnects",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: conReconnect,
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Connack Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: conLatency,
						Type:  metrics.Histogram,
					},
				},
			},
		},
	}
	pubGroup := metrics.Group{
		Name: "NAT Publishers",
		Graphs: []metrics.Graph{
			{
				Title: "Publish to Puback Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: pubLatency,
						Type:  metrics.Histogram,
					},
				},
			},
			{
				Title: "Pubacks Received Total",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: pubTotal,
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Outstanding Pubacks (Waiting Acks)",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: pubWaiting,
						Type:  metrics.Counter,
					},
				},
			},
		},
	}
	consumerGroup := metrics.Group{
		Name: "NATS Consumers",
		Graphs: []metrics.Graph{
			{
				Title: "Suback Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: subLatency,
						Type:  metrics.Histogram,
					},
				},
			},
			{
				Title: "Subscriber Total/Errors",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: subTotal,
						Type:  metrics.Counter,
					},
					{
						Title: subError,
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Unsuback Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: unsubLatency,
						Type:  metrics.Histogram,
					},
				},
			},
			{
				Title: "Total Consumed Message",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: msgSubTotal,
						Type:  metrics.Counter,
					},
				},
			},
		},
	}
	reqGroup := metrics.Group{
		Name: "NAT Request",
		Graphs: []metrics.Graph{
			{
				Title: "Request Number",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: reqOk,
						Type:  metrics.Counter,
					},
					{
						Title: reqErr,
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Request Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: reqLatency,
						Type:  metrics.Histogram,
					},
				},
			},
		},
	}
	return []metrics.Group{
		conGroup,
		pubGroup,
		consumerGroup,
		reqGroup,
	}
}

func disconnectErrHandler(nc *nats.Conn, err error) {
	if err != nil {
		log.Printf("Disconnected due to: %v\n", err)
	}
	executor.Notify(conTotal, -1)
}

func reconnectHandler(nc *nats.Conn) {
	log.Printf("Reconnected [%s]\n", nc.ConnectedUrl())
	executor.Notify(conTotal, 1)
	executor.Notify(conReconnect, 1)
}

func closeHandler(nc *nats.Conn) {
	log.Printf("Exiting: %v\n", nc.LastError())
}

// A Conn represents a bare connection to a nats-server. It can send and receive
// []byte payloads. It is a type definition of nats.Conn
type Conn nats.Conn

// Msg is a structure used by Subscribers and PublishMsg()
type Msg struct {
	nats.Msg
}

// NewNatClient creates a NATS client that connect to given url
func NewNatClient(ctx context.Context, url string) (*Conn, error) {
	if err := executor.Setup(groups()); err != nil {
		return nil, err
	}

	// nats opts
	opts := []nats.Option{nats.Name("NATS Benchmark")}
	opts = append(opts, nats.DisconnectErrHandler(disconnectErrHandler))
	opts = append(opts, nats.ReconnectHandler(reconnectHandler))
	opts = append(opts, nats.ClosedHandler(closeHandler))

	begin := time.Now()
	c, err := nats.Connect(url, opts...)
	if err != nil {
		log.Printf("Connect error: %v\n", err)
		executor.Notify(conError, 1)

		return nil, err
	}

	diff := time.Since(begin)

	executor.Notify(conTotal, 1)
	executor.Notify(conLatency, diff.Microseconds())

	return (*Conn)(c), nil
}

// Publish publishes given data to a topic
func (c *Conn) Publish(ctx context.Context, topic string, data []byte) error {
	begin := time.Now()
	if err := (*nats.Conn)(c).Publish(topic, data); err != nil {
		return err
	}
	diff := time.Since(begin)

	executor.Notify(pubTotal, 1)
	executor.Notify(pubLatency, diff.Microseconds())

	return nil
}

// Subscribe to a topic, given callback function
func (c *Conn) Subscribe(ctx context.Context, topic string, cb func(msg *Msg)) error {
	ch := make(chan *nats.Msg, 1)
	begin := time.Now()

	if _, err := (*nats.Conn)(c).ChanSubscribe(topic, ch); err != nil {
		executor.Notify(subError, 1)
		return err
	}
	diff := time.Since(begin)

	// notify sub total and latency
	executor.Notify(subTotal, 1)
	executor.Notify(subLatency, diff.Microseconds())

	go func(ch chan *nats.Msg) {
		for {
			select {
			case msg := <-ch:
				executor.Notify(msgSubTotal, 1)
				if cb != nil {
					cb(&Msg{*msg})
				}
			}
		}
	}(ch)

	return nil
}

// QueueSubscribe to a topic, given callback function
func (c *Conn) QueueSubscribe(ctx context.Context, sub, group string, cb func(msg *Msg)) error {
	ch := make(chan *nats.Msg, 1)
	begin := time.Now()

	if _, err := (*nats.Conn)(c).ChanQueueSubscribe(sub, group, ch); err != nil {
		executor.Notify(subError, 1)
		return err
	}
	diff := time.Since(begin)

	// notify sub total and latency
	executor.Notify(subTotal, 1)
	executor.Notify(subLatency, diff.Microseconds())

	go func(ch chan *nats.Msg) {
		for {
			select {
			case msg := <-ch:
				executor.Notify(msgSubTotal, 1)
				if cb != nil {
					cb(&Msg{*msg})
				}
			}
		}
	}(ch)

	return nil
}

// Request will send a request payload and deliver the response message, or an
// error, including a timeout if no message was received correctly
func (c *Conn) Request(ctx context.Context, sub string, data []byte,
	timeout time.Duration) (*Msg, error) {
	begin := time.Now()

	m, err := (*nats.Conn)(c).Request(sub, data, timeout)

	diff := time.Since(begin)

	executor.Notify(reqLatency, diff.Microseconds())
	if err != nil {
		executor.Notify(reqErr, 1)
		return nil, err
	}

	executor.Notify(reqOk, 1)

	return &Msg{*m}, nil
}

func (c *Conn) Flush() error {
	return (*nats.Conn)(c).Flush()
}

// Disconnect drains and closes the connection
func (c *Conn) Disconnect(ctx context.Context) error {
	(*nats.Conn)(c).Drain()
	(*nats.Conn)(c).Close()

	return nil
}
