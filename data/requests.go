package data

type Request struct {
	Follow     FollowRequest     `json:"follow"`
	Register   RegisterRequest   `json:"register"`
	GetFriends GetFriendsRequest `json:"getfriends"`
}

type FollowRequest struct {
	CurrentUserName string `json:"currentuser"`
	UserName        string `json:"username"`
}

type RegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type GetFriendsRequest struct {
	CurrentUserName string `json:"currentuser"`
}
