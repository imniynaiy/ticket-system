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
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=30"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	AgentFlag  string `validate:"max=100"`
	FamilyName string `validate:"max=20"`
	FirstName  string `validate:"max=20"`
	Sex        int
	Tel        int
	Address    string `validate:"max=1000"`
	Email      string `validate:"required,email"`
	Password   string `validate:"required,min=8,max=30"`
}
