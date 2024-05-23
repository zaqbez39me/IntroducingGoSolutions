package main

import (
	"time"
)

func Sleep(duration time.Duration) {
	for {
		select {
		case <-time.After(duration):
			return
		}
	}
}

func main() {
	Sleep(time.Second)
}
