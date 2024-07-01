package handlers

type ErrorMsg struct {
	Status  int    `json:"status"`
	Message string `json:"meesage"`
}

func (e ErrorMsg) Error() string {
	return e.Message
}

func ErrorMessage(status int, message string) ErrorMsg {
	return ErrorMsg{
		Status:  status,
		Message: message,
	}
}
