package virtual_rooms

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/virtual_rooms/virtual_room_dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VirtualRoomsHandlers struct {
	Service *application.VirtualRoomService
}

func NewHandler(service *application.VirtualRoomService) *VirtualRoomsHandlers {
	return &VirtualRoomsHandlers{
		Service: service,
	}
}

// CreateVirtualRoom создает новую виртуальную комнату
// @Summary      Создать виртуальную комнату
// @Description  Добавляет новую виртуальную комнату в систему
// @Tags         VirtualRooms
// @Accept       json
// @Produce      json
// @Param        room  body      virtual_room_dto.CreateVirtualRoom  true  "Данные комнаты"
// @Success      200   {string}  string                              "Название комнаты"
// @Failure      400   {object}  map[string]interface{}                              "Ошибка валидации"
// @Failure      500   {object}  map[string]interface{}                              "Ошибка сервера"
// @Router       /virtualRooms/create [post]
func (vrh *VirtualRoomsHandlers) CreateVirtualRoom(ctx *gin.Context) {
	dto := virtual_room_dto.CreateVirtualRoom{}

	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	if dto.Capacity <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid capacity value"})
		return
	}

	if err = vrh.Service.Create(dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Virtual room created successfully", "room_name": dto.RoomName})
}

// GetByName возвращает комнату по названию
// @Summary      Получить виртуальную комнату
// @Description  Возвращает данные виртуальной комнаты по её названию
// @Tags         VirtualRooms
// @Produce      json
// @Param        name  path      string  true  "Название комнаты"
// @Success      200   {object}  map[string]interface{}   "Информация о комнате"
// @Failure      500   {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /virtualRooms/{name} [get]
func (vrh *VirtualRoomsHandlers) GetByName(ctx *gin.Context) {
	name := ctx.Param("name")

	room, err := vrh.Service.GetByName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"room": room})
}

// GetAllVirtualRooms возвращает все комнаты
// @Summary      Список виртуальных комнат
// @Description  Возвращает список всех виртуальных комнат
// @Tags         VirtualRooms
// @Produce      json
// @Success      200   {object}  map[string]interface{}   "Список виртуальных комнат"
// @Failure      500   {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /virtualRooms/all [get]
func (vrh *VirtualRoomsHandlers) GetAllVirtualRooms(ctx *gin.Context) {
	rooms, err := vrh.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

// DeleteVirtualRoom удаляет комнату
// @Summary      Удалить виртуальную комнату
// @Description  Удаляет виртуальную комнату по названию
// @Tags         VirtualRooms
// @Param        name  path      string  true  "Название комнаты"
// @Success      200   {object}  map[string]interface{}  "Название удалённой комнаты"
// @Failure      500   {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /virtualRooms/{name} [delete]
func (vrh *VirtualRoomsHandlers) DeleteVirtualRoom(ctx *gin.Context) {
	name := ctx.Param("name")

	if err := vrh.Service.Delete(name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully", "room_name": name})
}

// UpdateVirtualRoom обновляет данные комнаты
// @Summary      Обновить виртуальную комнату
// @Description  Обновляет данные виртуальной комнаты по названию
// @Tags         VirtualRooms
// @Accept       json
// @Produce      json
// @Param        name  path      string               true  "Название комнаты"
// @Param        data  body      map[string]interface{} true  "Обновляемые данные"
// @Success      200   {object}  map[string]interface{}   "Обновлённые данные комнаты"
// @Failure      400   {object}  map[string]interface{}               "Ошибка валидации"
// @Failure      500   {object}  map[string]interface{}               "Ошибка сервера"
// @Router       /virtualRooms/update/{name} [put]
func (vrh *VirtualRoomsHandlers) UpdateVirtualRoom(ctx *gin.Context) {
	name := ctx.Param("name")

	existingRoom, err := vrh.Service.GetByName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Room not found"})
		return
	}

	var input map[string]interface{}
	if err = ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if roomName, ok := input["room_name"].(string); ok {
		existingRoom.RoomName = roomName
	}
	if capacity, ok := input["capacity"].(float64); ok {
		existingRoom.Capacity = int(capacity)
	}

	if err = vrh.Service.Update(existingRoom); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Room updated successfully", "room": existingRoom})
}
