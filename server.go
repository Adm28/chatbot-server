package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms map[string]*room
	commands chan command
}

// Constructor of the server
func newServer() *server {
	return &server {
		rooms : make(map[string]*room),
		commands : make(chan command),
	}
}


func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)	
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args) 
		default:
			
		 }
	}
}

func (s *server) nick(c* client, args []string) {
	// Create an user with name, assign the name, encounter a nick command
	if len(args) < 2 {
		c.msg("nick is required. usage: /nick NAME")
		return
	}

	c.nick = args[1]
	c.msg(fmt.Sprintf("You have been assigned with name : %s", args[1]))

}

func (s *server) join(c *client, args[]string) {
	if len(args) < 2 {
		c.msg("room name is required. usage: /join ROOM_NAME")
		return
	}
	// check if an existing room exists , add the client in the room client map
	// else create a function to make a new room
	existingRoom, doesExist := s.rooms[args[1]]
	if !doesExist {
		existingRoom = &room {
			name : args[1],
			members : make(map[net.Conn]*client),
		}
		s.rooms[args[0]] = existingRoom
	}
	existingRoom.members[c.conn] = c
	c.room = existingRoom
	s.rooms[existingRoom.name] = existingRoom
	existingRoom.brodcastMessage(fmt.Sprintf("%s has been added to the room %s", c.nick, existingRoom.name))
	
}

func (s *server) listRooms(c *client, args[]string) {
	// list all the rooms.
	roomList := make([]string, 0, len(s.rooms))
	for k := range s.rooms {
		roomList = append(roomList, k)
	}
	roomString := strings.Join(roomList, " ,")
	c.msg(fmt.Sprintf("Current List of Rooms available : %s", roomString))

}


func (s *server) msg(c *client, args[]string) {
	if len(args) < 2 {
		c.msg("message is required, usage: /msg MSG")
		return
	}
	// brodcast the message to all clients in the same room
	c.room.brodcastMessage(args[1])
	
}

func (s *server) quit(c *client, args[]string) {
	// remove the userfrom the map and terminate the connection

}


func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		oldRoom := s.rooms[c.room.name]
		delete(s.rooms[c.room.name].members, c.conn)
		oldRoom.brodcastMessage(fmt.Sprintf("%s has left the room", c.nick))
	}
}

func (s *server) newClient(conn net.Conn) *client {

	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	// Assigning the new client witht the server command channel;.
	c := &client{
		conn: conn,
		nick: "anonymous",
		commands: s.commands,
	}

	return c

}