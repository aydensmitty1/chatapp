package main

import (
	"github.houston.softwaregrp.net/CSB/chatapp/pkg/client"
	"flag"
)

func main(){
	port:=flag.String("port", "9992", "A port")
	address:=flag.String("ipad","10.12.182.133", "an Ip address")
	flag.Parse()
	cli:= client.Client{}
	cli.Call(address, port)
}