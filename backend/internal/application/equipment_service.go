package application

import (
	"context"
	"fmt"
	"git.ai-space.tech/coursework/backend/internal/domain/equipment/entity"
	er "git.ai-space.tech/coursework/backend/internal/domain/equipment/repository"
	eu "git.ai-space.tech/coursework/backend/internal/domain/user/entity"
	ur "git.ai-space.tech/coursework/backend/internal/domain/user/repository"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/equipments/equipment_dto"
)

type EquipmentService struct {
	EquipmentRepo er.EquipmentRepository
	UserRepo      ur.UserRepository
	ctx           context.Context
}

func NewEquipmentService(equipmentRepo er.EquipmentRepository, userRepo ur.UserRepository) *EquipmentService {
	return &EquipmentService{
		EquipmentRepo: equipmentRepo,
		UserRepo:      userRepo,
	}
}

func (s *EquipmentService) CreateEquipment(dto equipment_dto.CreateEquipment) error {

	if _, err := s.EquipmentRepo.GetUserByEquipmentName(dto.Name); err == nil {
		return fmt.Errorf("оборудование с названием ''%s'' уже существует", dto.Name)
	}

	user, err := s.UserRepo.GetByEmail(dto.UserLogin)

	if err != nil {
		return err
	}

	equipment := entity.NewEquipment(dto.Name, user.ID)

	return s.EquipmentRepo.CreateEquipment(equipment)
}

func (s *EquipmentService) GetUserByEquipmentName(equipmentName string) (*eu.User, error) {
	return s.EquipmentRepo.GetUserByEquipmentName(equipmentName)
}
