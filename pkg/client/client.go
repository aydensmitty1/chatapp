package client

import (
	"net"
	"bufio"
	"fmt"
	"os"
	"time"
	"encoding/json"
	"github.houston.softwaregrp.net/CSB/chatapp/pkg/message"
	"github.houston.softwaregrp.net/CSB/chatapp/pkg/clientregistration"
)

type Client struct {
}
var messer struct {
	message.Message
}
func (s *Client) Call(address *string, port *string, username *string, chatrm *string) {
	addr := fmt.Sprintf("%s:%s", *address, *port)
	fmt.Printf("dialing %s\n", addr)
	conn, _ := net.Dial("tcp", addr)
	clientregistration.Clientreg(*chatrm, conn)
	go s.handleServerMessages(conn)
	for {
		t := time.Now()
		x := fmt.Sprintf(t.Format(time.UnixDate))
		tx := fmt.Sprintf("%s:", x)

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		username := fmt.Sprint(*username)
		chatroom := fmt.Sprintf(*chatrm)
		msg := &message.Message {Author: username, Chatroom: chatroom, Time: tx, Text: text}

		jsonEncoder := json.NewEncoder(conn)
		err := jsonEncoder.Encode(msg)
		if err != nil {
			fmt.Println("error marshalling message")
			fmt.Println(err)
		}
	}
}
func (s *Client) handleServerMessages(conn net.Conn) {
	for {
		jsonDecoder := json.NewDecoder(conn)
		err := jsonDecoder.Decode(&messer)
		fmt.Println(messer)
		if err !=nil{
			return
		}

	}

}
