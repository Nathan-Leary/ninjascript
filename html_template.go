package ninjascript
import (
html_template "html/template"
)
func init() {if _, ok := Api["html/template"]; !ok {
   Api["html/template"] = map[string]interface{}{}
}
Api["html/template"]["HTMLEscape"] = html_template.HTMLEscape
Api["html/template"]["HTMLEscapeString"] = html_template.HTMLEscapeString
Api["html/template"]["HTMLEscaper"] = html_template.HTMLEscaper
Api["html/template"]["IsTrue"] = html_template.IsTrue
Api["html/template"]["JSEscape"] = html_template.JSEscape
Api["html/template"]["JSEscapeString"] = html_template.JSEscapeString
Api["html/template"]["JSEscaper"] = html_template.JSEscaper
Api["html/template"]["URLQueryEscaper"] = html_template.URLQueryEscaper
Api["html/template"]["Error"] = html_template.Error{}
Api["html/template"]["OK"] = html_template.OK
Api["html/template"]["ErrAmbigContext"] = html_template.ErrAmbigContext
Api["html/template"]["ErrBadHTML"] = html_template.ErrBadHTML
Api["html/template"]["ErrBranchEnd"] = html_template.ErrBranchEnd
Api["html/template"]["ErrEndContext"] = html_template.ErrEndContext
Api["html/template"]["ErrNoSuchTemplate"] = html_template.ErrNoSuchTemplate
Api["html/template"]["ErrOutputContext"] = html_template.ErrOutputContext
Api["html/template"]["ErrPartialCharset"] = html_template.ErrPartialCharset
Api["html/template"]["ErrPartialEscape"] = html_template.ErrPartialEscape
Api["html/template"]["ErrRangeLoopReentry"] = html_template.ErrRangeLoopReentry
Api["html/template"]["ErrSlashAmbig"] = html_template.ErrSlashAmbig
Api["html/template"]["ErrPredefinedEscaper"] = html_template.ErrPredefinedEscaper
Api["html/template"]["Template"] = html_template.Template{}
Api["html/template"]["Must"] = html_template.Must
Api["html/template"]["New"] = html_template.New
Api["html/template"]["ParseFS"] = html_template.ParseFS
Api["html/template"]["ParseFiles"] = html_template.ParseFiles
Api["html/template"]["ParseGlob"] = html_template.ParseGlob
Api["html/template"]["HTMLEscape"] = html_template.HTMLEscape
Api["html/template"]["HTMLEscapeString"] = html_template.HTMLEscapeString
Api["html/template"]["HTMLEscaper"] = html_template.HTMLEscaper
Api["html/template"]["IsTrue"] = html_template.IsTrue
Api["html/template"]["JSEscape"] = html_template.JSEscape
Api["html/template"]["JSEscapeString"] = html_template.JSEscapeString
Api["html/template"]["JSEscaper"] = html_template.JSEscaper
Api["html/template"]["URLQueryEscaper"] = html_template.URLQueryEscaper
Api["html/template"]["Must"] = html_template.Must
Api["html/template"]["New"] = html_template.New
Api["html/template"]["ParseFS"] = html_template.ParseFS
Api["html/template"]["ParseFiles"] = html_template.ParseFiles
Api["html/template"]["ParseGlob"] = html_template.ParseGlob

}