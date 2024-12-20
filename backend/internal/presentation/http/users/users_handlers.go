package users

import (
	"errors"
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/domain/user_role/entity"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres/pg_errors"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/users/user_dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type UserHandler struct {
	Service *application.UserService
}

func NewHandler(service *application.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// Auth аутентификация пользователя
// @Summary      Аутентификация пользователя
// @Description  Проверяет учетные данные пользователя и возвращает токен
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        credentials  body      user_dto.Auth  true  "Учетные данные пользователя"
// @Success      200          {object}  map[string]interface{}   "Токен аутентификации"
// @Failure      400          {object}  map[string]interface{}  "Ошибка валидации"
// @Failure      404          {object}  map[string]interface{}  "Пользователь не найден"
// @Failure      500          {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /users/auth [post]
func (h *UserHandler) Auth(ctx *gin.Context) {
	dto := user_dto.Auth{}

	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	creds, err := h.Service.Auth(dto)
	if err != nil {
		if errors.Is(err, pg_errors.NewUserNotFoundError(dto.Email)) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": creds})
}

// SignUp регистрация нового пользователя
// @Summary      Регистрация пользователя
// @Description  Регистрирует нового пользователя в системе
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      user_dto.SignUp  true  "Данные пользователя"
// @Success      200   {object}  map[string]interface{}  "Успешная регистрация"
// @Failure      400   {object}  map[string]interface{}  "Ошибка валидации"
// @Failure      500   {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /users/signup [post]
func (h *UserHandler) SignUp(ctx *gin.Context) {
	dto := user_dto.SignUp{}

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	/*
		if !IsValidEmail(dto.Email) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
			return
		}

	*/if err := h.Service.SingUp(dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User creation is successful", "user": dto.Email})
}

// GetAllUsers возвращает список всех пользователей
// @Summary      Список пользователей
// @Description  Возвращает список всех зарегистрированных пользователей
// @Tags         Users
// @Produce      json
// @Success      200   {object}  map[string]interface{}   "Список пользователей"
// @Failure      500   {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /users/all [get]
func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUserByLogin возвращает пользователя по логину
// @Summary      Получить пользователя
// @Description  Возвращает данные пользователя по его логину
// @Tags         Users
// @Produce      json
// @Param        login  path      string  true  "Логин пользователя"
// @Success      200    {object}  map[string]interface{}   "Информация о пользователе"
// @Failure      500    {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /users/{login} [get]
func (h *UserHandler) GetUserByLogin(ctx *gin.Context) {
	login := ctx.Param("login")

	user, err := h.Service.GetUserByLogin(login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// DeleteUser удаляет пользователя
// @Summary      Удалить пользователя
// @Description  Удаляет пользователя по его логину
// @Tags         Users
// @Param        login  path      string  true  "Логин пользователя"
// @Success      200    {object}  map[string]interface{}  "Успешное удаление"
// @Failure      500    {object}  map[string]interface{}  "Ошибка сервера"
// @Router       /users/{login} [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	login := ctx.Param("login")

	if err := h.Service.DeleteUser(login); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deletion is successful", "login": login})
}

// UpdateUser обновляет данные пользователя
// @Summary      Обновить пользователя
// @Description  Обновляет данные пользователя по его логину
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        login  path      string               true  "Логин пользователя"
// @Param        data   body      map[string]interface{} true  "Обновляемые данные"
// @Success      200    {object}  map[string]interface{}               "Успешное обновление"
// @Failure      400    {object}  map[string]interface{}               "Ошибка валидации"
// @Failure      500    {object}  map[string]interface{}               "Ошибка сервера"
// @Router       /users/{login} [put]
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	login := ctx.Param("login")

	existingUser, err := h.Service.GetUserByLogin(login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	var input map[string]interface{}
	if err = ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if newName, ok := input["name"].(string); ok {
		existingUser.Name = newName
	}
	if newRole, ok := input["role"].(entity.UserRole); ok {
		existingUser.Role = newRole
	}

	if err = h.Service.UpdateUser(existingUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User update is successful", "login": login})
}

// IsValidEmail проверяет корректность email
func IsValidEmail(email string) bool {
	emailRegex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}
