package routes

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/virtual_rooms"
	"github.com/gin-gonic/gin"
)

type VirtualRoomRouter struct {
	Connections *infrastructure.Connections
	engine      *gin.Engine
}

func NewVirtualRoomRouter(connections *infrastructure.Connections, engine *gin.Engine) *VirtualRoomRouter {
	return &VirtualRoomRouter{
		connections,
		engine,
	}
}

func initVirtualRoomDeps(connections *infrastructure.Connections) *application.VirtualRoomService {
	repo := postgres.NewVirtualRoomRepository(connections)

	return application.NewVirtualRoomService(repo)
}

func (vr *VirtualRoomRouter) RegisterRoutes() {
	apiGroup := vr.engine.Group("/virtualRooms")

	service := initVirtualRoomDeps(vr.Connections)

	handler := virtual_rooms.NewHandler(service)

	{
		apiGroup.POST("/create", handler.CreateVirtualRoom)
		apiGroup.GET("/:name", handler.GetByName)
		apiGroup.GET("/all", handler.GetAllVirtualRooms)
		apiGroup.DELETE("/:name", handler.DeleteVirtualRoom)
		apiGroup.PUT("/update/:name", handler.UpdateVirtualRoom)
	}

}
