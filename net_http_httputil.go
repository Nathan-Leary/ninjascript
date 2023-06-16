package ninjascript
import (
net_http_httputil "net/http/httputil"
)
func init() {if _, ok := Api["net/http/httputil"]; !ok {
   Api["net/http/httputil"] = map[string]interface{}{}
}
Api["net/http/httputil"]["ErrLineTooLong"] = net_http_httputil.ErrLineTooLong
Api["net/http/httputil"]["DumpRequest"] = net_http_httputil.DumpRequest
Api["net/http/httputil"]["DumpRequestOut"] = net_http_httputil.DumpRequestOut
Api["net/http/httputil"]["DumpResponse"] = net_http_httputil.DumpResponse
Api["net/http/httputil"]["NewChunkedReader"] = net_http_httputil.NewChunkedReader
Api["net/http/httputil"]["NewChunkedWriter"] = net_http_httputil.NewChunkedWriter
Api["net/http/httputil"]["ClientConn"] = net_http_httputil.ClientConn{}
Api["net/http/httputil"]["NewClientConn"] = net_http_httputil.NewClientConn
Api["net/http/httputil"]["NewProxyClientConn"] = net_http_httputil.NewProxyClientConn
Api["net/http/httputil"]["ProxyRequest"] = net_http_httputil.ProxyRequest{}
Api["net/http/httputil"]["ReverseProxy"] = net_http_httputil.ReverseProxy{}
Api["net/http/httputil"]["NewSingleHostReverseProxy"] = net_http_httputil.NewSingleHostReverseProxy
Api["net/http/httputil"]["ServerConn"] = net_http_httputil.ServerConn{}
Api["net/http/httputil"]["NewServerConn"] = net_http_httputil.NewServerConn
Api["net/http/httputil"]["DumpRequest"] = net_http_httputil.DumpRequest
Api["net/http/httputil"]["DumpRequestOut"] = net_http_httputil.DumpRequestOut
Api["net/http/httputil"]["DumpResponse"] = net_http_httputil.DumpResponse
Api["net/http/httputil"]["NewChunkedReader"] = net_http_httputil.NewChunkedReader
Api["net/http/httputil"]["NewChunkedWriter"] = net_http_httputil.NewChunkedWriter
Api["net/http/httputil"]["NewClientConn"] = net_http_httputil.NewClientConn
Api["net/http/httputil"]["NewProxyClientConn"] = net_http_httputil.NewProxyClientConn
Api["net/http/httputil"]["NewSingleHostReverseProxy"] = net_http_httputil.NewSingleHostReverseProxy
Api["net/http/httputil"]["NewServerConn"] = net_http_httputil.NewServerConn

}