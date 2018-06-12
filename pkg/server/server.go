package server

import (
	"net"
	"fmt"
	"bufio"
	)

type Server struct {
	conns []net.Conn
}


func (s *Server) Listen(port *string) error {
	prt:= *port
	fmt.Printf("prt: %s\n", prt)
	ln, err := net.Listen("tcp", ":" + prt)
	if err != nil {
		return err
	}
	fmt.Printf("server is listening on %s\n", prt)

	for {
		conn, err := ln.Accept()
		fmt.Printf("client connected: %s\n", conn.RemoteAddr().String())
		if err != nil {
			fmt.Println("yo someone had issues connecting")
			continue
		}
		s.conns = append(s.conns, conn)
		fmt.Printf("number of clients connected: %d\n", len(s.conns))
		go s.handle(conn)
	}
}

// loop through the s.conns slice of connections
// and send the message to each of them.
// if an error occurs when you are parroting the message to one of the clients
// make sure it doesn't prevent the rest of the clients from getting the message. (handle the error here, don't propagate it)
func (s *Server) handle(conn net.Conn) {
	fmt.Println("handling the connection")
	for {

		message, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Client:", message)
		if err !=nil{
			return
		}

		// for each connection in s.conns
		// send the message to them
		for i := 0; i < len(s.conns); i++ {
			peerConn := s.conns[i]

			fmt.Fprintf(peerConn, message,s.conns)
			if err != nil {
				//fmt.Println("there was an error forwarding message to peer")
				continue
			}
		}
	}

	fmt.Println("connection closed")
}