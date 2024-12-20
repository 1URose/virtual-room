package routes

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/sponsors/handlers"
	"github.com/gin-gonic/gin"
)

type SponsorRouter struct {
	Connections *infrastructure.Connections
	engine      *gin.Engine
}

func NewSponsorRouter(connections *infrastructure.Connections, engine *gin.Engine) *SponsorRouter {
	return &SponsorRouter{
		connections,
		engine,
	}
}

func (sr *SponsorRouter) initDeps(connections *infrastructure.Connections) *application.SponsorService {
	sponsorRepo := postgres.NewSponsorRepository(connections)
	eventRepo := postgres.NewEventRepository(connections)

	return application.NewSponsorService(sponsorRepo, eventRepo)
}

func (sr *SponsorRouter) RegisterRoutes() {
	apiGroup := sr.engine.Group("/sponsors")

	service := sr.initDeps(sr.Connections)

	handler := handlers.NewHandler(service)

	{
		apiGroup.POST("/create", handler.CreateSponsor)
		apiGroup.GET("/:name", handler.GetSponsorByName)
		apiGroup.GET("/all", handler.GetAllSponsors)
		apiGroup.DELETE("/:sponsorName", handler.DeleteSponsor)
		apiGroup.PUT("/:name", handler.UpdateSponsor)
	}
}
