package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/KadirOzerOzturk/procguard-agent/app/entities"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	gopsnet "github.com/shirou/gopsutil/v3/net"
)

// GetAgentID returns a unique agent ID based on local IP
func GetAgentID() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "host-unknown"
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return "host-" + ipNet.IP.String()
		}
	}
	return "host-unknown"
}

// CollectSystemStats gathers all required system statistics
func CollectSystemStats() (entities.SystemStats, error) {
	cpu.Percent(0, false)
	time.Sleep(500 * time.Millisecond)

	// Gerçek ölçüm
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return entities.SystemStats{}, err
	}
	memStats, _ := mem.VirtualMemory()
	netStats, _ := gopsnet.IOCounters(false)
	diskStats, _ := disk.Usage(getDiskMountPoint())
	topProcs, _ := GetTopProcesses(5)

	return entities.SystemStats{
		AgentID:     GetAgentID(),
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
		CPUUsage:    safeFloat(cpuPercent),
		MemoryUsage: memStats.UsedPercent,
		Network: entities.NetworkStats{
			SentMB:     bytesToMB(netStats),
			ReceivedMB: bytesReceivedToMB(netStats),
		},
		TopProcesses: topProcs,
		Disk: entities.DiskStats{
			TotalGB:     toGB(diskStats.Total),
			UsedGB:      toGB(diskStats.Used),
			UsedPercent: diskStats.UsedPercent,
		},
	}, nil
}

// SendStatsToAPI sends collected stats to the backend
func SendStatsToAPI(stats entities.SystemStats) error {
	data, err := json.Marshal(stats)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		return fmt.Errorf("API_URL environment variable is not set")
	}

	req, err := http.NewRequest("POST", apiURL+"/stats", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("request creation failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	fmt.Println("✅ Stats successfully sent to backend.")
	return nil
}

// PrintStatsAsJSON pretty prints the stats
func PrintStatsAsJSON(stats entities.SystemStats) {
	if data, err := json.MarshalIndent(stats, "", "  "); err == nil {
		fmt.Println(string(data))
	}
}

func getDiskMountPoint() string {
	parts, _ := disk.Partitions(false)
	for _, part := range parts {
		if strings.HasPrefix(part.Device, "\\Device\\Harddisk") || part.Mountpoint == "/" || strings.HasPrefix(part.Mountpoint, "C:") {
			return part.Mountpoint
		}
	}
	return "/"
}

func toGB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}

func safeFloat(arr []float64) float64 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}

func bytesToMB(stats []gopsnet.IOCountersStat) float64 {
	if len(stats) > 0 {
		return float64(stats[0].BytesSent) / (1024 * 1024)
	}
	return 0
}

func bytesReceivedToMB(stats []gopsnet.IOCountersStat) float64 {
	if len(stats) > 0 {
		return float64(stats[0].BytesRecv) / (1024 * 1024)
	}
	return 0
}
