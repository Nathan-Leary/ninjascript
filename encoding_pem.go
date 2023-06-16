package ninjascript
import (
encoding_pem "encoding/pem"
)
func init() {if _, ok := Api["encoding/pem"]; !ok {
   Api["encoding/pem"] = map[string]interface{}{}
}
Api["encoding/pem"]["Encode"] = encoding_pem.Encode
Api["encoding/pem"]["EncodeToMemory"] = encoding_pem.EncodeToMemory
Api["encoding/pem"]["Block"] = encoding_pem.Block{}
Api["encoding/pem"]["Decode"] = encoding_pem.Decode
Api["encoding/pem"]["Encode"] = encoding_pem.Encode
Api["encoding/pem"]["EncodeToMemory"] = encoding_pem.EncodeToMemory
Api["encoding/pem"]["Decode"] = encoding_pem.Decode

}