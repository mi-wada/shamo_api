package repository

import (
	"api/domain/entity"
	"api/infra"

	"github.com/guregu/dynamo"
)

type UserRepository struct {
	db *dynamo.DB
}

func NewUserRepository() *UserRepository {
	user_repository := new(UserRepository)
	user_repository.db = infra.ConnectDB()
	return user_repository
}

func (r *UserRepository) GetAllByRoomId(roomId string) []entity.User {
	table := r.db.Table("shamo_user")

	var users []entity.User
	table.Get("room_id", roomId).All(&users)
	return users
}

func (r *UserRepository) Save(user *entity.User) {
	table := r.db.Table("shamo_user")

	table.Put(user).Run()
}
