package main

import (
	"log"
	"net"
	"os"
)

func main() {

	s := newServer()
	// start the server in a separate thread
	go s.run()

	listener,err := net.Listen("tcp",":8888")

	if err!=nil {
		log.Fatalf("unable to start server : %s", err.Error())
	}
	log.Println("Os args " , len(os.Args))	

	defer listener.Close()
	log.Printf("started server on :8888")

	for {

		conn ,err := listener.Accept()
		
		log.Printf("Connection details localserver: %s , ", conn.LocalAddr())

		if err != nil {
			log.Printf("Unable to accept connections : %s",err.Error())
			continue
		}
	
		c := s.newClient(conn);
		go c.readInput()
		
	}
}