package application

import (
	"fmt"
	"git.ai-space.tech/coursework/backend/internal/domain/virtual_room/entity"
	"git.ai-space.tech/coursework/backend/internal/domain/virtual_room/repository"
	presentation "git.ai-space.tech/coursework/backend/internal/presentation/http/virtual_rooms/virtual_room_dto"
	"golang.org/x/net/context"
)

type VirtualRoomService struct {
	repo repository.VirtualRoomRepository
	ctx  context.Context
}

func NewVirtualRoomService(repo repository.VirtualRoomRepository) *VirtualRoomService {
	return &VirtualRoomService{
		repo: repo,
	}
}

func (s *VirtualRoomService) Create(dto presentation.CreateVirtualRoom) error {
	if _, err := s.repo.GetByName(dto.RoomName); err == nil {
		return fmt.Errorf("комната с названием ''%s'' уже существует", dto.RoomName)
	}

	virtualRoom := entity.NewVirtualRoom(dto.RoomName, dto.Capacity)

	return s.repo.CreateVirtualRoom(virtualRoom)
}

func (s *VirtualRoomService) GetByName(name string) (*entity.VirtualRoom, error) {
	return s.repo.GetByName(name)
}

func (s *VirtualRoomService) GetAll() ([]entity.VirtualRoom, error) {
	return s.repo.GetAll()
}

func (s *VirtualRoomService) Delete(name string) error {
	return s.repo.DeleteByName(name)
}

func (s *VirtualRoomService) Update(room *entity.VirtualRoom) error {
	return s.repo.UpdateVirtualRoom(room)
}
