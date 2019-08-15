---
title: Quick Monitor and Restart in Golang
author: Luis Rodriguez
type: post
date: 2019-08-14
url: /2019/08/14/quick-monitor-restart-golang
categories:
  - Projects
  - Golang
tags:
  - go
  - tutorial
  - monitor
  - service
  - golang

---

## Go Applications

On my desktop I was running a processing script that would occasionally hang at high CPU, so I wrote a quick script to monitor it and restart it. You can see it below and change it as needed.


Run it like so after compiling to a binary

```bash
nohup ./restart &
```

The code

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

func main() {

	// Check every 30 seconds
	for {
		time.Sleep(30 * time.Second)

		idle0, total0 := getCPUSample()
		time.Sleep(3 * time.Second)
		idle1, total1 := getCPUSample()

		idleTicks := float64(idle1 - idle0)
		totalTicks := float64(total1 - total0)
		cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

		if cpuUsage > 50 {

			cmd := exec.Command("sudo", "pm2", "restart", "all")
			if runtime.GOOS == "windows" {
				cmd = exec.Command("tasklist")
			}
			err := cmd.Run()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}

			fmt.Printf("CPU usage is %f%%\n", cpuUsage)
		}
	}

}
```