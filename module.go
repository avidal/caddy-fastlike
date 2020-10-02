package caddyfastlike

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"

	"fastlike.dev"
)

func init() {
	caddy.RegisterModule(Module{})
}

type Module struct {
	proxy *fastlike.Fastlike
}

func (m Module) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "http.handlers.fastlike",
		New: func() caddy.Module {
			return Module{
				proxy: fastlike.New(""),
			}
		},
	}
}

func (m Module) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	return next.ServeHTTP(w, r)
}

var (
	_ caddyhttp.MiddlewareHandler = (*Module)(nil)
)
