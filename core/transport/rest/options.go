package rest

func WithBody[T any](body T) ResponseOpt[T] {
	return func(res *response[T]) {
		res.body = body
	}
}
