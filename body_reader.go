package caddy_plugins

import (
	"bytes"
	"encoding/json"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"io/ioutil"
	"net/http"
)

func init() {
	caddy.RegisterModule(MiddlewareBodyReader{})
}

// header检查
type MiddlewareBodyReader struct {
	headerMapRaw	json.RawMessage		`json:"body_reader,omitempty" caddy:"namespace=http.handlers.body_reader"`
	headerMap 		map[string]string   `json:"-"`
}


func (MiddlewareBodyReader) CaddyModule() caddy.ModuleInfo {
	caddy.Log().Named("body_reader").Info("init module body_reader")
	return caddy.ModuleInfo{
		ID:  "http.handlers.body_reader",
		New: func() caddy.Module { return new(MiddlewareBodyReader) },
	}
}

// 初始化处理
func (m MiddlewareBodyReader) Provision(ctx caddy.Context) error {
	return nil
}

// 检查header头
func (m MiddlewareBodyReader) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	//var reader




	//io.Copy(reader, r.Body)
	data, _ := ioutil.ReadAll(r.Body)
	caddy.Log().Named("body_reader").Info(string(data))

	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	return next.ServeHTTP(w, r)
}


// 接口保证
var (
	_ caddyhttp.MiddlewareHandler = (*MiddlewareBodyReader)(nil)
)
