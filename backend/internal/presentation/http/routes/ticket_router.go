package routes

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/tickets/handlers"
	"github.com/gin-gonic/gin"
)

type TicketRouter struct {
	Connections *infrastructure.Connections
	engine      *gin.Engine
}

func NewTicketRouter(connections *infrastructure.Connections, engine *gin.Engine) *TicketRouter {
	return &TicketRouter{
		connections,
		engine,
	}
}

func initDeps(conn *infrastructure.Connections) *application.TicketService {
	ticketRepo := postgres.NewTicketRepository(conn)
	eventRepo := postgres.NewEventRepository(conn)
	userRepo := postgres.NewUserRepository(conn)

	return application.NewTicketService(ticketRepo, eventRepo, userRepo)
}

func (t *TicketRouter) RegisterRoutes() {
	apiGroup := t.engine.Group("/tickets")

	service := initDeps(t.Connections)

	handler := handlers.NewHandler(service)

	{
		apiGroup.POST("/create", handler.CreateTicket)
		apiGroup.GET("/:login/:event_name", handler.GetTicketByUserLoginAndEventName)
		apiGroup.GET("/all", handler.GetAllTickets)
		apiGroup.DELETE("/:login/:event_name", handler.DeleteTicket)
		apiGroup.PUT("/:login/:event_name", handler.UpdateTicket)
	}
}
