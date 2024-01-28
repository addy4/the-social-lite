package data

type User struct {
	UserName       string `json:"username,omitempty"`
	EmailID        string `json:"email,omitempty"`
	FriendRequests []User `json:"requests,omitempty"`
	Friends        []User `json:"friends,omitempty"`
	Password       string `json:"-"`
}

type Party struct {
	PartyTitle string `json:"partytitle,omitempty"`
	Members    []User `json:"members,omitempty"`
}
