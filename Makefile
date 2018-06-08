

all: client server clean

client:
	go install github.houston.softwaregrp.net/CSB/chatapp/cmd/client

server:
	go install github.houston.softwaregrp.net/CSB/chatapp/cmd/server

clean:
	rm -rf ./buildtmp
