package handlers

type ErrorMsg struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ErrorMsg) Error() string {
	return e.Message
}

func ErrorMessage(staus int, message string) ErrorMsg {
	return ErrorMsg{
		Status:  staus,
		Message: message,
	}
}
