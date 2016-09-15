package gocket;

import (
	"log"
	"github.com/pborman/uuid"
	"golang.org/x/net/websocket"
)

type Conn struct {
	id 		string
	rooms 	[]string
	ws 		websocket.Conn
}

type WebsocketResponse struct {
	MessageType string 		`json:"type"`
	Data 		interface{} `json:"data"`
}

type Gocket struct {
	listeners 	map[string]func(c Conn, data interface{})
	connections []Conn
}

var gocket = Gocket{
	listeners: map[string]func(c Conn, data interface{}){},
	connections: []Conn{},
}

func Handler() websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		log.Println("Connection Started!")

		var conn = Conn{
			id: uuid.New(),
			rooms: []string{"self"},
			ws: *ws,
		}

		gocket.listeners["Connect"](conn, nil)
		gocket.connections = append(gocket.connections, conn)
		
		messageReceive(conn)
	})
}

func messageReceive(conn Conn) {
	var response WebsocketResponse;

	for {
		err := websocket.JSON.Receive(&conn.ws, &response)
	
		if err != nil {
			log.Println("Connection Closed!")
			log.Println(err)
			conn.ws.Close()
			break
		}

		log.Println("EVENT DISPATCHED: ", response.MessageType)
		go gocket.listeners[response.MessageType](conn, response.Data)
	}
}

func On(mType string, callback func(c Conn, d interface{})) {
	gocket.listeners[mType] = callback
	log.Println("NEW EVENT ATTACHED: ", mType)
}


func (conn *Conn) Emit(mType string, data interface{}) {
	log.Println("EMIT EVENT: ", mType)
	websocket.JSON.Send(&conn.ws, WebsocketResponse{
		MessageType: mType,
		Data: data,
	})
}

func (conn *Conn) EmitFor(room string, mType string, data interface{}) {
	log.Println("----> EMIT FOR CONN ID", conn.id)

	for _, connection := range gocket.connections {
		if conn.id != connection.id {
			log.Println("------> CONNECTIONS", connection.id)
			go connection.Emit(mType, data)
		}
	}
}

func (conn *Conn) ConnectOnRoom(room string) {
	conn.rooms = append(conn.rooms, room)
}

