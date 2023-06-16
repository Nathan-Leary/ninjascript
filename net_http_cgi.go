package ninjascript
import (
net_http_cgi "net/http/cgi"
)
func init() {if _, ok := Api["net/http/cgi"]; !ok {
   Api["net/http/cgi"] = map[string]interface{}{}
}
Api["net/http/cgi"]["Request"] = net_http_cgi.Request
Api["net/http/cgi"]["RequestFromMap"] = net_http_cgi.RequestFromMap
Api["net/http/cgi"]["Serve"] = net_http_cgi.Serve
Api["net/http/cgi"]["Handler"] = net_http_cgi.Handler{}
Api["net/http/cgi"]["Request"] = net_http_cgi.Request
Api["net/http/cgi"]["RequestFromMap"] = net_http_cgi.RequestFromMap
Api["net/http/cgi"]["Serve"] = net_http_cgi.Serve

}