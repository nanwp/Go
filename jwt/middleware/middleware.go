package middleware

import "net/http"

type MuxServer struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *MuxServer) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *MuxServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}
	current.ServeHTTP(w, r)
}
