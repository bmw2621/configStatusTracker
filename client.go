package main

import (
	"math/rand"
	"time"

	"golang.org/x/net/websocket"
)


var clients []*Client

type Client struct {
	id 		string
	ws 		*websocket.Conn
}

func (client *Client) StartListening() {
	buffer := make([]byte, 512)

	for {
		n, err := client.ws.Read(buffer[0:])

		if err != nil {
			ReleaseConnection(client)
		} else {
		HandleInputMessage(client, buffer[0:n])
		}
	}
}

func HandleNewConnection(c *websocket.Conn) {
	newClient := Client{
		id: genId(),
		ws: c,
	}

	clients = append(clients, &newClient)

	newClient.StartListening()
}

func ReleaseConnection(client *Client){
	index := -1
	for idx, val := range clients {
		if client == val {
			index = idx
			break
		}
	}

	if index >= 0 {
		clients = append(clients[:index], clients[index + 1:]...)
	}

	client.ws.Close()
}

func genId() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}