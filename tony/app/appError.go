package app

type APPError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (a APPError) Error() string {
	return a.Message
}

func NewAPPError(status int, message string) APPError {
	return APPError{
		Status:  status,
		Message: message,
	}
}
