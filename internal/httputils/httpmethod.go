package httputils

type HttpMethod int

const (
	GET HttpMethod = iota
	POST
	PUT
	DELETE
)
