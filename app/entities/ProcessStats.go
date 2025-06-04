package entities

type ProcessStats struct {
	Pid        int32   `json:"pid"`
	Name       string  `json:"name"`
	CPUPercent float64 `json:"cpu_percent"`
}
