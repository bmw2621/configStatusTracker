package main

import (
	"encoding/json"
	"math/rand"
	"time"
)

var statuses = []*Status{
	{
		Service: "service1",
		Status: "initializing",
		Progress: 0,
	},
	{
		Service: "service2",
		Status: "phase 1",
		Progress: 10,
	},
	{
		Service: "service3",
		Status: "phase 2",
		Progress: 30,
	},
	{
		Service: "service4",
		Status: "phase 4",
		Progress: 50,
	},
	{
		Service: "service5",
		Status: "complete",
		Progress: 100,
	},
}

func init() {
	for _, status := range statuses {
		go simulate(status)
	}
}

// Status struct represents a service status
type Status struct {
	Service			string	`json:"service"`
	Status			string	`json:"status"`
	Progress		int		`json:"progress"`
}

// Broadcast method sends status to all clients in client slice
func (s *Status) Broadcast() {
	for _, client := range clients {
		s.BroadcastTo(client)
	}
}

// BroadcastTo method sends status to passed client
func (s *Status) BroadcastTo(to *Client) {
	byteMessage, _ := json.Marshal(s)
	to.ws.Write(byteMessage)
}

func simulate(s *Status) {
	time.Sleep(time.Millisecond * 500)

	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	gains := r1.Intn(5 - 1) + 1

	if s.Progress + gains <= 100{
		s.Progress = s.Progress + gains
	}

	switch s.Progress {
	case 10:
		s.Status = "phase 1"
	case 30:
		s.Status = "phase 2"
	case 50:
		s.Status = "phase 3"
	case 70:
		s.Status = "phase 4"
	case 90:
		s.Status = "finalizing"
	case 100:
		s.Status = "complete"
	}
	s.Broadcast()
	if s.Progress < 100 {
		go simulate(s)
	}
}