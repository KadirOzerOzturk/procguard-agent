package services

import (
	"net"
	"os"
	"runtime"
	"time"
)

func GetAgentInfo() (map[string]interface{}, error) {
	agentInfo := make(map[string]interface{})

	// OS bilgisi
	agentInfo["os"] = runtime.GOOS

	// Architecture bilgisi
	agentInfo["architecture"] = runtime.GOARCH

	// Hostname (name olarak kullanılabilir)
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	agentInfo["name"] = hostname

	ipAddress := ""
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ipAddress = ipnet.IP.String()
				break
			}
		}
	}
	agentInfo["ipAddress"] = ipAddress

	agentInfo["status"] = "active"

	agentInfo["version"] = "1.0.0"

	now := time.Now()

	agentInfo["lastSeenAt"] = now.Format("2006-01-02T15:04:05")

	return agentInfo, nil
}
