package ninjascript
import (
encoding_hex "encoding/hex"
)
func init() {if _, ok := Api["encoding/hex"]; !ok {
   Api["encoding/hex"] = map[string]interface{}{}
}
Api["encoding/hex"]["Decode"] = encoding_hex.Decode
Api["encoding/hex"]["DecodeString"] = encoding_hex.DecodeString
Api["encoding/hex"]["DecodedLen"] = encoding_hex.DecodedLen
Api["encoding/hex"]["Dump"] = encoding_hex.Dump
Api["encoding/hex"]["Dumper"] = encoding_hex.Dumper
Api["encoding/hex"]["Encode"] = encoding_hex.Encode
Api["encoding/hex"]["EncodeToString"] = encoding_hex.EncodeToString
Api["encoding/hex"]["EncodedLen"] = encoding_hex.EncodedLen
Api["encoding/hex"]["NewDecoder"] = encoding_hex.NewDecoder
Api["encoding/hex"]["NewEncoder"] = encoding_hex.NewEncoder
Api["encoding/hex"]["Decode"] = encoding_hex.Decode
Api["encoding/hex"]["DecodeString"] = encoding_hex.DecodeString
Api["encoding/hex"]["DecodedLen"] = encoding_hex.DecodedLen
Api["encoding/hex"]["Dump"] = encoding_hex.Dump
Api["encoding/hex"]["Dumper"] = encoding_hex.Dumper
Api["encoding/hex"]["Encode"] = encoding_hex.Encode
Api["encoding/hex"]["EncodeToString"] = encoding_hex.EncodeToString
Api["encoding/hex"]["EncodedLen"] = encoding_hex.EncodedLen
Api["encoding/hex"]["NewDecoder"] = encoding_hex.NewDecoder
Api["encoding/hex"]["NewEncoder"] = encoding_hex.NewEncoder

}