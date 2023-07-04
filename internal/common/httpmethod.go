package common

type HttpMethod int

const (
	GET HttpMethod = iota
	POST
	PUT
	DELETE
)
