package studychanger

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"net/http"
	"strings"
)

func init() {
	caddy.RegisterModule(MiddlewareHeaderCheck{})
}

// header检查
type MiddlewareHeaderCheck struct {

}

// CaddyModule returns the Caddy module information.
func (MiddlewareHeaderCheck) CaddyModule() caddy.ModuleInfo {
	caddy.Log().Named("header_check").Info("init module header_check")
	return caddy.ModuleInfo{
		ID:  "http.handlers.header_check",
		New: func() caddy.Module { return new(MiddlewareHeaderCheck) },
	}
}

// 初始化处理
func (m MiddlewareHeaderCheck) Provision(ctx caddy.Context) error {


	return nil
}

// 检查header头
func (m MiddlewareHeaderCheck) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	val := r.Header.Values("User-Agent")
	if len(val) == 0 || !strings.Contains(val[0], "MSIE 11, Windows NT 6.3;"){
		//w.Write([]byte("illegal request"))
		w.WriteHeader(401)
		return nil
	}

	return next.ServeHTTP(w, r)
}


// 接口保证
var (
	_ caddyhttp.MiddlewareHandler = (*MiddlewareHeaderCheck)(nil)
)
