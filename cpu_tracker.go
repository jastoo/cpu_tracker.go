package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/process"
)

func main() {
	r := gin.Default()
	r.StaticFile("/", "./index.html")

	r.GET("/cpu", func(c *gin.Context) {
		percentages, err := cpu.Percent(0, false)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cpu": percentages[0]})
	})

	r.GET("/processes", func(c *gin.Context) {
		processes, err := process.Processes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var procs []gin.H
		for _, proc := range processes {
			name, err := proc.Name()
			if err != nil {
				name = "unknown"
			}
			pid := proc.Pid
			cpuPercent, err := proc.CPUPercent()
			if err != nil {
				cpuPercent = 0.0
			}
			procs = append(procs, gin.H{
				"pid":  pid,
				"name": name,
				"cpu":  cpuPercent,
			})
		}
		c.JSON(http.StatusOK, procs)
	})

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port

	go func() {
		if err := http.Serve(listener, r); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Server listening on port %d\n", port)
}
