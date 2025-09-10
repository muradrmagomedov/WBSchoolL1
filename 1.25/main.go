package main

import (
	"fmt"
	"time"
)

func main() {
	sleep(3 * time.Second)
	fmt.Println("Wake up")
}

func sleep(t time.Duration) {
	<-time.After(t)
}
