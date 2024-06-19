package main

import (
	"net"
)

type room struct {
	name string
	members map[net.Conn]*client
}

// broadcasting messages to all memebers of a room
func (r *room) brodcastMessage(message string) {
	for _,client := range r.members {
		client.msg(message)
	}
}

