package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"super.com/networking/data"
)

var ctxx = context.Background()

func createRedisClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

// Add User to redis DB
func addUserToDB(client *redis.Client, user data.User) error {

	// User to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Store User as JSON in Hash "users" with key username
	err = client.HSet(ctxx, "users", user.UserName, userJSON).Err()
	if err != nil {
		return err
	}

	return nil
}

// Modify User such that friend/User is added to Friends[]data.User and then marshalled again as JSON which will be stored to redis DB using above addUserToDB
func addFriendToUser(client *redis.Client, username string, newFriend data.User) error {

	status, err := client.HGet(ctxx, "links", fmt.Sprintf("%s:%s", username, newFriend.UserName)).Result()
	if err == redis.Nil {
		log.Printf("Link does not exist between %s:%s", username, newFriend.UserName)
	} else if status == "true" {
		log.Printf("Link already exists between %s:%s", username, newFriend.UserName)
		return nil
	} else if status != "true" {
		log.Printf("Link does not exist between %s:%s", username, newFriend.UserName)
	}

	// Retrieve User JSON from Redis hash "users"
	userJSON, err := client.HGet(ctxx, "users", username).Result()
	if err != nil {
		return err
	}

	// Convert JSON to User struct
	var user data.User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return err
	}

	// Add new friend/User to the User's Friends[] list
	user.Friends = append(user.Friends, newFriend)

	// Re-Insert User using addUserToDB call
	err = addUserToDB(client, user)
	if err != nil {
		return err
	}

	// Store link in "links" redis hash
	err = client.HSet(ctxx, "links", fmt.Sprintf("%s:%s", username, newFriend.UserName), "true").Err()
	if err != nil {
		return err
	}

	return nil
}

// Gets JSON with key as username which will then be converted to User struct using Unmarshalling
func getUserFromDB(client *redis.Client, username string) (data.User, error) {

	// Retrieve User as JSON from Redis hash "users"
	userJSON, err := client.HGet(ctxx, "users", username).Result()
	if err != nil {
		return data.User{}, err
	}

	// Convert JSON string to User struct
	var user data.User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return data.User{}, err
	}

	return user, nil
}
