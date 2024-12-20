package routes

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/events/handlers"
	"github.com/gin-gonic/gin"
)

type EventRouter struct {
	Connections *infrastructure.Connections
	engine      *gin.Engine
}

func NewEventRouter(connections *infrastructure.Connections, engine *gin.Engine) *EventRouter {
	return &EventRouter{
		connections,
		engine,
	}
}

func (er *EventRouter) initDeps(connections *infrastructure.Connections) *application.EventService {
	eventRepo := postgres.NewEventRepository(connections)
	userRepo := postgres.NewUserRepository(connections)
	virtualRoomRepo := postgres.NewVirtualRoomRepository(connections)

	return application.NewEventService(eventRepo, userRepo, virtualRoomRepo)
}

func (er *EventRouter) RegisterRoutes() {
	apiGroup := er.engine.Group("/events")

	service := er.initDeps(er.Connections)

	handler := handlers.NewHandler(service)

	{
		apiGroup.POST("/create", handler.CreateEvent)
		apiGroup.GET("/:name", handler.GetEventByName)
		apiGroup.GET("/all", handler.GetAllEvents)
		apiGroup.DELETE("/:name", handler.DeleteEvent)
		apiGroup.PUT("/:name", handler.UpdateEvent)
	}

}
