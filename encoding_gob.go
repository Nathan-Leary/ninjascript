package ninjascript
import (
encoding_gob "encoding/gob"
)
func init() {if _, ok := Api["encoding/gob"]; !ok {
   Api["encoding/gob"] = map[string]interface{}{}
}
Api["encoding/gob"]["Register"] = encoding_gob.Register
Api["encoding/gob"]["RegisterName"] = encoding_gob.RegisterName
Api["encoding/gob"]["CommonType"] = encoding_gob.CommonType{}
Api["encoding/gob"]["Decoder"] = encoding_gob.Decoder{}
Api["encoding/gob"]["NewDecoder"] = encoding_gob.NewDecoder
Api["encoding/gob"]["Encoder"] = encoding_gob.Encoder{}
Api["encoding/gob"]["NewEncoder"] = encoding_gob.NewEncoder
Api["encoding/gob"]["Register"] = encoding_gob.Register
Api["encoding/gob"]["RegisterName"] = encoding_gob.RegisterName
Api["encoding/gob"]["NewDecoder"] = encoding_gob.NewDecoder
Api["encoding/gob"]["NewEncoder"] = encoding_gob.NewEncoder

}