package application

import (
	"api/domain/entity"
	"api/domain/repository"

	"github.com/google/uuid"
)

type RoomService struct {
	room_repository repository.RoomRepository
}

func NewRoomService(room_repository repository.RoomRepository) *RoomService {
	room_service := new(RoomService)
	room_service.room_repository = room_repository
	return room_service
}

func (s *RoomService) CreateRoom() *entity.Room {
	room := entity.Room{
		Id:     uuid.NewString(),
		RoomId: uuid.NewString(),
	}
	s.room_repository.Save(&room)
	return &room
}
