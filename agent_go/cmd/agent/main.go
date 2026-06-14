package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/tekkpriest/cerberus/agent_go/internal/collector"
)

const (
	gb = 1024 * 1024 * 1024
)

func main() {
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	totalResources, err := collector.CollectSystemSpecs()
	if err != nil {
		fmt.Printf("error getting system information: %v", err)
	}

	fmt.Println("Cerberus Agent has been started... Press CTRL+C to quit.")

	for {
		select {
		case <-exitChan:
			fmt.Println("\nCerberus Agent has been stopped.")
			return
		case <-ticker.C:
			snapshot, err := collector.CollectSnapshots()
			if err != nil {
				fmt.Printf("Fehler beim Sammeln der Metriken: %v", err)
				continue
			}
			fmt.Println(strings.Repeat("-", 20))
			fmt.Printf("Current Time: %v\n", time.Now().Format("15:04:05"))

			fmt.Printf("CPU-Usage: %.2f%%\n", snapshot.SystemCPU)

			maxRAMGB := float64(totalResources.MaxRAM) / gb
			usedRAMGB := maxRAMGB * (snapshot.SystemRAM / 100)
			fmt.Printf("RAM-Usage: %.2f%% (%.2f GB of %.1f GB used)\n", snapshot.SystemRAM, usedRAMGB, maxRAMGB)

			maxDiskGB := float64(totalResources.MaxDiskSpace) / gb
			usedDiskGB := maxDiskGB * (snapshot.SystemDisk / 100)
			fmt.Printf("Disk-Usage: %.2f%% (%.2f GB of %.1f GB used)\n", snapshot.SystemDisk, usedDiskGB, maxDiskGB)
		}
	}
}
