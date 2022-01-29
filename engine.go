package main

import (
	"encoding/json"
)

// HandleInputMessage handles a new client connection and sends all current status
func HandleInputMessage(from *Client, data []byte) {
	var input map[string]string
	json.Unmarshal(data, &input)

	
	switch input["action"]{
		case "init":
			for _, status := range statuses {
				status.BroadcastTo(from)
			}
	}
}