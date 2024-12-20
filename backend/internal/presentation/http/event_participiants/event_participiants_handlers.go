package event_participiants

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/domain/event/entity"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/event_participiants/event_participiant_dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EventParticipiantHandler struct {
	Service *application.EventParticipiantService
}

func NewHandler(service *application.EventParticipiantService) *EventParticipiantHandler {
	return &EventParticipiantHandler{Service: service}
}

// BecomeParticipant handles the process of a user becoming a participant in an event.
// @Summary User becomes a participant in an event
// @Description Allows a user to join a specific event by becoming its participant
// @Tags event_participants
// @Accept json
// @Produce json
// @Param participant body event_participiant_dto.BecomeParticipantDto true "Event participant details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /eventParticipants [post]
func (h *EventParticipiantHandler) BecomeParticipant(ctx *gin.Context) {
	dto := event_participiant_dto.BecomeParticipantDto{}

	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	err = h.Service.BecomeParticipiant(&dto.Event, dto.UserId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// LeaveEvent handles the process of a user leaving an event.
// @Summary User leaves an event
// @Description Allows a user to leave a specific event they are participating in
// @Tags event_participants
// @Accept json
// @Produce json
// @Param leaveEvent body event_participiant_dto.LeaveEventDto true "Event leave details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /eventParticipants [delete]
func (h *EventParticipiantHandler) LeaveEvent(ctx *gin.Context) {
	dto := event_participiant_dto.LeaveEventDto{}

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})

		return
	}

	if err := h.Service.LeaveEvent(&dto.Event, dto.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user successfully leave event"})
}

// GetAllEventParticipants retrieves a list of all participants for a given event.
// @Summary Get all participants of an event
// @Description Fetches a list of users participating in a specified event
// @Tags event_participants
// @Accept json
// @Produce json
// @Param event body entity.Event true "Event details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /eventParticipants/getAllParticipantsByEvent [post]
func (h *EventParticipiantHandler) GetAllEventParticipants(ctx *gin.Context) {
	dto := entity.Event{}

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})

		return
	}

	users, err := h.Service.GetAllEventParticipants(&dto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

// GetAllEventsByParticipantId retrieves a list of all events a user is participating in.
// @Summary Get all events by participant ID
// @Description Fetches all events for a user identified by their participant ID
// @Tags event_participants
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /eventParticipants/getAllEventsByParticipant [get]
func (h *EventParticipiantHandler) GetAllEventsByParticipantId(ctx *gin.Context) {

	userIdString := ctx.Param("userId")

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})

		return
	}

	events, err := h.Service.GetAllEventsByParticipantId(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"events": events})
}
