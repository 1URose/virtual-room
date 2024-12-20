package repository

import (
	"git.ai-space.tech/coursework/backend/internal/domain/equipment/entity"
	eu "git.ai-space.tech/coursework/backend/internal/domain/user/entity"
)

type EquipmentRepository interface {
	CreateEquipment(equipment *entity.Equipment) error
	GetUserByEquipmentName(equipmentName string) (*eu.User, error)
}
