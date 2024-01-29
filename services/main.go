package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	//"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"super.com/networking/data"
	"super.com/networking/routes"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var addr = flag.String("addr", "localhost:8089", "ws service address")

func main() {

	data.Register()

	data.ServerConnections = data.Connections{
		Ch: make(chan string),
	}

	go broadcast()

	mux := mux.NewRouter()
	mux.HandleFunc("/api/register", routes.RegisterUser).Methods("POST")
	mux.HandleFunc("/api/friends", routes.AddFriend).Methods("POST")
	mux.HandleFunc("/api/friends", routes.GetFriends).Methods("GET")
	mux.HandleFunc("/api/party", routes.CreateParty).Methods("POST")
	mux.HandleFunc("/api/party/membership", routes.AddMemberToTheParty).Methods("POST")
	mux.HandleFunc("/api/party/members", routes.GetPartyMembers).Methods("GET")

	http.HandleFunc("/play", play)

	go func() {
		log.Fatal(http.ListenAndServe(*addr, nil))
	}()

	srv := &http.Server{
		Addr:         ":4010",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	log.Fatal(err)

}

func play(w http.ResponseWriter, r *http.Request) {

	username, _, _ := r.BasicAuth()

	// WebSocket Upgrade
	c, err := upgrader.Upgrade(w, r, nil)

	data.ServerConnections.WsConnections = append(data.ServerConnections.WsConnections, c)
	fmt.Println("... here")
	data.ServerConnections.Ch <- username

	if err != nil {
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
	}

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

	defer c.Close()

}

func broadcast() {

	for {
		newUser := <-data.ServerConnections.Ch
		for _, clients := range data.ServerConnections.WsConnections {

			msg := map[string]interface{}{
				"Online": newUser,
			}

			clients.WriteJSON(msg)
		}
	}

}
