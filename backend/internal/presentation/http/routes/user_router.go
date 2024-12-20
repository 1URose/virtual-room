package routes

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/users"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	Connections *infrastructure.Connections
	engine      *gin.Engine
}

func NewUserRouter(connections *infrastructure.Connections, engine *gin.Engine) *UserRouter {
	return &UserRouter{
		connections,
		engine,
	}
}

func initUserDeps(connections *infrastructure.Connections) *application.UserService {
	repo := postgres.NewUserRepository(connections)

	return application.NewUserService(repo)
}

func (ur *UserRouter) RegisterRoutes() {
	apiGroup := ur.engine.Group("/users")

	service := initUserDeps(ur.Connections)

	handler := users.NewHandler(service)

	{
		apiGroup.POST("/auth", handler.Auth)
		apiGroup.POST("/signUp", handler.SignUp)
		apiGroup.GET("/all", handler.GetAllUsers)
		apiGroup.GET("/:login", handler.GetUserByLogin)
		apiGroup.DELETE("/:login", handler.DeleteUser)
		apiGroup.PUT("/:login", handler.UpdateUser)
	}
}
