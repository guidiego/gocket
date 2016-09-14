package main;

import (	
	"log"
	"net/http"
	"github.com/guidiego/gocket/test"
)

func main() {
	gocket.On("HU3", func(data interface{}, send func(m string, d interface{})) {
		send("HU3_RESPONSE", data)
	})

	gocket.On("MICHAEL", func(data interface{}, send func(m string, d interface{})) {
		send("MICHAEL_GAY", data)
	})

	http.Handle("/ws", gocket.SocketHandler())

	log.Fatal(http.ListenAndServe(":8888", nil))
}