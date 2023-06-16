package ninjascript
import (
encoding_base64 "encoding/base64"
)
func init() {if _, ok := Api["encoding/base64"]; !ok {
   Api["encoding/base64"] = map[string]interface{}{}
}
Api["encoding/base64"]["StdPadding"] = encoding_base64.StdPadding
Api["encoding/base64"]["NoPadding"] = encoding_base64.NoPadding
Api["encoding/base64"]["RawStdEncoding"] = encoding_base64.RawStdEncoding
Api["encoding/base64"]["RawURLEncoding"] = encoding_base64.RawURLEncoding
Api["encoding/base64"]["StdEncoding"] = encoding_base64.StdEncoding
Api["encoding/base64"]["URLEncoding"] = encoding_base64.URLEncoding
Api["encoding/base64"]["NewDecoder"] = encoding_base64.NewDecoder
Api["encoding/base64"]["NewEncoder"] = encoding_base64.NewEncoder
Api["encoding/base64"]["Encoding"] = encoding_base64.Encoding{}
Api["encoding/base64"]["NewEncoding"] = encoding_base64.NewEncoding
Api["encoding/base64"]["NewDecoder"] = encoding_base64.NewDecoder
Api["encoding/base64"]["NewEncoder"] = encoding_base64.NewEncoder
Api["encoding/base64"]["NewEncoding"] = encoding_base64.NewEncoding

}