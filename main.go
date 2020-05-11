package main

import (
	"time"
	. "yogurt/nats"
)

const (
	subj = "weather"
)

func main() {

	StartServer(subj, "s1")
	StartServer(subj, "s2")
	StartServer(subj, "s3")
	//wait for subscribe complete
	time.Sleep(1 * time.Second)

	StartClient(subj)

	select {}
}
