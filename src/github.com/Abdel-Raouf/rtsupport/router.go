main package

import (
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type router struct {}

func (e *router) serveHTTP(w http.ResponseWriter, r *http.Request){
		socket, err := upgrader.Upgrade(w, r, nil)
}
