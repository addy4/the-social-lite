package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"super.com/networking/data"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the protected handler")
}

func UnprotectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the unprotected handler")
}

func AddFriend(w http.ResponseWriter, r *http.Request) {

	// Read
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Read Error")
	}

	// Unmarshall
	message := &data.Request{}
	json.Unmarshal(body, &message)

	// Redis Client
	client := createRedisClient()
	defer client.Close()

	// Get User from DB using above message
	UserToAddAsFriend, err := getUserFromDB(client, message.Follow.UserName)
	if err != nil {
		fmt.Println("Cannot follow this user, not in our database")
		w.WriteHeader(400)
		return
	}

	// Remove Friend Requests, Friends
	UserToAddAsFriend.FriendRequests = nil
	UserToAddAsFriend.Friends = nil

	// Adding as friend
	addFriendToUser(client, message.Follow.CurrentUserName, UserToAddAsFriend)

	w.WriteHeader(201)
	response, _ := json.Marshal(message)
	w.Write(response)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// Read
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Read Error")
		return
	}

	// Unmarshall
	message := &data.Request{}
	json.Unmarshal(body, &message)

	// Creating User struct using above JSON
	registeredUser := data.User{
		UserName: message.Register.UserName,
		Password: message.Register.Password,
	}

	// Redis Client
	client := createRedisClient()

	defer client.Close()

	// Save
	addUserToDB(client, registeredUser)

	response, _ := json.Marshal(registeredUser)

	w.WriteHeader(201)
	w.Write(response)
}

func GetFriends(w http.ResponseWriter, r *http.Request) {

	// Read
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Read Error")
		return
	}

	// Unmarshall
	message := &data.Request{}
	json.Unmarshal(body, &message)

	// Redis
	client := createRedisClient()
	defer client.Close()

	// Remove Unnecessary fields
	currentUser, err := getUserFromDB(client, message.GetFriends.CurrentUserName)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	currentUser.UserName = ""
	currentUser.EmailID = ""

	// Convert Response to JSON
	response, err := json.Marshal(currentUser)

	// Write Back
	w.WriteHeader(200)
	w.Write(response)
}
