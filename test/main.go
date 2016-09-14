package gocket;

import (
	"log"
	"golang.org/x/net/websocket"
	"github.com/chuckpreslar/emission"
)

var emitter = emission.NewEmitter()

type WsTyper struct {
	MessageType string 		 `json:"type"`
	Data 		interface {} `json:"data"`
}

func webSocketServer(ws *websocket.Conn) {
	log.Println("Connection Started!")

	var response WsTyper;
	send := func (messageType string, data interface {}) {
		websocket.JSON.Send(ws, WsTyper{
			MessageType: messageType,
			Data: data,
		})
	}

	for {
		err := websocket.JSON.Receive(ws, &response)
	
		if err != nil {
			log.Println("Connection Closed!")
			ws.Close()
			break
		}

		emitter.Emit(response.MessageType, response.Data, send)
	}


}

func On (messageType string, callback func(data interface{}, callback func(m string, d interface{}))) {
	emitter.On(messageType, callback)
}

func SocketHandler() websocket.Handler {
	return websocket.Handler(webSocketServer)
}