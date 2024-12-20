package validator

import (
	"errors"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/events/event_dto"
)

type EventValidator struct {
	eDto event_dto.CreateEvent
}

func NewEventValidator(eDto event_dto.CreateEvent) *EventValidator {
	return &EventValidator{eDto: eDto}
}

func (ev *EventValidator) TimeValidate() error {
	if ev.eDto.StartDate.After(ev.eDto.EndDate) {
		return errors.New("время начала не может быть позже времени окончания")
	}
	return nil
}

func (ev *EventValidator) TitleValidate() error {
	if ev.eDto.Title == "" {
		return errors.New("заголовок события не может быть пустым")
	}
	return nil
}

func (ev *EventValidator) DescriptionValidate() error {
	if ev.eDto.Description == "" {
		return errors.New("описание события не может быть пустым")
	}
	return nil
}

func (ev *EventValidator) VirtualRoomNameValidate() error {
	if ev.eDto.VirtualRoomName == "" {
		return errors.New("название виртуальной комнаты не может быть пустым")
	}
	return nil
}

func (ev *EventValidator) OrganizerLoginValidate() error {
	if ev.eDto.OrganizerLogin == "" {
		return errors.New("логин организатора не может быть пустым")
	}
	return nil
}

// Validate выполняет все проверки сразу
func (ev *EventValidator) Validate() error {
	if err := ev.TitleValidate(); err != nil {
		return err
	}
	if err := ev.DescriptionValidate(); err != nil {
		return err
	}
	if err := ev.VirtualRoomNameValidate(); err != nil {
		return err
	}
	if err := ev.OrganizerLoginValidate(); err != nil {
		return err
	}
	if err := ev.TimeValidate(); err != nil {
		return err
	}
	return nil
}
