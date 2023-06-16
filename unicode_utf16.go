package ninjascript
import (
unicode_utf16 "unicode/utf16"
)
func init() {if _, ok := Api["unicode/utf16"]; !ok {
   Api["unicode/utf16"] = map[string]interface{}{}
}
Api["unicode/utf16"]["AppendRune"] = unicode_utf16.AppendRune
Api["unicode/utf16"]["Decode"] = unicode_utf16.Decode
Api["unicode/utf16"]["DecodeRune"] = unicode_utf16.DecodeRune
Api["unicode/utf16"]["Encode"] = unicode_utf16.Encode
Api["unicode/utf16"]["EncodeRune"] = unicode_utf16.EncodeRune
Api["unicode/utf16"]["IsSurrogate"] = unicode_utf16.IsSurrogate
Api["unicode/utf16"]["AppendRune"] = unicode_utf16.AppendRune
Api["unicode/utf16"]["Decode"] = unicode_utf16.Decode
Api["unicode/utf16"]["DecodeRune"] = unicode_utf16.DecodeRune
Api["unicode/utf16"]["Encode"] = unicode_utf16.Encode
Api["unicode/utf16"]["EncodeRune"] = unicode_utf16.EncodeRune
Api["unicode/utf16"]["IsSurrogate"] = unicode_utf16.IsSurrogate

}