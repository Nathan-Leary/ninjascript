package ninjascript
import (
encoding_json "encoding/json"
)
func init() {if _, ok := Api["encoding/json"]; !ok {
   Api["encoding/json"] = map[string]interface{}{}
}
Api["encoding/json"]["Compact"] = encoding_json.Compact
Api["encoding/json"]["HTMLEscape"] = encoding_json.HTMLEscape
Api["encoding/json"]["Indent"] = encoding_json.Indent
Api["encoding/json"]["Marshal"] = encoding_json.Marshal
Api["encoding/json"]["MarshalIndent"] = encoding_json.MarshalIndent
Api["encoding/json"]["Unmarshal"] = encoding_json.Unmarshal
Api["encoding/json"]["Valid"] = encoding_json.Valid
Api["encoding/json"]["Decoder"] = encoding_json.Decoder{}
Api["encoding/json"]["NewDecoder"] = encoding_json.NewDecoder
Api["encoding/json"]["Encoder"] = encoding_json.Encoder{}
Api["encoding/json"]["NewEncoder"] = encoding_json.NewEncoder
Api["encoding/json"]["InvalidUTF8Error"] = encoding_json.InvalidUTF8Error{}
Api["encoding/json"]["InvalidUnmarshalError"] = encoding_json.InvalidUnmarshalError{}
Api["encoding/json"]["MarshalerError"] = encoding_json.MarshalerError{}
Api["encoding/json"]["SyntaxError"] = encoding_json.SyntaxError{}
Api["encoding/json"]["UnmarshalFieldError"] = encoding_json.UnmarshalFieldError{}
Api["encoding/json"]["UnmarshalTypeError"] = encoding_json.UnmarshalTypeError{}
Api["encoding/json"]["UnsupportedTypeError"] = encoding_json.UnsupportedTypeError{}
Api["encoding/json"]["UnsupportedValueError"] = encoding_json.UnsupportedValueError{}
Api["encoding/json"]["Compact"] = encoding_json.Compact
Api["encoding/json"]["HTMLEscape"] = encoding_json.HTMLEscape
Api["encoding/json"]["Indent"] = encoding_json.Indent
Api["encoding/json"]["Marshal"] = encoding_json.Marshal
Api["encoding/json"]["MarshalIndent"] = encoding_json.MarshalIndent
Api["encoding/json"]["Unmarshal"] = encoding_json.Unmarshal
Api["encoding/json"]["Valid"] = encoding_json.Valid
Api["encoding/json"]["NewDecoder"] = encoding_json.NewDecoder
Api["encoding/json"]["NewEncoder"] = encoding_json.NewEncoder

}