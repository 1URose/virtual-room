package routes

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/equipments/handlers"
	"github.com/gin-gonic/gin"
)

type EquipmentRouter struct {
	Connections *infrastructure.Connections
	engine      *gin.Engine
}

func NewEquipmentRouter(connections *infrastructure.Connections, engine *gin.Engine) *EquipmentRouter {
	return &EquipmentRouter{
		connections,
		engine,
	}
}

func (er *EquipmentRouter) initDeps(connections *infrastructure.Connections) *application.EquipmentService {
	equipmentRepo := postgres.NewEquipmentRepository(connections)
	userRepo := postgres.NewUserRepository(connections)

	return application.NewEquipmentService(equipmentRepo, userRepo)
}

func (er *EquipmentRouter) RegisterRoutes() {
	apiGroup := er.engine.Group("/equipment")

	service := er.initDeps(er.Connections)

	handler := handlers.NewHandler(service)

	{
		apiGroup.POST("/create", handler.CreateEquipment)
		apiGroup.GET("/:name", handler.GetUserByEquipmentName)
	}

}
