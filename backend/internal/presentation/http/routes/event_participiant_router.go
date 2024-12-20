package routes

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/event_participiants"
	"github.com/gin-gonic/gin"
)

type EventParticipantRouter struct {
	Connections *infrastructure.Connections
	engine      *gin.Engine
}

func NewEventParticipantRouter(connections *infrastructure.Connections, engine *gin.Engine) *EventParticipantRouter {
	return &EventParticipantRouter{
		connections,
		engine,
	}
}

func initEventParticipantRouterDeps(connections *infrastructure.Connections) *application.EventParticipiantService {
	repo := postgres.NewEventParticipantRouter(connections)

	return application.NewEventParticipantService(repo)
}

func (vr *EventParticipantRouter) RegisterRoutes() {
	apiGroup := vr.engine.Group("/eventParticipants")

	service := initEventParticipantRouterDeps(vr.Connections)

	handler := event_participiants.NewHandler(service)

	{
		apiGroup.POST("/", handler.BecomeParticipant)
		apiGroup.DELETE("/", handler.LeaveEvent)
		apiGroup.POST("/getAllParticipantsByEvent", handler.GetAllEventParticipants)
		apiGroup.POST("/getAllEventsByParticipant", handler.GetAllEventsByParticipantId)
	}

}
