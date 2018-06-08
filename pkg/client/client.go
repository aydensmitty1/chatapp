package client

import (
	"net"
	"bufio"
		"fmt"
	"os"
)

type Client struct {
}

func (s *Client) Call(address *string,port *string) error {
	addr := fmt.Sprintf("%s:%s", *address, *port)
	fmt.Printf("dialing %s\n", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil{
		return err
	}
	go s.handleServerMessages(conn)
	for {
		reader:=bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		fmt.Fprintf(conn,text+"\n")
	}
}


func (s *Client) handleServerMessages(conn net.Conn) {
	for {
		message, err:= bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server:"+message)

		if err!=nil{
			break
		}
	}
}
