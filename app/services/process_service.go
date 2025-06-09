package services

import (
	"time"

	"github.com/KadirOzerOzturk/procguard-agent/app/entities"
	"github.com/shirou/gopsutil/v3/process"
)

func KillProcess(pid int32) error {
	p, err := process.NewProcess(pid)
	if err != nil {
		return err
	}
	return p.Kill()
}
func GetTopProcesses(limit int) ([]entities.ProcessStats, error) {
	procs, _ := process.Processes()
	var result []entities.ProcessStats

	for _, p := range procs {
		cpuPercent, err := p.CPUPercent()
		if err != nil || cpuPercent < 0.5 {
			continue
		}
		name, _ := p.Name()
		result = append(result, entities.ProcessStats{
			Pid:        p.Pid,
			Name:       name,
			CPUPercent: cpuPercent,
		})
		if len(result) >= limit {
			break
		}
	}
	return result, nil
}

func GetAllProcesses() ([]entities.ProcessStats, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	time.Sleep(500 * time.Millisecond)

	var result []entities.ProcessStats
	for _, p := range procs {
		name, err := p.Name()
		if err != nil || name == "" {
			name = "Unknown"
		}

		cpuPercent, err := p.CPUPercent()
		if err != nil {
			cpuPercent = 0.0
		}

		if cpuPercent == 0.0 {
			continue
		}

		memInfo, err := p.MemoryInfo()
		var memoryMB float64
		if err == nil && memInfo != nil {
			memoryMB = float64(memInfo.RSS) / 1024.0 / 1024.0
		} else {
			memoryMB = 0.0
		}

		result = append(result, entities.ProcessStats{
			Pid:        p.Pid,
			Name:       name,
			CPUPercent: cpuPercent,
			MemoryMB:   memoryMB,
		})
	}

	return result, nil
}
