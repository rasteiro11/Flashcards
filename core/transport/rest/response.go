package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type (
	response[T any] struct {
		statusCode int
		body       T
		headers    http.Header
	}
)

func newResponse[T any](c *fiber.Ctx, data T, opts ...ResponseOpt[T]) *response[T] {
	res := &response[T]{
		body:    data,
		headers: make(http.Header),
	}

	for _, opt := range opts {
		opt(res)
	}

	return res
}

type ResponseOpt[T any] func(r *response[T])

func WithBody[T any](body T) ResponseOpt[T] {
	return func(res *response[T]) {
		res.body = body
	}
}

func WithStatusCode[T any](statusCode int) ResponseOpt[T] {
	return func(res *response[T]) {
		res.statusCode = statusCode
	}
}

func WithHeader[T any](key, value string) ResponseOpt[T] {
	return func(res *response[T]) {
		res.headers.Add(key, value)
	}
}

//func (r *response[T]) WithStatusCode(statusCode int) *response[T] {
//	return r
//}

//func NewStatusOkResponse[T any](c *fiber.Ctx, data T) error {
//	c.Response().SetStatusCode(http.StatusOK)
//	return c.JSON(data)
//}
