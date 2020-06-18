package benchclient

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/worker"
)

// cpu
const cpuCount string = "cpu.count"
const cpuCgoCall string = "cpu.cgo_calls"
const cpuGoroutines string = "cpu.goroutines"

// mem gc
const memAlloc string = "mem.alloc"
const memFrees string = "mem.frees"
const memGcCount string = "mem.gc.count"
const memGcLast string = "mem.gc.last"
const memGcNext string = "mem.gc.next"

// const memGcPause string = "mem.gc.pause"
const memGcPauseTotal string = "mem.gc.pause_total"
const memGcSys string = "mem.gc.sys"

const memLookups string = "mem.lookups"
const memMalloc string = "mem.malloc"
const memOthersys string = "mem.othersys"
const memSys string = "mem.sys"
const memTotalalloc string = "mem.totalalloc"

// mem heap
const memHeapAlloc string = "mem.heap.alloc"
const memHeapIdle string = "mem.heap.idle"
const memHeapInuse string = "mem.heap.inuse"
const memHeapObjects string = "mem.heap.objects"
const memHeapReleased string = "mem.heap.released"
const memHeapSys string = "mem.heap.sys"

// mem stack
const memStackInuse string = "mem.stack.inuse"
const memStackSys string = "mem.stack.sys"
const memStackMcacheInuse string = "mem.stack.mcache_inuse"
const memStackMcacheSys string = "mem.stack.mcache_sys"
const memStackMspanInuse string = "mem.stack.mspan_inuse"
const memStackMspanSys string = "mem.stack.mspan_sys"

type InternalClient struct {
	run bool
}

func groups() []metrics.Group {
	cpuGroup := metrics.Group{
		Name: "CPU",
		Graphs: []metrics.Graph{
			metrics.Graph{
				Title: "CPU Cores",
				Unit:  "N",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: cpuCount,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "CGO Calls",
				Unit:  "N",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: cpuCgoCall,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Goroutines",
				Unit:  "N",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: cpuGoroutines,
						Type:  metrics.Gauge,
					},
				},
			},
		},
	}
	memGroup := metrics.Group{
		Name: "Memory",
		Graphs: []metrics.Graph{
			metrics.Graph{
				Title: "Allocated Heap",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memAlloc,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Cumulative Freed Heap",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memFrees,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Completed GC Cycles",
				Unit:  "N",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memGcCount,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Last GC Finished",
				Unit:  "ns (Unix epoch)",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memGcLast,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Target Next GC Cycle",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memGcNext,
						Type:  metrics.Gauge,
					},
				},
			},
			// metrics.Graph{
			// 	Title: "Recent GC Pause",
			// 	Unit:  "ns",
			// 	Metrics: []metrics.Metric{
			// 		metrics.Metric{
			// 			Title:   memGcPause,
			// 			Type: metrics.Gauge,
			// 		},
			// 	},
			// },
			metrics.Graph{
				Title: "Total GC Pause",
				Unit:  "ns",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memGcPauseTotal,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "GC Metadata",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memGcSys,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Pointer Lookups",
				Unit:  "N",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memLookups,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Cumulative Heap Objects Allocated",
				Unit:  "N",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memMalloc,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Miscellaneous off-heap Runtime Allocation",
				Unit:  "N",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memOthersys,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Total Mem from the OS",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memSys,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Cumulative Allocated for Heap Objects",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memTotalalloc,
						Type:  metrics.Gauge,
					},
				},
			},
		},
	}
	heapGroup := metrics.Group{
		Name: "Heap",
		Graphs: []metrics.Graph{
			metrics.Graph{
				Title: "Heap Allocated Objects",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memHeapAlloc,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Heap Idle (unused) Spans",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memHeapIdle,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Heap In-use Spans",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memHeapInuse,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Heap Allocated Object",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memHeapObjects,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Physical Memory Returned to OS",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memHeapReleased,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Heap Memory Obtained from the OS",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memHeapSys,
						Type:  metrics.Gauge,
					},
				},
			},
		},
	}
	memStackGroup := metrics.Group{
		Name: "Mem Stack",
		Graphs: []metrics.Graph{
			metrics.Graph{
				Title: "In-use Stack Spans",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memStackInuse,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Stack Memory Obtained from the OS",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memStackSys,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Allocated mcache Structures",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memStackMcacheInuse,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Memory Obtained from the OS for mcache Structures",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memStackMcacheSys,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Allocated mspan Structures",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memStackMspanInuse,
						Type:  metrics.Gauge,
					},
				},
			},
			metrics.Graph{
				Title: "Memory Obtained from the OS for mspan Structures",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					metrics.Metric{
						Title: memStackMspanSys,
						Type:  metrics.Gauge,
					},
				},
			},
		},
	}

	return []metrics.Group{
		cpuGroup,
		memGroup,
		heapGroup,
		memStackGroup,
	}
}

