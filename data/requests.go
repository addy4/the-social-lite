package data

type Request struct {
	Follow          FollowRequest          `json:"follow,omitempty"`
	Register        RegisterRequest        `json:"register,omitempty"`
	GetFriends      GetFriendsRequest      `json:"getfriends,omitempty"`
	CreateNewParty  CreatePartyRequest     `json:"createparty,omitempty"`
	AddToParty      AddToPartyRequest      `json:"addtoparty,omitempty"`
	GetPartyMembers GetPartyMembersRequest `json:"getpartymembers,omitempty"`
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

type CreatePartyRequest struct {
	PartyTitle string `json:"partytitle"`
}

type AddToPartyRequest struct {
	PartyTitle string `json:"partytitle,omitempty"`
	UserName   string `json:"username,omitempty"`
}

type GetPartyMembersRequest struct {
	PartyTitle string `json:"partytitle,omitempty"`
}
