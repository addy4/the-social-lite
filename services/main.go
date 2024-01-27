package main

import (
	"fmt"
	"log"
	"net/http"

	//"os"
	"time"

	"github.com/gorilla/mux"
	"super.com/networking/auth"
	"super.com/networking/data"
	"super.com/networking/routes"
)

func main() {

	data.Register()

	mux := mux.NewRouter()
	mux.HandleFunc("/unprotected", ItProtectedHandler)
	mux.HandleFunc("/friends", auth.BasicAuth(routes.UnprotectedHandler)).Methods("GET")
	mux.HandleFunc("/api/friends", auth.BasicAuth(routes.AddFriend)).Methods("POST")
	mux.HandleFunc("/api/register", routes.RegisterUser).Methods("POST")
	mux.HandleFunc("/api/friends", routes.GetFriends).Methods("GET")

	srv := &http.Server{
		Addr:         ":4010",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//routes.RedisClientDemo()
	//routes.MainRedis()

	log.Printf("starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	log.Fatal(err)

}

func ItProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the unprotected handler")
}
