package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type (
	ServerOpt func(*server)

	route struct {
		method, fullPath string
	}

	server struct {
		engine *fiber.App
		prefix string
		routes map[string]route
	}

	Server interface {
		AddHandler(
			path, group, method string,
			handler fiber.Handler,
			middlewares ...fiber.Handler,
		)
		Use(group string, middlewares ...fiber.Handler)
		Start(port string) error
		PrintRouter()
	}
)

func WithPrefix(prefix string) ServerOpt {
	return func(s *server) {
		s.prefix = prefix
	}
}

func (s *server) Use(group string, middlewares ...fiber.Handler) {
	fullPath := s.prefix + group

	for _, middleware := range middlewares {
		s.engine.Use(fullPath, middleware)
	}
}

func (s *server) AddHandler(path, group, method string, handler fiber.Handler, middlewares ...fiber.Handler) {
	fullPath := s.prefix + group + path

	_, ok := s.routes[fullPath+method]
	if ok {
		log.Println("[server.AddHandler] there is already a handler for this path")
		return
	}

	s.routes[fullPath+method] = route{
		method:   method,
		fullPath: fullPath,
	}

	for _, middleware := range middlewares {
		s.engine.Use(fullPath, middleware)
	}
	s.engine.Add(method, fullPath, handler)
}

func (s *server) PrintRouter() {
	for _, route := range s.routes {
		log.Printf("[METHOD] %s - [PATH] %s", route.method, route.fullPath)
	}
}

func (s *server) Start(port string) error {
	return s.engine.Listen(port)
}

func NewServer(opts ...ServerOpt) Server {
	s := &server{
		routes: make(map[string]route),
	}

	s.engine = fiber.New()

	for _, opt := range opts {
		opt(s)
	}

	return s
}
