package exception

type UnauthorizedError struct{}

func (e *UnauthorizedError) Error() string {
	return "unauthorized access"
}

func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{}
}
