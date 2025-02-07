package constant

const (
	UserInfoJoinTypeJOIN = iota + 1
	UserInfoJoinTypeAPPLY
)

const (
	FriendRequestDefault = 0
	FriendResponseAgree  = 1
	FriendResponseRefuse = -1

	GroupResponseAgree  = 1
	GroupResponseRefuse = -1

	GroupOrdinaryUser = 1
	GroupOwner        = 2
	GroupAdmin        = 3

	Male   = 1
	Female = 2
)
