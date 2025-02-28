package constant

const (
	UserInfoJoinTypeJOIN = iota + 1
	UserInfoJoinTypeAPPLY
)

const (
	FriendResponseDefault = 0
	FriendResponseAgree   = 1
	FriendResponseRefuse  = -1

	GroupResponseDefault = 0
	GroupResponseAgree   = 1
	GroupResponseRefuse  = -1

	GroupOrdinaryUser = 1
	GroupOwner        = 2
	GroupAdmin        = 3

	Male   = 1
	Female = 2
)
