package caddy_plugins

import (
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("header_check", headerCheckParseCaddyfile)
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func headerCheckParseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	m := new(MiddlewareHeaderCheck)
	//var m Middleware
	//err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, nil
}