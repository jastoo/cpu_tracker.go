package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func main() {
	for {
		percentages, err := cpu.Percent(time.Second, false)
		if err != nil {
			fmt.Println("Error getting CPU usage:", err)
			return
		}
		fmt.Println("CPU Usage: %.2f%%\n", percentages[0])
		time.Sleep(time.Second)
	}
}
