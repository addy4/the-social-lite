package main

import (
	"log"
	"net/http"

	//"os"
	"time"

	"github.com/gorilla/mux"
	"super.com/networking/data"
	"super.com/networking/routes"
)

func main() {

	data.Register()

	mux := mux.NewRouter()
	mux.HandleFunc("/api/register", routes.RegisterUser).Methods("POST")
	mux.HandleFunc("/api/friends", routes.AddFriend).Methods("POST")
	mux.HandleFunc("/api/friends", routes.GetFriends).Methods("GET")
	mux.HandleFunc("/api/party", routes.CreateParty).Methods("POST")
	mux.HandleFunc("/api/party/membership", routes.AddMemberToTheParty).Methods("POST")
	mux.HandleFunc("/api/party/members", routes.GetPartyMembers).Methods("GET")

	srv := &http.Server{
		Addr:         ":4010",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//routes.RedisClientDemo()
	//routes.MainRedis()
	//routes.Testing()

	log.Printf("starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	log.Fatal(err)

}
