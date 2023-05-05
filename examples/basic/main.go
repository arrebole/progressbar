package main

import (
	"time"

	"github.com/arrebole/progressbar"
)

func main() {
	bar := progressbar.Default(10000)
	bar.SetOffset(100)
	for i := 0; i < 1000; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
}
