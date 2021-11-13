package service

import (
	"api/app/model"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func CreateUser(user *model.User) *model.User {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_DEFAULT_REGION"))})
	table := db.Table("shamo_user")

	err := table.Put(user).Run()

	if err != nil {
		println(err.Error())
	}

	return user
}

func GetUsersByRoomId(roomId string) []model.User {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_DEFAULT_REGION"))})
	table := db.Table("shamo_user")

	var users []model.User
	table.Scan().Filter("room_id = ?", roomId).All(&users)

	return users
}
