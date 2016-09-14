package main;

import (	
	"log"
	"net/http"
	"github.com/guidiego/gocket/test"
)

func main() {

	gocket.On("Connect", func (conn gocket.Conn, data interface{}) {
		conn.ConnectOnRoom("test")
	})

	gocket.On("HU3", func (conn gocket.Conn, data interface{}) {
		conn.EmitFor("test", "HUE", data)
	})

	http.Handle("/ws", gocket.Handler())

	log.Fatal(http.ListenAndServe(":8888", nil))
}