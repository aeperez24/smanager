package httputils

type HttpMethod int

const (
	GET HttpMethod = iota
	POST
	PUT
	DELETE
)

func (httpMethod HttpMethod) String() string {
	return []string{"GET", "POST", "PUT", "DELETE"}[httpMethod]
}
