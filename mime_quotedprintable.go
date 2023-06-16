package ninjascript
import (
mime_quotedprintable "mime/quotedprintable"
)
func init() {if _, ok := Api["mime/quotedprintable"]; !ok {
   Api["mime/quotedprintable"] = map[string]interface{}{}
}
Api["mime/quotedprintable"]["Reader"] = mime_quotedprintable.Reader{}
Api["mime/quotedprintable"]["NewReader"] = mime_quotedprintable.NewReader
Api["mime/quotedprintable"]["Writer"] = mime_quotedprintable.Writer{}
Api["mime/quotedprintable"]["NewWriter"] = mime_quotedprintable.NewWriter
Api["mime/quotedprintable"]["NewReader"] = mime_quotedprintable.NewReader
Api["mime/quotedprintable"]["NewWriter"] = mime_quotedprintable.NewWriter

}