package utils

type Request struct {
	Message Message
	Writer  Writer
}

func NewRequest(msg Message, writer Writer) *Request {
	return &Request{
		Message: msg,
		Writer:  writer,
	}
}
