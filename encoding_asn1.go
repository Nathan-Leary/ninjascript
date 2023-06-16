package ninjascript
import (
encoding_asn1 "encoding/asn1"
)
func init() {if _, ok := Api["encoding/asn1"]; !ok {
   Api["encoding/asn1"] = map[string]interface{}{}
}
Api["encoding/asn1"]["TagBoolean"] = encoding_asn1.TagBoolean
Api["encoding/asn1"]["TagInteger"] = encoding_asn1.TagInteger
Api["encoding/asn1"]["TagBitString"] = encoding_asn1.TagBitString
Api["encoding/asn1"]["TagOctetString"] = encoding_asn1.TagOctetString
Api["encoding/asn1"]["TagNull"] = encoding_asn1.TagNull
Api["encoding/asn1"]["TagOID"] = encoding_asn1.TagOID
Api["encoding/asn1"]["TagEnum"] = encoding_asn1.TagEnum
Api["encoding/asn1"]["TagUTF8String"] = encoding_asn1.TagUTF8String
Api["encoding/asn1"]["TagSequence"] = encoding_asn1.TagSequence
Api["encoding/asn1"]["TagSet"] = encoding_asn1.TagSet
Api["encoding/asn1"]["TagNumericString"] = encoding_asn1.TagNumericString
Api["encoding/asn1"]["TagPrintableString"] = encoding_asn1.TagPrintableString
Api["encoding/asn1"]["TagT61String"] = encoding_asn1.TagT61String
Api["encoding/asn1"]["TagIA5String"] = encoding_asn1.TagIA5String
Api["encoding/asn1"]["TagUTCTime"] = encoding_asn1.TagUTCTime
Api["encoding/asn1"]["TagGeneralizedTime"] = encoding_asn1.TagGeneralizedTime
Api["encoding/asn1"]["TagGeneralString"] = encoding_asn1.TagGeneralString
Api["encoding/asn1"]["TagBMPString"] = encoding_asn1.TagBMPString
Api["encoding/asn1"]["ClassUniversal"] = encoding_asn1.ClassUniversal
Api["encoding/asn1"]["ClassApplication"] = encoding_asn1.ClassApplication
Api["encoding/asn1"]["ClassContextSpecific"] = encoding_asn1.ClassContextSpecific
Api["encoding/asn1"]["ClassPrivate"] = encoding_asn1.ClassPrivate
Api["encoding/asn1"]["NullBytes"] = encoding_asn1.NullBytes
Api["encoding/asn1"]["NullRawValue"] = encoding_asn1.NullRawValue
Api["encoding/asn1"]["Marshal"] = encoding_asn1.Marshal
Api["encoding/asn1"]["MarshalWithParams"] = encoding_asn1.MarshalWithParams
Api["encoding/asn1"]["Unmarshal"] = encoding_asn1.Unmarshal
Api["encoding/asn1"]["UnmarshalWithParams"] = encoding_asn1.UnmarshalWithParams
Api["encoding/asn1"]["BitString"] = encoding_asn1.BitString{}
Api["encoding/asn1"]["RawValue"] = encoding_asn1.RawValue{}
Api["encoding/asn1"]["StructuralError"] = encoding_asn1.StructuralError{}
Api["encoding/asn1"]["SyntaxError"] = encoding_asn1.SyntaxError{}
Api["encoding/asn1"]["Marshal"] = encoding_asn1.Marshal
Api["encoding/asn1"]["MarshalWithParams"] = encoding_asn1.MarshalWithParams
Api["encoding/asn1"]["Unmarshal"] = encoding_asn1.Unmarshal
Api["encoding/asn1"]["UnmarshalWithParams"] = encoding_asn1.UnmarshalWithParams

}