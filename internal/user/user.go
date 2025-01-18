package user

type user struct {
	id uint
}

func NewUser(id uint) *user {
	return &user{
		id: id,
	}
}
