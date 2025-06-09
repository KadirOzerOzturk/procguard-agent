package entities

type ProcessStats struct {
	Pid        int32   `json:"pid"`
	Name       string  `json:"name"`
	CPUPercent float64 `json:"cpuUsage"`
	MemoryMB   float64 `json:"memoryUsage"`
}
