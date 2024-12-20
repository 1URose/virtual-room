package pg_errors

type UserNotFoundError struct {
	Email string
}

func NewUserNotFoundError(email string) *UserNotFoundError {
	return &UserNotFoundError{
		Email: email,
	}
}

func (e *UserNotFoundError) Error() string {
	return "User not found " + e.Email
}
