package contact

type UserIdType string

const (
	UserIdTypeUserId  UserIdType = "user_id"
	UserIdTypeOpenId             = "open_id"
	UserIdTypeUnionId            = "union_id"
)
