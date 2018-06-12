package client

import (
	"net"
	"bufio"
		"fmt"
	"os"
	"github.houston.softwaregrp.net/CSB/chatapp/pkg/message"
	"encoding/json"
	"time"
)

type Client struct {
}
func (s *Client) Call(address *string,port *string, username *string) {
	addr := fmt.Sprintf("%s:%s", *address, *port)
	fmt.Printf("dialing %s\n", addr)
	conn, _ := net.Dial("tcp", addr)

	go s.handleServerMessages(conn, username)
	for {
		t:= time.Now()
		x:=fmt.Sprintf(t.Format(time.UnixDate))
		tx:=fmt.Sprintf("%s:", x)

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		username := fmt.Sprint(*username)

		msg := &message.Message{username, tx, text}
		S, err := json.Marshal(msg)

		message.Decode(S)
		fmt.Sprint(*msg)
		fmt.Fprint(conn,msg)
		if err != nil {
			continue
		}
	}
}
func (s *Client) handleServerMessages(conn net.Conn, username *string) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(message)
		if err != nil {
			return
		}
	}

}

