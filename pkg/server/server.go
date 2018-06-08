package server

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

type Server struct {
}


func (s *Server) Listen(port *string) error {
	prt:= *port
	fmt.Printf("prt: %s\n", prt)
	ln, err := net.Listen("tcp", ":" + prt)
	if err != nil {
		return err
	}
	fmt.Printf("server is listening on %s", prt)

	for {
		conn, err := ln.Accept()
		fmt.Println("we have accepted a new connection")
		if err != nil {
			fmt.Println("yo someone had issues connecting")
			continue
		}
		go s.handle(conn)
	}
}


func (s *Server) handle(conn net.Conn) {
	fmt.Println("handling the connection")
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Client:", message)

		reader:=bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		fmt.Fprintf(conn,text+"\n")
		if err != nil {
			break
		}

	}

	fmt.Println("connection closed")
}