func NewInternalClient(ctx *context.Context) (InternalClient, error) {
	client := InternalClient{}

	if err := worker.Setup(groups()); err != nil {
		return client, err
	}
	return client, nil
}

func (c *InternalClient) Start() error {
	c.run = true
	c.operate()

	return nil
}

func (c *InternalClient) Stop() error {
	c.run = false
	return nil
}

func (c *InternalClient) operate() error {
	ch := make(chan interface{})
	freq := 1 * time.Second

	go func(channel chan interface{}) {
		for range time.Tick(freq) {
			channel <- struct{}{}
		}
	}(ch)

	for range ch {
		if !c.run {
			break
		}
		worker.Notify(cpuCount, int64(runtime.NumCPU()))
		worker.Notify(cpuCgoCall, int64(runtime.NumCgoCall()))
		worker.Notify(cpuGoroutines, int64(runtime.NumGoroutine()))

		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)
		worker.Notify(memAlloc, int64(stats.Alloc))
		worker.Notify(memFrees, int64(stats.Frees))
		worker.Notify(memGcCount, int64(stats.NumGC))
		worker.Notify(memGcLast, int64(stats.LastGC))
		worker.Notify(memGcNext, int64(stats.NextGC))
		// worker.Notify(memGcPause, int64(stats.PauseNs))
		worker.Notify(memGcPauseTotal, int64(stats.PauseTotalNs))

		worker.Notify(memLookups, int64(stats.Lookups))
		worker.Notify(memMalloc, int64(stats.Mallocs))
		worker.Notify(memOthersys, int64(stats.OtherSys))
		worker.Notify(memSys, int64(stats.Sys))
		worker.Notify(memTotalalloc, int64(stats.TotalAlloc))

		// heap
		worker.Notify(memHeapAlloc, int64(stats.HeapAlloc))
		worker.Notify(memHeapIdle, int64(stats.HeapIdle))
		worker.Notify(memHeapInuse, int64(stats.HeapInuse))
		worker.Notify(memHeapObjects, int64(stats.HeapObjects))
		worker.Notify(memHeapReleased, int64(stats.HeapReleased))
		worker.Notify(memHeapSys, int64(stats.HeapReleased))

		// stack
		worker.Notify(memStackInuse, int64(stats.StackInuse))
		worker.Notify(memStackSys, int64(stats.StackSys))
		worker.Notify(memStackMcacheInuse, int64(stats.MCacheInuse))
		worker.Notify(memStackMcacheSys, int64(stats.MCacheSys))
		worker.Notify(memStackMspanInuse, int64(stats.MSpanInuse))
		worker.Notify(memStackMspanSys, int64(stats.MSpanSys))
	}

	return nil
}

// InternalMonitor start the gobench client metrics collection should be run in
// a goroutine
func InternalMonitor() {
	ctx := context.Background()
	client, err := NewInternalClient(&ctx)
	if err != nil {
		log.Println(err)
		return
	}

	client.Start()
}
