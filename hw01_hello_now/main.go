package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	currentTime := time.Now()
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	fmt.Printf("current time: %s\n", currentTime)
	fmt.Printf("exact time: %s\n", exactTime)
}
