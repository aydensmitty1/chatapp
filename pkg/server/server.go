package server

import (
	"net"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"github.houston.softwaregrp.net/CSB/chatapp/pkg/message"
	"github.houston.softwaregrp.net/CSB/chatapp/pkg/clientregistration"
	)
type Server struct {
	conns []net.Conn
	conner []net.Conn
	messer []string
}

type Event struct {
	Type    string
	Payload string
}
var Msg struct{
	messer []string
	message.Message
}
func (s *Server) Listen(port *string) error {

	prt := *port
	fmt.Printf("prt: %s\n", prt)
	ln, err := net.Listen("tcp", ":"+prt)
	if err != nil {
		return err
	}
	fmt.Printf("server is listening on %s\n", prt)
	chatroomConns := make(map[string][]net.Conn)
	for {

		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("yo failed to accept connection. skipping")
			continue
		}
		//s.conns = append(s.conns, conn)
		var event Event

		decoder := json.NewDecoder(conn)

		rip := decoder.Decode(&event)
		if rip != nil {
			fmt.Println("rip dudes aaaahhh")
			fmt.Println(rip)
			continue
		}
		decode, err := base64.StdEncoding.DecodeString(event.Payload)
		dec := fmt.Sprintf("%s\n", decode)

		if err != nil {
			return err
		}

		if err != nil {
			fmt.Println("yo someone had issues connecting")
			continue
		}
		//go s.Handle(conn, chatroomConns, dec)
		switch event.Type{
		case "ClientRegistration":
			var cr clientregistration.Clientregistration
			err := json.Unmarshal([]byte(dec), &cr)
			if err != nil {
				// handle
			}
			//var chatroom string
			chatroomConns= clientregistration.Addclient(chatroomConns, cr,conn)


		fallthrough
		case "Msg":
			var mes clientregistration.Clientregistration
			err:= json.Unmarshal([]byte(dec),&mes)
			fmt.Println("mess address:", mes)
			if err !=nil{
				fmt.Println(err)
				return err
			}

			go s.Handle(conn, chatroomConns,mes)

		default:
			fmt.Println("sorry event not found")
		}
	}
}
func (s *Server) Handle(conn net.Conn, chatroomConns map[string][]net.Conn, mes clientregistration.Clientregistration) {
		fmt.Println("map:", chatroomConns)
		for {
			fmt.Println("handling the connection")
			jsonDecoder := json.NewDecoder(conn)
			err := jsonDecoder.Decode(&Msg)
			fmt.Println("client:", Msg)
			if err != nil {
				return
			}
			fmt.Println("mess:",mes)
			for _, messer := range mes.Chatrooms{
				fmt.Println(messer)
				for _, conn := range chatroomConns[messer] {
					_ = json.NewEncoder(conn).Encode(Msg)
						continue
					}
				}
			}

				fmt.Println("connection closed")
}



		//chatroomConns["bogus"] = append(chatroomConns["bogus"], bogusConn)
	// access a specific chatroom and send every client the message


// iterate over every single chat room and send the message to every client inside of each chatroom
//for chatroom, conns := range chatroomConns {
//for _, conn := range conns {
//err := json.NewEncoder(conn).Encode(Msg)
//if err != nil {
// do stuff
