package config

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

type WebsocketConfig struct {
	*socketio.Server
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func (s *WebsocketConfig) Setup() {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "join", func(s socketio.Conn, msg map[string]string) {
		log.Println("join " + msg["roomName"])
		s.Join(msg["roomName"])
	})

	server.OnEvent("/", "unjoin", func(s socketio.Conn, msg map[string]string) {
		log.Println("unjoin " + msg["roomName"])
		s.Leave(msg["roomName"])
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg map[string]string) {
		s.Emit("notice_1", "have "+msg["msg"])
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	s.Server = server

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Websocket Serving at localhost:15001...")
	log.Fatal(http.ListenAndServe(":15001", nil))

}
