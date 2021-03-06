package repository

import (
	"api/domain/entity"
	"api/infrastructure/database"

	"github.com/guregu/dynamo"
)

type UserRepository struct {
	db *dynamo.DB
}

func NewUserRepository() *UserRepository {
	user_repository := new(UserRepository)
	user_repository.db = database.Connect()
	return user_repository
}

func (r *UserRepository) GetListByRoomId(roomId string) []entity.User {
	table := r.db.Table("shamo_user")

	var users []entity.User
	table.Scan().Filter("room_id = ?", roomId).All(&users)
	return users
}

func (r *UserRepository) Save(user *entity.User) {
	table := r.db.Table("shamo_user")

	table.Put(user).Run()
}

func (r *UserRepository) Exist(id string) bool {
	table := r.db.Table("shamo_user")

	cnt, _ := table.Get("id", id).Count()
	return cnt > 0
}
