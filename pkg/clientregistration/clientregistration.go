package clientregistration

import (
	"encoding/json"
	"fmt"
	"net"
	"encoding/base64"
	"strings"
)

type Clientregistration struct{
	Chatrooms []string
}

type Event struct{
	Type string
	Payload string
}

func Clientreg(chatroom string, conn net.Conn) error {
	chatrooms:=strings.Split(chatroom,",")
	clientregistration := &Clientregistration{Chatrooms: chatrooms}
	crPayload, err := json.Marshal(clientregistration)
	if err != nil {
		return err
	}
	fmt.Println(string(crPayload))
	event := &Event{Type: "ClientRegistration", Payload: base64.StdEncoding.EncodeToString(crPayload)}
	return json.NewEncoder(conn).Encode(event)
}



func Readclientreg(eventPayload []byte,conn net.Conn){
	var event Event
	err:=json.Unmarshal(eventPayload, &event)
	if err !=nil {
	return
	}
}

func Addclient(chatroomConns map[string][]net.Conn, cr Clientregistration,conn net.Conn) (map[string][]net.Conn,) {
	for _, chatroom := range cr.Chatrooms {
		chatroomConns[chatroom] = append(chatroomConns[chatroom], conn)

		fmt.Println("map before:", chatroomConns)
	}
		return chatroomConns

}
