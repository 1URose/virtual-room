package repository

import "git.ai-space.tech/coursework/backend/internal/domain/virtual_room/entity"

type VirtualRoomRepository interface {
	CreateVirtualRoom(virtualRoom *entity.VirtualRoom) error
	GetByName(name string) (*entity.VirtualRoom, error)
	GetAll() ([]entity.VirtualRoom, error)
	DeleteByName(name string) error
	UpdateVirtualRoom(virtualRoom *entity.VirtualRoom) error
}
