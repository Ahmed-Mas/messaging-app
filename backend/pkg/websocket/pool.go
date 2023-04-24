package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
	Clients    map[*Client]bool
}

func NewPool() *Pool {
	pool := &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
		Clients:    make(map[*Client]bool),
	}
	return pool
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(p.Clients))
			for client := range p.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
		case client := <-p.Unregister:
			delete(p.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(p.Clients))
			for client := range p.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
		case msg := <-p.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client := range p.Clients {
				fmt.Println(client)
				if err := client.Conn.WriteJSON(msg); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
