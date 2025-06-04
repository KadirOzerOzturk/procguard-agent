package entities

type SystemStats struct {
	AgentID      string         `json:"agentId"`
	Timestamp    string         `json:"timestamp"`
	CPUUsage     float64        `json:"cpuUsage"`
	MemoryUsage  float64        `json:"memoryUsage"`
	Network      NetworkStats   `json:"network"`
	TopProcesses []ProcessStats `json:"topProcesses"`
	Disk         DiskStats      `json:"disk"`
}
