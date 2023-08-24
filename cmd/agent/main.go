package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type Monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs,
	BuckHashSys,

	GCSys,
	HeapAlloc,
	HeapIdle,
	HeapInuse,
	HeapObjects,
	HeapReleased,
	HeapSys,
	LastGC,
	Lookups,
	MCacheInuse,
	MCacheSys,
	MSpanInuse,
	MSpanSys,
	NextGC,

	OtherSys,

	StackInuse,
	StackSys uint64
	GCCPUFraction float64

	NumForcedGC, NumGC uint32
	NumGoroutine       int
}

func main() {
	var m Monitor
	var rtm runtime.MemStats
	var interval = time.Duration(1) * time.Second
	for {
		<-time.After(interval)

		// Read full mem stats
		runtime.ReadMemStats(&rtm)

		// Number of goroutines
		m.NumGoroutine = runtime.NumGoroutine()

		// Misc memory stats
		m.Alloc = rtm.Alloc
		m.BuckHashSys = rtm.BuckHashSys
		m.GCCPUFraction = rtm.GCCPUFraction
		m.TotalAlloc = rtm.TotalAlloc
		m.Sys = rtm.Sys
		m.Mallocs = rtm.Mallocs
		m.Frees = rtm.Frees
		m.GCSys = rtm.GCSys
		m.HeapAlloc = rtm.HeapAlloc
		m.HeapIdle = rtm.HeapIdle
		m.HeapInuse = rtm.HeapInuse
		m.HeapObjects = rtm.HeapObjects
		m.HeapReleased = rtm.HeapReleased
		m.HeapSys = rtm.HeapSys
		m.Lookups = rtm.Lookups
		m.MCacheInuse = rtm.MCacheInuse
		m.MCacheSys = rtm.MCacheSys
		m.MSpanInuse = rtm.MSpanInuse
		m.MSpanSys = rtm.MSpanSys
		m.NextGC = rtm.NextGC
		m.NumForcedGC = rtm.NumForcedGC
		m.OtherSys = rtm.OtherSys
		m.PauseTotalNs = rtm.PauseTotalNs
		m.StackInuse = rtm.StackInuse
		m.StackSys = rtm.StackSys
		m.NumGC = rtm.NumGC

		// Just encode to json and print
		b, _ := json.Marshal(m)
		fmt.Println(string(b))
	}
}
