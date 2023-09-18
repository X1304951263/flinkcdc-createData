package model

type User struct {
	Id   int
	Name string
}

func (u User) TableName() string {
	return "user1"
}
