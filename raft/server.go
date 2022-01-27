package raft

import (
	"fmt"
	"time"
)

func StartServer() {
	et := NewTimer()
	go et.WaitTillElection()

	// TODO: Use wait group maybe
	time.Sleep(10 * time.Second)
	et.Stop()
	fmt.Println("some timer impl")
}
