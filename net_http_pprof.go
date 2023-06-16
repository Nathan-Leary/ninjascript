package ninjascript
import (
net_http_pprof "net/http/pprof"
)
func init() {if _, ok := Api["net/http/pprof"]; !ok {
   Api["net/http/pprof"] = map[string]interface{}{}
}
Api["net/http/pprof"]["Cmdline"] = net_http_pprof.Cmdline
Api["net/http/pprof"]["Handler"] = net_http_pprof.Handler
Api["net/http/pprof"]["Index"] = net_http_pprof.Index
Api["net/http/pprof"]["Profile"] = net_http_pprof.Profile
Api["net/http/pprof"]["Symbol"] = net_http_pprof.Symbol
Api["net/http/pprof"]["Trace"] = net_http_pprof.Trace
Api["net/http/pprof"]["Cmdline"] = net_http_pprof.Cmdline
Api["net/http/pprof"]["Handler"] = net_http_pprof.Handler
Api["net/http/pprof"]["Index"] = net_http_pprof.Index
Api["net/http/pprof"]["Profile"] = net_http_pprof.Profile
Api["net/http/pprof"]["Symbol"] = net_http_pprof.Symbol
Api["net/http/pprof"]["Trace"] = net_http_pprof.Trace

}