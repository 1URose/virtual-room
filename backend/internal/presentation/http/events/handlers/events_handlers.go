package handlers

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/events/event_dto"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/events/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type EventHandler struct {
	Service *application.EventService
}

func NewHandler(service *application.EventService) *EventHandler {
	return &EventHandler{Service: service}
}

// CreateEvent handles the creation of a new event.
// @Summary Create an event
// @Description Creates a new event with the provided details
// @Tags events
// @Accept json
// @Produce json
// @Param event body event_dto.CreateEvent true "Event details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /events/create [post]
func (eh *EventHandler) CreateEvent(ctx *gin.Context) {
	dto := event_dto.CreateEvent{}

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if err := validator.NewEventValidator(dto).Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "message": err})
		return
	}

	if err := eh.Service.Create(dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event creation is successful", "event": dto.Title})
}

// GetEventByName retrieves an event by its name.
// @Summary Get an event by name
// @Description Fetches event details by its name
// @Tags events
// @Accept json
// @Produce json
// @Param name path string true "Event name"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /events/{name} [get]
func (eh *EventHandler) GetEventByName(ctx *gin.Context) {
	name := ctx.Param("name")
	event, err := eh.Service.GetEventByName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event found", "event": event})
}

// GetAllEvents retrieves all events.
// @Summary Get all events
// @Description Retrieves a list of all events
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /events/all [get]
func (eh *EventHandler) GetAllEvents(ctx *gin.Context) {
	events, err := eh.Service.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Events retrieved successfully", "events": events})
}

// UpdateEvent updates an existing event.
// @Summary Update an event
// @Description Updates details of an existing event
// @Tags events
// @Accept json
// @Produce json
// @Param name path string true "Event name"
// @Param event body map[string]interface{} true "Event update details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /events/{name} [put]
func (eh *EventHandler) UpdateEvent(ctx *gin.Context) {
	name := ctx.Param("name")
	event, err := eh.Service.GetEventByName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found", "details": err.Error()})
		return
	}

	var input map[string]interface{}
	if err = ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if newName, ok := input["new_name"].(string); ok {
		event.EventName = newName
	}
	if newDescription, ok := input["new_description"].(string); ok {
		event.Description = newDescription
	}
	if newOrganizerLogin, ok := input["new_organizer_login"].(string); ok {
		organizer, err := eh.Service.UserRepository.GetByEmail(newOrganizerLogin)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Organizer not found", "details": err.Error()})
			return
		}
		event.OrganizerID = organizer.ID
	}
	if newStartTime, ok := input["new_start_time"].(time.Time); ok {
		event.StartTime = newStartTime
	}
	if newEndTime, ok := input["new_end_time"].(time.Time); ok {
		event.EndTime = newEndTime
	}
	if newVirtualRoomName, ok := input["new_virtual_room_name"].(string); ok {
		virtualRoom, err := eh.Service.VirtualRoomRepository.GetByName(newVirtualRoomName)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Virtual room not found", "details": err.Error()})
			return
		}
		event.VirtualRoomID = virtualRoom.ID
	}

	if err = eh.Service.UpdateEvent(event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": event})
}

// DeleteEvent deletes an existing event.
// @Summary Delete an event
// @Description Deletes an event by its name
// @Tags events
// @Accept json
// @Produce json
// @Param name path string true "Event name"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /events/{name} [delete]
func (eh *EventHandler) DeleteEvent(ctx *gin.Context) {
	name := ctx.Param("name")

	if err := eh.Service.DeleteEvent(name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully", "event_name": name})
}
