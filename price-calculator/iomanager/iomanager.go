package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteJSON(data interface{}) error
}
