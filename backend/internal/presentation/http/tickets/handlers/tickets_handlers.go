package handlers

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/domain/ticket_type/entity"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/tickets/ticket_dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TicketHandler struct {
	Service *application.TicketService
}

func NewHandler(service *application.TicketService) *TicketHandler {
	return &TicketHandler{Service: service}
}

// CreateTicket создаёт новый билет
// @Summary      Создать билет
// @Description  Создаёт новый билет для события
// @Tags         Tickets
// @Accept       json
// @Produce      json
// @Param        ticket  body      ticket_dto.CreateTicket  true  "Данные билета"
// @Success      200     {object}  map[string]interface{}  "Успешное создание билета"
// @Failure      400     {object}  map[string]interface{}  "Ошибка валидации"
// @Failure      500     {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /tickets/create [post]
func (th *TicketHandler) CreateTicket(ctx *gin.Context) {
	dto := ticket_dto.CreateTicket{}

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ticket not found", "details": err.Error()})
		return
	}

	if err := th.Service.CreateTicket(dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ticket not found", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ticket creation is successful", "event": dto.EventName, "user": dto.UserLogin})
}

// GetTicketByUserLoginAndEventName возвращает билет по логину пользователя и названию события
// @Summary      Получить билет
// @Description  Возвращает билет по логину пользователя и названию события
// @Tags         Tickets
// @Produce      json
// @Param        userLogin  path      string                true  "Логин пользователя"
// @Param        eventName  path      string                true  "Название события"
// @Success      200        {object}  map[string]interface{}  "Информация о билете"
// @Failure      500        {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /tickets/{login}/{event_name} [get]
func (th *TicketHandler) GetTicketByUserLoginAndEventName(ctx *gin.Context) {
	userLogin := ctx.Param("login")
	eventName := ctx.Param("event_name")

	ticket, err := th.Service.GetTicket(userLogin, eventName)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ticket not found", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ticket": ticket})
}

// GetAllTickets возвращает все билеты
// @Summary      Список билетов
// @Description  Возвращает все билеты
// @Tags         Tickets
// @Produce      json
// @Success      200   {object}  map[string]interface{}  "Список билетов"
// @Failure      500   {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /tickets/all [get]
func (th *TicketHandler) GetAllTickets(ctx *gin.Context) {
	tickets, err := th.Service.GetAllTickets()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ticket not found", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tickets": tickets})
}

// DeleteTicket удаляет билет
// @Summary      Удалить билет
// @Description  Удаляет билет по логину пользователя и названию события
// @Tags         Tickets
// @Param        userLogin  path      string  true  "Логин пользователя"
// @Param        eventName  path      string  true  "Название события"
// @Success      200        {object}  map[string]interface{}  "Успешное удаление"
// @Failure      500        {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /tickets/{login}/{event_name} [delete]
func (th *TicketHandler) DeleteTicket(ctx *gin.Context) {
	userLogin := ctx.Param("login")
	eventName := ctx.Param("event_name")

	if err := th.Service.DeleteTicket(userLogin, eventName); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ticket not found", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ticket deletion is successful", "event": eventName, "user": userLogin})
}

// UpdateTicket обновляет билет
// @Summary      Обновить билет
// @Description  Обновляет данные билета
// @Tags         Tickets
// @Accept       json
// @Produce      json
// @Param        userLogin      path      string                true  "Логин пользователя"
// @Param        eventName      path      string                true  "Название события"
// @Param        updatedTicket  body      map[string]interface{} true  "Обновляемые данные"
// @Success      200            {object}  map[string]interface{}  "Успешное обновление"
// @Failure      400            {object}  map[string]interface{}  "Ошибка валидации"
// @Failure      500            {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /tickets/{login}/{event_name} [put]
func (th *TicketHandler) UpdateTicket(ctx *gin.Context) {
	userLogin := ctx.Param("login")
	eventName := ctx.Param("event_name")

	existingTicket, err := th.Service.GetTicket(userLogin, eventName)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ticket not found", "details": err.Error()})
		return
	}

	var input map[string]interface{}

	if err = ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	if newTicketType, ok := input["ticket_type"].(entity.TicketType); ok {
		existingTicket.TicketType = newTicketType
	}

	if newPrice, ok := input["price"].(float64); ok {
		existingTicket.Price = newPrice
	}

	if newPurchaseDate, ok := input["purchase_date"].(time.Time); ok {
		existingTicket.PurchaseDate = newPurchaseDate
	}

	if err = th.Service.UpdateTicket(existingTicket); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ticket not found", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ticket update is successful", "event": eventName, "user": userLogin})
}
