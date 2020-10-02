package caddyfastlike

import (
	"fmt"
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

	Wasm   string `json:"wasm"`
	Warmup int    `json:"warmup"`
}

func (m Module) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.fastlike",
		New: func() caddy.Module { return new(Module) },
	}
}

func (m *Module) Provision(ctx caddy.Context) error {
	proxy := fastlike.New(m.Wasm)
	m.proxy = proxy

	if m.Warmup > 0 {
		proxy.Warmup(m.Warmup)
	}

	return nil
}

func (m *Module) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	defer func() {
		if exc := recover(); exc != nil {
			fmt.Printf("RECOVER %q\n", exc)
		}
	}()
	var remoteaddr = r.RemoteAddr
	var requesturi = r.RequestURI
	var ctx = r.Context()

	inst := m.proxy.Instantiate(fastlike.BackendHandlerOption(func(_ string) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Fix these in fastlike
			r.RemoteAddr = remoteaddr
			r.RequestURI = requesturi

			// TODO: This is because the original request context doesn't carry, there's an issue
			// in the fastlike repo
			r = r.WithContext(ctx)

			// TODO: Figure out how to adapt and capture this error
			if err := next.ServeHTTP(w, r); err != nil {
				panic(err)
			}
		})
	}))

	inst.ServeHTTP(w, r)
	return nil
}

var (
	_ caddyhttp.MiddlewareHandler = (*Module)(nil)
	_ caddy.Provisioner           = (*Module)(nil)
)
