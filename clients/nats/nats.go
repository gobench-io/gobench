package nats

import (
	"context"
	"log"
	"time"

	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/worker"
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
const subTotal string = "nats.subsciber.current_total"
const subError string = "nats.subsciber.suback.error"

func groups() []metrics.Group {
	conGroup := metrics.Group{
		Name: "NAT Connections",
		Graphs: []metrics.Graph{
			{
				Title: "Total Connections",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: conTotal,
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Connection Errors",
				Unit:  "N",
				Metrics: []metrics.Metric{
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
		Name: "MQTT Consumers",
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
				Title: "Subscriber Total",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: subTotal,
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Subscriber Suback Errors",
				Unit:  "N",
				Metrics: []metrics.Metric{
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
		},
	}
	return []metrics.Group{
		conGroup,
		pubGroup,
		consumerGroup,
	}
}

func disconnectErrHandler(nc *nats.Conn, err error) {
	if err != nil {
		log.Printf("Disconnected due to: %v\n", err)
	}
	worker.Notify(conTotal, -1)
}

func reconnectHandler(nc *nats.Conn) {
	log.Printf("Reconnected [%s]\n", nc.ConnectedUrl())
	worker.Notify(conTotal, 1)
	worker.Notify(conReconnect, 1)
}
func closeHandler(nc *nats.Conn) {
	log.Printf("Exiting: %v\n", nc.LastError())
}

type NatsClient struct {
	conn *nats.Conn
}

func NewNatClient(ctx context.Context, url string) (natsClient NatsClient, err error) {
	if err := worker.Setup(groups()); err != nil {
		return natsClient, err
	}

	// nats opts
	opts := []nats.Option{nats.Name("NATS Benchmark")}
	opts = append(opts, nats.DisconnectErrHandler(disconnectErrHandler))
	opts = append(opts, nats.ReconnectHandler(reconnectHandler))
	opts = append(opts, nats.ClosedHandler(closeHandler))

	begin := time.Now()
	if natsClient.conn, err = nats.Connect(url, opts...); err != nil {
		log.Printf("Connect error: %v\n", err)
		worker.Notify(conError, 1)
		return natsClient, err
	}

	diff := time.Since(begin)

	worker.Notify(conTotal, 1)
	worker.Notify(conLatency, diff.Microseconds())

	return natsClient, nil
}

func (c *NatsClient) Publish(ctx context.Context, topic string, data []byte) error {
	begin := time.Now()
	if err := c.conn.Publish(topic, data); err != nil {
		return err
	}
	diff := time.Since(begin)

	worker.Notify(pubTotal, 1)
	worker.Notify(pubLatency, diff.Microseconds())

	return nil
}

func (c *NatsClient) Subscribe(ctx context.Context, topic string) error {
	ch := make(chan *nats.Msg, 1)
	begin := time.Now()

	// begin to sub, ignore sub
	if _, err := c.conn.ChanSubscribe(topic, ch); err != nil {
		worker.Notify(subError, 1)
		return err
	}
	diff := time.Since(begin)

	// notify sub total and latency
	worker.Notify(subTotal, 1)
	worker.Notify(subLatency, diff.Microseconds())

	go func(ch chan *nats.Msg) {
		<-ch
	}(ch)

	return nil
}

func (c *NatsClient) Disconnect(ctx context.Context) error {
	c.conn.Drain()
	c.conn.Close()

	return nil
}
