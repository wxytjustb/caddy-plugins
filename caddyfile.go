package studychanger

import (
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("read_cache", parseCaddyfile)
	httpcaddyfile.RegisterHandlerDirective("write_cache", parseCaddyfile)
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	m := new(MiddlewareCache)
	//var m Middleware
	//err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, nil
}