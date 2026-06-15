package collector

import "time"

type Snapshot struct {
	ServerID   string    `json:"-"`
	SystemCPU  float64   `json:"cpu_usage"`
	SystemRAM  float64   `json:"ram_usage"`
	SystemDisk float64   `json:"disk_usage"`
	Timestamp  time.Time `json:"timestamp"`
}

type SystemSpecs struct {
	ServerID     string `json:"-"`
	MaxRAM       uint64 `json:"max_ram"`
	MaxDiskSpace uint64 `json:"max_disk_space"`
}
