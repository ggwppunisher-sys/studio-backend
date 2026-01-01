package domain

type TgUserInfo struct {
	TgId      int64
	FirstName string
	LastName  string
	Username  string
	TgChatId  int64
}

type User struct {
	Id int64

	TgUserInfo
}

type Engineer struct {
	User

	EngineerId int64
}
