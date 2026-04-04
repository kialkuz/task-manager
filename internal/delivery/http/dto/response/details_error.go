package dto

type DetailsErrorResponse[T any] struct {
	ErrorText string `json:"error"`
	Details   T      `json:"details"`
}

func NewDetailsError[T any](errorText string, details T) DetailsErrorResponse[T] {
	return DetailsErrorResponse[T]{ErrorText: errorText, Details: details}
}
