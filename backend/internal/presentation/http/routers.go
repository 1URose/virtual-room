package ginConfig

import (
	"git.ai-space.tech/coursework/backend/docs"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
}

func RegisterRoutes(engine *gin.Engine, connections *infrastructure.Connections) {
	docs.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRouter := routes.NewUserRouter(connections, engine)
	virtualRoomRouter := routes.NewVirtualRoomRouter(connections, engine)
	eventRouter := routes.NewEventRouter(connections, engine)
	ticketRouter := routes.NewTicketRouter(connections, engine)
	equipmentRouter := routes.NewEquipmentRouter(connections, engine)
	eventParticipantRouter := routes.NewEventParticipantRouter(connections, engine)
	sponsorRouter := routes.NewSponsorRouter(connections, engine)

	userRouter.RegisterRoutes()
	virtualRoomRouter.RegisterRoutes()
	eventRouter.RegisterRoutes()
	ticketRouter.RegisterRoutes()
	equipmentRouter.RegisterRoutes()
	eventParticipantRouter.RegisterRoutes()
	sponsorRouter.RegisterRoutes()
}
