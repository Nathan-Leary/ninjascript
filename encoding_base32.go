package ninjascript
import (
encoding_base32 "encoding/base32"
)
func init() {if _, ok := Api["encoding/base32"]; !ok {
   Api["encoding/base32"] = map[string]interface{}{}
}
Api["encoding/base32"]["StdPadding"] = encoding_base32.StdPadding
Api["encoding/base32"]["NoPadding"] = encoding_base32.NoPadding
Api["encoding/base32"]["HexEncoding"] = encoding_base32.HexEncoding
Api["encoding/base32"]["StdEncoding"] = encoding_base32.StdEncoding
Api["encoding/base32"]["NewDecoder"] = encoding_base32.NewDecoder
Api["encoding/base32"]["NewEncoder"] = encoding_base32.NewEncoder
Api["encoding/base32"]["Encoding"] = encoding_base32.Encoding{}
Api["encoding/base32"]["NewEncoding"] = encoding_base32.NewEncoding
Api["encoding/base32"]["NewDecoder"] = encoding_base32.NewDecoder
Api["encoding/base32"]["NewEncoder"] = encoding_base32.NewEncoder
Api["encoding/base32"]["NewEncoding"] = encoding_base32.NewEncoding

}