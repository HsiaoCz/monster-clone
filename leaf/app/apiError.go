package app

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (a APIError) Error() string {
	return a.Message
}

func NewAPIError(status int, message string) APIError {
	return APIError{
		Status:  status,
		Message: message,
	}
}
