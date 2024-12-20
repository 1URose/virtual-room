package jwt_errors

type GenerateTokenError struct {
}

func NewGenerateTokenError() *GenerateTokenError {
	return &GenerateTokenError{}
}

func (e *GenerateTokenError) Error() string {
	return "Ошибка создания токена"
}
