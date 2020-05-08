package mqtt

import (
	"context"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gobench-io/gobench"
	"github.com/gobench-io/gobench/metrics"
)

const conTotal string = "mqtt.connection.current_total"
const conError string = "mqtt.connection.connect.errors"
const conReconnect string = "mqtt.connection.reconnects"
const conLatency string = "mqtt.connection.connack.latency"

const pubQos0Latency string = "mqtt.publisher.qos0.puback.latency"
const pubQos0Total string = "mqtt.publisher.qos0.puback.in.total"
const pubQos1Latency string = "mqtt.publisher.qos1.puback.latency"
const pubQos1Total string = "mqtt.publisher.qos1.puback.in.total"
const pubQos2Latency string = "mqtt.publisher.qos2.puback.latency"
const pubQos2Total string = "mqtt.publisher.qos2.puback.in.total"

const subLatency string = "mqtt.subscriber.suback.latency"
const subTotal string = "mqtt.subscriber.current_total"
const subError string = "mqtt.subscriber.suback.error"
const unsubLatency string = "mqtt.subscriber.unsuback.latency"
const unsubError string = "mqtt.subscriber.unsuback.error"

// ContextKey is the type for context
type ContextKey string

type ClientOptions struct {
	*mqtt.ClientOptions
}

func NewClientOptions() *ClientOptions {
	t := mqtt.NewClientOptions()
	o := &ClientOptions{ClientOptions: t}
	return o
}

type MqttClient struct {
	client mqtt.Client
}

func groups() []metrics.Group {
	conGroup := metrics.Group{
		Name: "MQTT Connections",
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
	qos0PubGroup := metrics.Group{
		Name: "MQTT Publishers QoS 0",
		Graphs: []metrics.Graph{
			{
				Title: "QoS0 Publish to Puback Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: pubQos0Latency,
						Type:  metrics.Histogram,
					},
				},
			},
			{
				Title: "QoS0 Pubacks Received Total",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: pubQos0Total,
						Type:  metrics.Counter,
					},
				},
			},
		},
	}
	qos1PubGroup := metrics.Group{
		Name: "MQTT Publishers QoS 1",
		Graphs: []metrics.Graph{
			{
				Title: "QoS1 Publish to Puback Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: pubQos1Latency,
						Type:  metrics.Histogram,
					},
				},
			},
			{
				Title: "QoS1 Pubacks Received Total",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: pubQos1Total,
						Type:  metrics.Counter,
					},
				},
			},
		},
	}
	qos2PubGroup := metrics.Group{
		Name: "MQTT Publishers QoS 2",
		Graphs: []metrics.Graph{
			{
				Title: "QoS2 Publish to Puback Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: pubQos2Latency,
						Type:  metrics.Histogram,
					},
				},
			},
			{
				Title: "QoS2 Pubacks Received Total",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: pubQos2Total,
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
						Title: unsubError,
						Type:  metrics.Histogram,
					},
				},
			},
			{
				Title: "Subscriber Unsuback Errors",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: unsubError,
						Type:  metrics.Counter,
					},
				},
			},
		},
	}
	return []metrics.Group{
		conGroup,
		qos0PubGroup,
		qos1PubGroup,
		qos2PubGroup,
		consumerGroup,
	}
}

func Init(client mqtt.Client) MqttClient {
	mqttClient := MqttClient{client}
	mqttClient.client = client

	return mqttClient
}

func NewMqttClient(ctx *context.Context, opts *ClientOptions) (MqttClient, error) {
	mqttClient := MqttClient{}

	gs := groups()
	if err := gobench.Setup(gs); err != nil {
		return mqttClient, err
	}

	clientID := opts.ClientID
	// generate random clientID if not provided
	if clientID == "" {
		clientID = gobench.RandomString(32)
	}

	opts.ClientOptions.SetClientID(clientID)

	// be called when the client is connected.
	// Both at initial connection time and upon automatic reconnect.
	OnConnect := opts.OnConnect
	opts.SetOnConnectHandler(func(c mqtt.Client) {
		gobench.Notify(conTotal, 1)
		if OnConnect != nil {
			OnConnect(c)
		}
	})

	// be executed in the case where the client unexpectedly loses connection with the MQTT broker.
	OnConnectionLost := opts.OnConnectionLost
	opts.SetConnectionLostHandler(func(c mqtt.Client, e error) {
		gobench.Notify(conTotal, -1)
		if OnConnectionLost != nil {
			OnConnectionLost(c, e)
		}
	})

	// be executed prior to the client attempting a reconnect to the MQTT broker.
	OnReconnecting := opts.OnReconnecting
	opts.SetReconnectingHandler(func(c mqtt.Client, o *mqtt.ClientOptions) {
		gobench.Notify(conReconnect, 1)
		if OnReconnecting != nil {
			OnReconnecting(c, o)
		}
	})

	client := mqtt.NewClient(opts.ClientOptions)

	mqttClient.client = client

	return mqttClient, nil
}

func (c *MqttClient) Connect(ctx *context.Context) error {
	begin := time.Now()

	token := c.client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}

	if err := token.Error(); err != nil {
		log.Printf("mqtt connect fail: %s\n", err.Error())

		gobench.Notify(conError, 1)
		return err
	}
	gobench.Notify(conLatency, time.Since(begin).Microseconds())

	return nil
}

func (c *MqttClient) Publish(ctx *context.Context, topic string, qos byte, data []byte) error {
	begin := time.Now()
	token := c.client.Publish(topic, qos, false, data)
	token.WaitTimeout(3 * time.Second)
	if err := token.Error(); err != nil {
		// todo: log the publish error
		log.Printf("mqtt publish fail: %s\n", err.Error())
		return err
	}

	switch qos {
	case 0:
		gobench.Notify(pubQos0Total, 1)
		gobench.Notify(pubQos0Latency, time.Since(begin).Microseconds())
	case 1:
		gobench.Notify(pubQos1Total, 1)
		gobench.Notify(pubQos1Latency, time.Since(begin).Microseconds())
	case 2:
		gobench.Notify(pubQos2Total, 1)
		gobench.Notify(pubQos2Latency, time.Since(begin).Microseconds())
	}

	return nil
}

func (c *MqttClient) Subscribe(ctx *context.Context, topic string, qos byte) error {
	begin := time.Now()
	token := c.client.Subscribe(topic, qos, nil)
	token.WaitTimeout(3 * time.Second)

	if err := token.Error(); err != nil {
		log.Printf("mqtt subscribe fail: %s\n", err.Error())
		gobench.Notify(subError, 1)
		return err
	}

	gobench.Notify(subLatency, time.Since(begin).Microseconds())
	gobench.Notify(subTotal, 1)

	return nil
}

func (c *MqttClient) Unsubscribe(ctx *context.Context, topics ...string) error {
	begin := time.Now()

	token := c.client.Unsubscribe(topics...)
	token.WaitTimeout(3 * time.Second)

	if err := token.Error(); err != nil {
		log.Printf("mqtt unsubscribe fail: %s\n", err.Error())
		gobench.Notify(unsubError, 1)
		return err
	}

	gobench.Notify(unsubLatency, time.Since(begin).Microseconds())

	return nil
}

func (c *MqttClient) Disconnect(ctx *context.Context) error {
	c.client.Disconnect(500)
	gobench.Notify(conTotal, -1)
	return nil
}
