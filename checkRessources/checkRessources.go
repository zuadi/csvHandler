package checkressources

import (
	"fmt"
	"runtime"
)

func PrintMemoryUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	if memStats.Alloc > 1000000000 {
		fmt.Printf("Total allocated memory %dGB\n", memStats.Alloc/1000000000)
	} else if memStats.Alloc > 1000000 {
		fmt.Printf("Total allocated memory %dMB\n", memStats.Alloc/1000000)
	} else if memStats.Alloc > 1000 {
		fmt.Printf("Total allocated memory %dkB\n", memStats.Alloc/1000)
	} else {
		fmt.Printf("Total allocated memory %dBytes\n", memStats.Alloc)
	}
	if memStats.HeapAlloc > 1000000000 {
		fmt.Printf("Heap memory %dGB\n", memStats.HeapAlloc/1000000000)
	} else if memStats.HeapAlloc > 1000000 {
		fmt.Printf("Heap memory %dMB\n", memStats.HeapAlloc/1000000)
	} else if memStats.HeapAlloc > 1000 {
		fmt.Printf("Heap memory %dkB\n", memStats.HeapAlloc/1000)
	} else {
		fmt.Printf("Heap memory %dBytes\n", memStats.HeapAlloc)
	}
	fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
}
