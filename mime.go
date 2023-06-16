package ninjascript
import (
mime "mime"
)
func init() {if _, ok := Api["mime"]; !ok {
   Api["mime"] = map[string]interface{}{}
}
Api["mime"]["BEncoding"] = mime.BEncoding
Api["mime"]["QEncoding"] = mime.QEncoding
Api["mime"]["ErrInvalidMediaParameter"] = mime.ErrInvalidMediaParameter
Api["mime"]["AddExtensionType"] = mime.AddExtensionType
Api["mime"]["ExtensionsByType"] = mime.ExtensionsByType
Api["mime"]["FormatMediaType"] = mime.FormatMediaType
Api["mime"]["ParseMediaType"] = mime.ParseMediaType
Api["mime"]["TypeByExtension"] = mime.TypeByExtension
Api["mime"]["WordDecoder"] = mime.WordDecoder{}
Api["mime"]["AddExtensionType"] = mime.AddExtensionType
Api["mime"]["ExtensionsByType"] = mime.ExtensionsByType
Api["mime"]["FormatMediaType"] = mime.FormatMediaType
Api["mime"]["ParseMediaType"] = mime.ParseMediaType
Api["mime"]["TypeByExtension"] = mime.TypeByExtension

}