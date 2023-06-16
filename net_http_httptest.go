package ninjascript
import (
net_http_httptest "net/http/httptest"
)
func init() {if _, ok := Api["net/http/httptest"]; !ok {
   Api["net/http/httptest"] = map[string]interface{}{}
}
Api["net/http/httptest"]["NewRequest"] = net_http_httptest.NewRequest
Api["net/http/httptest"]["ResponseRecorder"] = net_http_httptest.ResponseRecorder{}
Api["net/http/httptest"]["NewRecorder"] = net_http_httptest.NewRecorder
Api["net/http/httptest"]["Server"] = net_http_httptest.Server{}
Api["net/http/httptest"]["NewServer"] = net_http_httptest.NewServer
Api["net/http/httptest"]["NewTLSServer"] = net_http_httptest.NewTLSServer
Api["net/http/httptest"]["NewUnstartedServer"] = net_http_httptest.NewUnstartedServer
Api["net/http/httptest"]["NewRequest"] = net_http_httptest.NewRequest
Api["net/http/httptest"]["NewRecorder"] = net_http_httptest.NewRecorder
Api["net/http/httptest"]["NewServer"] = net_http_httptest.NewServer
Api["net/http/httptest"]["NewTLSServer"] = net_http_httptest.NewTLSServer
Api["net/http/httptest"]["NewUnstartedServer"] = net_http_httptest.NewUnstartedServer

}