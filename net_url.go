package ninjascript
import (
net_url "net/url"
)
func init() {if _, ok := Api["net/url"]; !ok {
   Api["net/url"] = map[string]interface{}{}
}
Api["net/url"]["JoinPath"] = net_url.JoinPath
Api["net/url"]["PathEscape"] = net_url.PathEscape
Api["net/url"]["PathUnescape"] = net_url.PathUnescape
Api["net/url"]["QueryEscape"] = net_url.QueryEscape
Api["net/url"]["QueryUnescape"] = net_url.QueryUnescape
Api["net/url"]["Error"] = net_url.Error{}
Api["net/url"]["URL"] = net_url.URL{}
Api["net/url"]["Parse"] = net_url.Parse
Api["net/url"]["ParseRequestURI"] = net_url.ParseRequestURI
Api["net/url"]["Userinfo"] = net_url.Userinfo{}
Api["net/url"]["User"] = net_url.User
Api["net/url"]["UserPassword"] = net_url.UserPassword
Api["net/url"]["ParseQuery"] = net_url.ParseQuery
Api["net/url"]["JoinPath"] = net_url.JoinPath
Api["net/url"]["PathEscape"] = net_url.PathEscape
Api["net/url"]["PathUnescape"] = net_url.PathUnescape
Api["net/url"]["QueryEscape"] = net_url.QueryEscape
Api["net/url"]["QueryUnescape"] = net_url.QueryUnescape
Api["net/url"]["Parse"] = net_url.Parse
Api["net/url"]["ParseRequestURI"] = net_url.ParseRequestURI
Api["net/url"]["User"] = net_url.User
Api["net/url"]["UserPassword"] = net_url.UserPassword
Api["net/url"]["ParseQuery"] = net_url.ParseQuery

}