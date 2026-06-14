package collector

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func CollectSnapshots() (*Snapshot, error) {
	systemCPU, err := cpu.Percent(0, false)
	if err != nil {
		return nil, fmt.Errorf("error reading CPU values: %w", err)
	}

	systemRAM, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error reading RAM values: %w", err)
	}

	systemDisk, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("error reading Disk values: %w", err)
	}

	return &Snapshot{
		ServerID:   "-",
		SystemCPU:  systemCPU[0],
		SystemRAM:  systemRAM.UsedPercent,
		SystemDisk: systemDisk.UsedPercent,
		Timestamp:  time.Now(),
	}, nil
}

func CollectSystemSpecs() (*SystemSpecs, error) {
	systemRAM, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error reading RAM values: %w", err)
	}

	systemDisk, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("error reading Disk values: %w", err)
	}

	return &SystemSpecs{
		ServerID:     "-",
		MaxRAM:       systemRAM.Total,
		MaxDiskSpace: systemDisk.Total,
	}, nil
}
