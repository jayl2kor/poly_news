package errors

type PolyNewsError struct {
	error
	Code          int
	Message       string
	InternalError error
}
