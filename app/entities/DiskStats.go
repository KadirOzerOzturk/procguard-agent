package entities

type DiskStats struct {
	TotalGB     float64 `json:"totalGB"`
	UsedGB      float64 `json:"usedGB"`
	UsedPercent float64 `json:"usedPercent"`
}
