package ninjascript
import (
encoding_ascii85 "encoding/ascii85"
)
func init() {if _, ok := Api["encoding/ascii85"]; !ok {
   Api["encoding/ascii85"] = map[string]interface{}{}
}
Api["encoding/ascii85"]["Decode"] = encoding_ascii85.Decode
Api["encoding/ascii85"]["Encode"] = encoding_ascii85.Encode
Api["encoding/ascii85"]["MaxEncodedLen"] = encoding_ascii85.MaxEncodedLen
Api["encoding/ascii85"]["NewDecoder"] = encoding_ascii85.NewDecoder
Api["encoding/ascii85"]["NewEncoder"] = encoding_ascii85.NewEncoder
Api["encoding/ascii85"]["Decode"] = encoding_ascii85.Decode
Api["encoding/ascii85"]["Encode"] = encoding_ascii85.Encode
Api["encoding/ascii85"]["MaxEncodedLen"] = encoding_ascii85.MaxEncodedLen
Api["encoding/ascii85"]["NewDecoder"] = encoding_ascii85.NewDecoder
Api["encoding/ascii85"]["NewEncoder"] = encoding_ascii85.NewEncoder

}