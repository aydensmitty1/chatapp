package main

import (
	"github.houston.softwaregrp.net/CSB/chatapp/pkg/server"
	"flag"
)

func main() {
	port:=flag.String("port", "9992", "A port for hosting server")
	flag.Parse()
	srv := server.Server{}
	srv.Listen(port)

}
