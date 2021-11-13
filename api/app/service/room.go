package service

import (
	"api/app/lib"
	"api/app/model"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func CreateRoom() *model.Room {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_DEFAULT_REGION"))})
	table := db.Table("shamo_room")

	roomId := lib.CreateRandomStr(32)
	room := model.Room{Id: 1, Room_id: roomId}
	table.Put(room).Run()

	return &room
}
