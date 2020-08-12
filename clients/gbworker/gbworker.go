package gbworker

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/gobench-io/gobench/executor/driver"
	"github.com/gobench-io/gobench/metrics"
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
			{
				Title: "CPU Cores",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: cpuCount,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "CGO Calls",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: cpuCgoCall,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Goroutines",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
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
			{
				Title: "Allocated Heap",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memAlloc,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Cumulative Freed Heap",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memFrees,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Completed GC Cycles",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: memGcCount,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Last GC Finished",
				Unit:  "ns (Unix epoch)",
				Metrics: []metrics.Metric{
					{
						Title: memGcLast,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Target Next GC Cycle",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
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
			{
				Title: "Total GC Pause",
				Unit:  "ns",
				Metrics: []metrics.Metric{
					{
						Title: memGcPauseTotal,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "GC Metadata",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memGcSys,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Pointer Lookups",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: memLookups,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Cumulative Heap Objects Allocated",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: memMalloc,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Miscellaneous off-heap Runtime Allocation",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: memOthersys,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Total Mem from the OS",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memSys,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Cumulative Allocated for Heap Objects",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
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
			{
				Title: "Heap Allocated Objects",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memHeapAlloc,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Heap Idle (unused) Spans",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memHeapIdle,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Heap In-use Spans",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memHeapInuse,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Heap Allocated Object",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memHeapObjects,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Physical Memory Returned to OS",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memHeapReleased,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Heap Memory Obtained from the OS",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
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
			{
				Title: "In-use Stack Spans",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memStackInuse,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Stack Memory Obtained from the OS",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memStackSys,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Allocated mcache Structures",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memStackMcacheInuse,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Memory Obtained from the OS for mcache Structures",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memStackMcacheSys,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Allocated mspan Structures",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
						Title: memStackMspanInuse,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "Memory Obtained from the OS for mspan Structures",
				Unit:  "Byte",
				Metrics: []metrics.Metric{
					{
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

	if err := driver.Setup(groups()); err != nil {
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
		driver.Notify(cpuCount, int64(runtime.NumCPU()))
		driver.Notify(cpuCgoCall, int64(runtime.NumCgoCall()))
		driver.Notify(cpuGoroutines, int64(runtime.NumGoroutine()))

		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)
		driver.Notify(memAlloc, int64(stats.Alloc))
		driver.Notify(memFrees, int64(stats.Frees))
		driver.Notify(memGcCount, int64(stats.NumGC))
		driver.Notify(memGcLast, int64(stats.LastGC))
		driver.Notify(memGcNext, int64(stats.NextGC))
		// driver.Notify(memGcPause, int64(stats.PauseNs))
		driver.Notify(memGcPauseTotal, int64(stats.PauseTotalNs))

		driver.Notify(memLookups, int64(stats.Lookups))
		driver.Notify(memMalloc, int64(stats.Mallocs))
		driver.Notify(memOthersys, int64(stats.OtherSys))
		driver.Notify(memSys, int64(stats.Sys))
		driver.Notify(memTotalalloc, int64(stats.TotalAlloc))

		// heap
		driver.Notify(memHeapAlloc, int64(stats.HeapAlloc))
		driver.Notify(memHeapIdle, int64(stats.HeapIdle))
		driver.Notify(memHeapInuse, int64(stats.HeapInuse))
		driver.Notify(memHeapObjects, int64(stats.HeapObjects))
		driver.Notify(memHeapReleased, int64(stats.HeapReleased))
		driver.Notify(memHeapSys, int64(stats.HeapReleased))

		// stack
		driver.Notify(memStackInuse, int64(stats.StackInuse))
		driver.Notify(memStackSys, int64(stats.StackSys))
		driver.Notify(memStackMcacheInuse, int64(stats.MCacheInuse))
		driver.Notify(memStackMcacheSys, int64(stats.MCacheSys))
		driver.Notify(memStackMspanInuse, int64(stats.MSpanInuse))
		driver.Notify(memStackMspanSys, int64(stats.MSpanSys))
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
