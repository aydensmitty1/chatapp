

all: client server

client:
	mkdir -p ./build/.buildtemp
	go install github.houston.softwaregrp.net/CSB/chatapp/cmd/client
	go build -o ./build/.buildtemp/client github.houston.softwaregrp.net/CSB/chatapp/cmd/client
	GOOS=darwin GOARCH=amd64 go build -o ./build/.buildtemp/client-darwin-amd64 github.houston.softwaregrp.net/CSB/chatapp/cmd/client

server:
	go install github.houston.softwaregrp.net/CSB/chatapp/cmd/server

clean:
	rm -rf ./build/.buildtemp
