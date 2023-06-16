package ninjascript
import (
mime_multipart "mime/multipart"
)
func init() {if _, ok := Api["mime/multipart"]; !ok {
   Api["mime/multipart"] = map[string]interface{}{}
}
Api["mime/multipart"]["ErrMessageTooLarge"] = mime_multipart.ErrMessageTooLarge
Api["mime/multipart"]["FileHeader"] = mime_multipart.FileHeader{}
Api["mime/multipart"]["Form"] = mime_multipart.Form{}
Api["mime/multipart"]["Part"] = mime_multipart.Part{}
Api["mime/multipart"]["Reader"] = mime_multipart.Reader{}
Api["mime/multipart"]["NewReader"] = mime_multipart.NewReader
Api["mime/multipart"]["Writer"] = mime_multipart.Writer{}
Api["mime/multipart"]["NewWriter"] = mime_multipart.NewWriter
Api["mime/multipart"]["NewReader"] = mime_multipart.NewReader
Api["mime/multipart"]["NewWriter"] = mime_multipart.NewWriter

}