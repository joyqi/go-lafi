package contact

type Group interface {
	MemberBelong(openId string)
}
