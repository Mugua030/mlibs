package mlibs

import (
	"mlibs/transport"
)

type Option func(a *App)

func Name(name string) Option {
	return func(a *App) { a.Name = name }
}
func Server(srv ...transport.Server) Option {
	return func(a *App) { a.servers = srv }
}
