package entities

type NetworkStats struct {
	SentMB     float64 `json:"sentMB"`
	ReceivedMB float64 `json:"receivedMB"`
}
