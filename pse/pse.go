package pse

import (
	"sync/atomic"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

var (
	cpuPrevTime time.Time
	cpuPrev     *cpu.Stats
	ipcpu       int64
)

func init() {
	periodic()
}

func periodic() {
	now := time.Now()
	cpuNow, err := cpu.Get()
	if err != nil {
		return
	}

	if cpuPrev != nil {
		seconds := now.Sub(cpuPrevTime).Seconds()

		atomic.StoreInt64(&ipcpu,
			int64((float64)(cpuNow.User+cpuNow.System-cpuPrev.User-cpuPrev.System)*10.0/seconds))
	}

	cpuPrevTime = now
	cpuPrev = cpuNow

	time.AfterFunc(1*time.Second, periodic)
}

// ProcUsage gets CPU usage and user memory usage
func ProcUsage(pcpu *float64, mem *uint64) error {
	*pcpu = float64(atomic.LoadInt64(&ipcpu)) / 10.0

	m, err := memory.Get()
	if err != nil {
		return err
	}
	*mem = m.Used

	return nil
}
