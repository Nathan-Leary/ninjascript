package ninjascript
import (
net_http_cookiejar "net/http/cookiejar"
)
func init() {if _, ok := Api["net/http/cookiejar"]; !ok {
   Api["net/http/cookiejar"] = map[string]interface{}{}
}
Api["net/http/cookiejar"]["Jar"] = net_http_cookiejar.Jar{}
Api["net/http/cookiejar"]["New"] = net_http_cookiejar.New
Api["net/http/cookiejar"]["Options"] = net_http_cookiejar.Options{}
Api["net/http/cookiejar"]["New"] = net_http_cookiejar.New

}