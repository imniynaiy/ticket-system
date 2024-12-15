package model

type User struct {
	UserId       uint
	AgentFlag    string
	FamilyName   string
	FirstName    string
	Sex          int
	Tel          int
	Address      string
	Email        string
	PasswordHash string
	IsAdmin      bool
}

func (User) TableName() string {
	return "user"
}

type LoginReq struct {
	Email    string
	Password string
}

type LoginResp struct {
	Token string
}

type RegisterReq struct {
	AgentFlag  string
	FamilyName string
	FirstName  string
	Sex        int
	Tel        int
	Address    string
	Email      string
	Password   string
}
