package ninjascript
import (
net_http_fcgi "net/http/fcgi"
)
func init() {if _, ok := Api["net/http/fcgi"]; !ok {
   Api["net/http/fcgi"] = map[string]interface{}{}
}
Api["net/http/fcgi"]["ErrConnClosed"] = net_http_fcgi.ErrConnClosed
Api["net/http/fcgi"]["ErrRequestAborted"] = net_http_fcgi.ErrRequestAborted
Api["net/http/fcgi"]["ProcessEnv"] = net_http_fcgi.ProcessEnv
Api["net/http/fcgi"]["Serve"] = net_http_fcgi.Serve
Api["net/http/fcgi"]["ProcessEnv"] = net_http_fcgi.ProcessEnv
Api["net/http/fcgi"]["Serve"] = net_http_fcgi.Serve

}