package caddy_plugins

import (
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("body_reader", checkParseCaddyfile)
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func checkParseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	m := MiddlewareBodyReader{
		headerMap: make(map[string]string),
	}

	for h.Next() {
		var segments []caddyfile.Segment
		for nesting := h.Nesting(); h.NextBlock(nesting); {
			segment := h.NextSegment()
			m.headerMap[segment[0].Text] = segment[1].Text
			segments = append(segments, segment)
		}
	}

	return m, nil
}