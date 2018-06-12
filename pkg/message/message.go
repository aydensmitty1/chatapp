package message

import (
	"time"
	"encoding/json"
	"fmt"
)
type Message struct {
	Author string
	Timestamp string
	Text string
}

func Encode(username string, x time.Time  , text string) {
	msg := &Message{username, "time", text}
	S, err := json.Marshal(msg)
	fmt.Println(string(S))
	if err != nil {
		fmt.Print("their was an error")
	}
}
func Decode(S []byte){
	var msg Message
	err:=json.Unmarshal(S, &msg)
	//fmt.Println(msg)
	//fmt.Print(err)
	if err != nil{
	}
}



