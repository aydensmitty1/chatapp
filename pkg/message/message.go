package message

import (
		"encoding/json"
	"fmt"
	)
type Message struct {
	Author string
	Chatroom string
	Time string
	Text string
}

func Encode(username string, tx string  , text string, chatroom string){
	msg := &Message{username, tx,chatroom,text}
	S, err := json.Marshal(msg)
	j:=json.Unmarshal(S, &msg)
	fmt.Println(msg)
	fmt.Println(j)
	if err!=nil{
		return
	}

}
func Decode(S []byte){
	var msg string
	err:=json.Unmarshal(S, &msg)
	fmt.Println(msg)
	if err != nil{
		return
	}
}



