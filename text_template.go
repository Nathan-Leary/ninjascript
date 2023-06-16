package ninjascript
import (
text_template "text/template"
)
func init() {if _, ok := Api["text/template"]; !ok {
   Api["text/template"] = map[string]interface{}{}
}
Api["text/template"]["HTMLEscape"] = text_template.HTMLEscape
Api["text/template"]["HTMLEscapeString"] = text_template.HTMLEscapeString
Api["text/template"]["HTMLEscaper"] = text_template.HTMLEscaper
Api["text/template"]["IsTrue"] = text_template.IsTrue
Api["text/template"]["JSEscape"] = text_template.JSEscape
Api["text/template"]["JSEscapeString"] = text_template.JSEscapeString
Api["text/template"]["JSEscaper"] = text_template.JSEscaper
Api["text/template"]["URLQueryEscaper"] = text_template.URLQueryEscaper
Api["text/template"]["ExecError"] = text_template.ExecError{}
Api["text/template"]["Template"] = text_template.Template{}
Api["text/template"]["Must"] = text_template.Must
Api["text/template"]["New"] = text_template.New
Api["text/template"]["ParseFS"] = text_template.ParseFS
Api["text/template"]["ParseFiles"] = text_template.ParseFiles
Api["text/template"]["ParseGlob"] = text_template.ParseGlob
Api["text/template"]["HTMLEscape"] = text_template.HTMLEscape
Api["text/template"]["HTMLEscapeString"] = text_template.HTMLEscapeString
Api["text/template"]["HTMLEscaper"] = text_template.HTMLEscaper
Api["text/template"]["IsTrue"] = text_template.IsTrue
Api["text/template"]["JSEscape"] = text_template.JSEscape
Api["text/template"]["JSEscapeString"] = text_template.JSEscapeString
Api["text/template"]["JSEscaper"] = text_template.JSEscaper
Api["text/template"]["URLQueryEscaper"] = text_template.URLQueryEscaper
Api["text/template"]["Must"] = text_template.Must
Api["text/template"]["New"] = text_template.New
Api["text/template"]["ParseFS"] = text_template.ParseFS
Api["text/template"]["ParseFiles"] = text_template.ParseFiles
Api["text/template"]["ParseGlob"] = text_template.ParseGlob

}