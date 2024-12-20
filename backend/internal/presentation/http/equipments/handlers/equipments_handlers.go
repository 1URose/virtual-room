package handlers

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/equipments/equipment_dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EquipmentHandler struct {
	Service *application.EquipmentService
}

func NewHandler(service *application.EquipmentService) *EquipmentHandler {
	return &EquipmentHandler{Service: service}
}

// CreateEquipment handles the creation of new equipment.
// @Summary Create a new equipment item
// @Description Creates a new equipment entry with the provided details
// @Tags equipment
// @Accept json
// @Produce json
// @Param equipment body equipment_dto.CreateEquipment true "Equipment details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /equipments/create [post]
func (eh *EquipmentHandler) CreateEquipment(ctx *gin.Context) {
	dto := equipment_dto.CreateEquipment{}

	if err := ctx.ShouldBindJSON(&dto); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})

		return
	}

	if err := eh.Service.CreateEquipment(dto); err != nil {

		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"equipment creation is successful": dto.Name})
}

// GetUserByEquipmentName retrieves the user associated with a specific equipment by its name.
// @Summary Get user by equipment name
// @Description Fetches the user who is associated with the specified equipment name
// @Tags equipment
// @Produce json
// @Param name path string true "Equipment name"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /equipments/{name} [get]
func (eh *EquipmentHandler) GetUserByEquipmentName(ctx *gin.Context) {
	name := ctx.Param("name")

	user, err := eh.Service.GetUserByEquipmentName(name)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, user)
}
