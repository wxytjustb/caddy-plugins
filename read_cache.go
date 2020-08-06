package caddy_plugins

import (
	"fmt"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func init() {
	caddy.RegisterModule(MiddlewareCache{})
}

// Gizmo is an example; put your own type here.
type MiddlewareCache struct {
}

// CaddyModule returns the Caddy module information.
func (MiddlewareCache) CaddyModule() caddy.ModuleInfo {
	caddy.Log().Named("middleware_cache").Info("init module middleware_read_cache")
	return caddy.ModuleInfo{
		ID:  "http.handlers.read_cache",
		New: func() caddy.Module { return new(MiddlewareCache) },
	}
}

func (m MiddlewareCache) Provision(ctx caddy.Context) error {

	go func() {
		for {
			select {
			case <-ctx.Done():
				caddy.Log().Named("cache").Info("缓存容量检测关闭")
				return
			case <-time.After(30 * time.Second):
				caddy.Log().Named("cache").Info(fmt.Sprintf("缓存容量:%d", CacheObj.GetLen()))
			}
		}
	}()

	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m MiddlewareCache) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {


	if r.RequestURI != "" && strings.HasSuffix(r.RequestURI, ".css") {
		if val, ok := CacheObj.GetData(r.RequestURI); ok {
			w.Header().Set("content-type", "text/css")

			// todo 研究一下为什么是这个数值
			bufLen := 32 * 1024
			for start := 0; start < len(val); start += bufLen {

				if start+bufLen >= len(val) {
					w.Write(val[start:])
				} else {
					w.Write(val[start : start+bufLen])
				}
			}
			return nil
		}
	}

	w = NewResponseWriterWithCache(w, r)

	return next.ServeHTTP(w, r)
}

type ResponseWriterWithCache struct {
	w    http.ResponseWriter
	r    *http.Request
	uuid string
}

func NewResponseWriterWithCache(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	return &ResponseWriterWithCache{
		w:    w,
		r:    r,
		uuid: uuid.New().String(),
	}
}

func (cacheWriter *ResponseWriterWithCache) Header() http.Header {
	return cacheWriter.w.Header()
}

func (cacheWriter *ResponseWriterWithCache) Write(data []byte) (int, error) {
	if cacheWriter.r.RequestURI != "" && strings.HasSuffix(cacheWriter.r.RequestURI, ".css") {
		CacheObj.SetData(cacheWriter.r.RequestURI, cacheWriter.uuid, data)
	}

	return cacheWriter.w.Write(data)
}

func (cacheWriter *ResponseWriterWithCache) WriteHeader(statusCode int) {
	cacheWriter.w.WriteHeader(statusCode)
}

// 接口保证
var (
	_ http.ResponseWriter         = (*ResponseWriterWithCache)(nil)
	_ caddyhttp.MiddlewareHandler = (*MiddlewareCache)(nil)
)
