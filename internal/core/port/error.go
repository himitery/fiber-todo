package port

type PortError struct {
	Code    int
	Message string
}

func (e *PortError) Error() string {
	return e.Message
}
