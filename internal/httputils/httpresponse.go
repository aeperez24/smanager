package httputil

type HttpResponseDto[T any] struct {
	Data         T      `json:"data"`
	ErrorMessage string `json:"errorMessage"`
}
