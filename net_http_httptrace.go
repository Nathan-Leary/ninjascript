package ninjascript
import (
net_http_httptrace "net/http/httptrace"
)
func init() {if _, ok := Api["net/http/httptrace"]; !ok {
   Api["net/http/httptrace"] = map[string]interface{}{}
}
Api["net/http/httptrace"]["WithClientTrace"] = net_http_httptrace.WithClientTrace
Api["net/http/httptrace"]["ClientTrace"] = net_http_httptrace.ClientTrace{}
Api["net/http/httptrace"]["ContextClientTrace"] = net_http_httptrace.ContextClientTrace
Api["net/http/httptrace"]["DNSDoneInfo"] = net_http_httptrace.DNSDoneInfo{}
Api["net/http/httptrace"]["DNSStartInfo"] = net_http_httptrace.DNSStartInfo{}
Api["net/http/httptrace"]["GotConnInfo"] = net_http_httptrace.GotConnInfo{}
Api["net/http/httptrace"]["WroteRequestInfo"] = net_http_httptrace.WroteRequestInfo{}
Api["net/http/httptrace"]["WithClientTrace"] = net_http_httptrace.WithClientTrace
Api["net/http/httptrace"]["ContextClientTrace"] = net_http_httptrace.ContextClientTrace

}