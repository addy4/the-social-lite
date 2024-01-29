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

func Testing() {

	client := createRedisClient()
	defer client.Close()

	newParty := data.Party{
		PartyTitle: "NewGame",
		Members:    nil,
	}

	addPartyToDB(client, newParty)

	partyForAdding, _ := getPartyFromDB(client, "NewGame")

	fmt.Println(partyForAdding.PartyTitle)

	member, _ := getUserFromDB(client, "kohli")

	fmt.Println(member.UserName)

	addMemberToParty(client, partyForAdding.PartyTitle, member)

}

func createRedisClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
		Password: "",
		DB:       0,
	})

	return client
}

// Add Party to redis DB
func addPartyToDB(client *redis.Client, party data.Party) error {

	// Party To JSON
	partyJSON, err := json.Marshal(party)
	if err != nil {
		return err
	}

	// Store Party as JSON in Hash "parties" with key party title
	err = client.HSet(ctxx, "parties", party.PartyTitle, partyJSON).Err()
	if err != nil {
		return err
	}

	return nil
}

// Gets JSON with key as party title which will then be converted to Party struct using Unmarshalling
func getPartyFromDB(client *redis.Client, partytitle string) (data.Party, error) {

	// Retrieve User as JSON from Redis hash "users"
	partyJSON, err := client.HGet(ctxx, "parties", partytitle).Result()
	if err != nil {
		return data.Party{}, err
	}

	// Convert JSON string to User struct
	var party data.Party
	err = json.Unmarshal([]byte(partyJSON), &party)
	if err != nil {
		return data.Party{}, err
	}

	return party, nil
}

// Modify Party such that friend/User is added to Members[]data.User of party struct then saved again as JSON
func addMemberToParty(client *redis.Client, partytitle string, newFriend data.User) error {

	// Retrieve User JSON from Redis hash "users"
	partyJSON, err := client.HGet(ctxx, "parties", partytitle).Result()
	if err != nil {
		return err
	}

	// Convert JSON to User struct
	var party data.Party
	err = json.Unmarshal([]byte(partyJSON), &party)
	if err != nil {
		return err
	}

	// Add new friend/User to the User's Friends[] list
	party.Members = append(party.Members, newFriend)

	// Re-Insert User using addUserToDB call
	err = addPartyToDB(client, party)
	if err != nil {
		return err
	}

	return nil
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
