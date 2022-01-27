package raft

import (
	"fmt"
	"time"
)

type Timer interface {
	WaitTillElection()
	Stop()
}

func NewTimer() Timer {
	return &electionTimer{
		ticker: time.NewTicker(getRandomTimeout()),
		done:   make(chan bool),
	}
}

type electionTimer struct {
	ticker *time.Ticker
	done   chan bool
}

func (et *electionTimer) WaitTillElection() {
	for {
		select {
		case <-et.done:
			return
		case t := <-et.ticker.C:
			fmt.Println("Ticked at ", t)
			// TODO: Invoke the TCP Client to send heartbeat to all other servers
		}
	}
}

func (et *electionTimer) Stop() {
	et.done <- true
}

// TODO: Make this randomized
func getRandomTimeout() time.Duration {
	return time.Second * 3
}
