# Gobench scenario

Example:

```{go}
func Export() {
    return Scenario{
        vus: []gobench.Vu{
            {
                Number: 100,
                Function: subVuPool,
            }
        }
    }
}

func subVuPool(i int, donewg *sync.WaitGroup) {
	ctx := context.Background()

	defer donewg.Done()

	opts := mqtt.NewClientOptions()
	opts.AddBroker("192.168.2.29:1883")

	client, err := mqtt.NewMqttClient(&ctx, opts)
	if err != nil {
		log.Println(err)
		return
	}

	if err = client.Connect(&ctx); err != nil {
		log.Println(err)
		return
	}

	_ = client.Subscribe(&ctx, "fixed/broadcast/topic", 0, nil)

	// finally
	// _ = client.Disconnect(&ctx)
}

func pubVuPool(i int, donewg *sync.WaitGroup) {
	ctx := context.Background()

	defer donewg.Done()

	opts := mqtt.NewClientOptions()
	opts.AddBroker("192.168.2.29:1883")

	client, err := mqtt.NewMqttClient(&ctx, opts)
	if err != nil {
		log.Println(err)
		return
	}

	if err = client.Connect(&ctx); err != nil {
		log.Println(err)
		return
	}

	rate := 1.0 // rps
	for j := 0; j < int(60*5*rate); j++ {
		gobench.SleepPoisson(rate)
		go func() {
			_ = client.Publish(&ctx, "fixed/broadcast/topic", 0, gobench.RandomByte(150))
		}()
	}

	// finally
	_ = client.Disconnect(&ctx)
}
```
