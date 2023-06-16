package ninjascript
import (
encoding_xml "encoding/xml"
)
func init() {if _, ok := Api["encoding/xml"]; !ok {
   Api["encoding/xml"] = map[string]interface{}{}
}
Api["encoding/xml"]["Header"] = encoding_xml.Header
Api["encoding/xml"]["HTMLAutoClose"] = encoding_xml.HTMLAutoClose
Api["encoding/xml"]["HTMLEntity"] = encoding_xml.HTMLEntity
Api["encoding/xml"]["Escape"] = encoding_xml.Escape
Api["encoding/xml"]["EscapeText"] = encoding_xml.EscapeText
Api["encoding/xml"]["Marshal"] = encoding_xml.Marshal
Api["encoding/xml"]["MarshalIndent"] = encoding_xml.MarshalIndent
Api["encoding/xml"]["Unmarshal"] = encoding_xml.Unmarshal
Api["encoding/xml"]["Attr"] = encoding_xml.Attr{}
Api["encoding/xml"]["NewDecoder"] = encoding_xml.NewDecoder
Api["encoding/xml"]["NewTokenDecoder"] = encoding_xml.NewTokenDecoder
Api["encoding/xml"]["Encoder"] = encoding_xml.Encoder{}
Api["encoding/xml"]["NewEncoder"] = encoding_xml.NewEncoder
Api["encoding/xml"]["EndElement"] = encoding_xml.EndElement{}
Api["encoding/xml"]["Name"] = encoding_xml.Name{}
Api["encoding/xml"]["ProcInst"] = encoding_xml.ProcInst{}
Api["encoding/xml"]["StartElement"] = encoding_xml.StartElement{}
Api["encoding/xml"]["SyntaxError"] = encoding_xml.SyntaxError{}
Api["encoding/xml"]["TagPathError"] = encoding_xml.TagPathError{}
Api["encoding/xml"]["CopyToken"] = encoding_xml.CopyToken
Api["encoding/xml"]["UnsupportedTypeError"] = encoding_xml.UnsupportedTypeError{}
Api["encoding/xml"]["Escape"] = encoding_xml.Escape
Api["encoding/xml"]["EscapeText"] = encoding_xml.EscapeText
Api["encoding/xml"]["Marshal"] = encoding_xml.Marshal
Api["encoding/xml"]["MarshalIndent"] = encoding_xml.MarshalIndent
Api["encoding/xml"]["Unmarshal"] = encoding_xml.Unmarshal
Api["encoding/xml"]["NewDecoder"] = encoding_xml.NewDecoder
Api["encoding/xml"]["NewTokenDecoder"] = encoding_xml.NewTokenDecoder
Api["encoding/xml"]["NewEncoder"] = encoding_xml.NewEncoder
Api["encoding/xml"]["CopyToken"] = encoding_xml.CopyToken

}