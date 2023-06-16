package ninjascript
import (
html "html"
)
func init() {if _, ok := Api["html"]; !ok {
   Api["html"] = map[string]interface{}{}
}
Api["html"]["EscapeString"] = html.EscapeString
Api["html"]["UnescapeString"] = html.UnescapeString
Api["html"]["EscapeString"] = html.EscapeString
Api["html"]["UnescapeString"] = html.UnescapeString

}