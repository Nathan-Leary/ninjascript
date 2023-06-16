package elliptic
import (
crypto_elliptic "crypto/elliptic"
)
func Load() {
if _, ok := Api["crypto/elliptic"]; !ok {
   Api["crypto/elliptic"] = map[string]interface{}{}
}
Api["crypto/elliptic"]["GenerateKey"] = crypto_elliptic.GenerateKey
Api["crypto/elliptic"]["Marshal"] = crypto_elliptic.Marshal
Api["crypto/elliptic"]["MarshalCompressed"] = crypto_elliptic.MarshalCompressed
Api["crypto/elliptic"]["Unmarshal"] = crypto_elliptic.Unmarshal
Api["crypto/elliptic"]["UnmarshalCompressed"] = crypto_elliptic.UnmarshalCompressed
Api["crypto/elliptic"]["P224"] = crypto_elliptic.P224
Api["crypto/elliptic"]["P256"] = crypto_elliptic.P256
Api["crypto/elliptic"]["P384"] = crypto_elliptic.P384
Api["crypto/elliptic"]["P521"] = crypto_elliptic.P521
Api["crypto/elliptic"]["CurveParams"] = crypto_elliptic.CurveParams{}
Api["crypto/elliptic"]["GenerateKey"] = crypto_elliptic.GenerateKey
Api["crypto/elliptic"]["Marshal"] = crypto_elliptic.Marshal
Api["crypto/elliptic"]["MarshalCompressed"] = crypto_elliptic.MarshalCompressed
Api["crypto/elliptic"]["Unmarshal"] = crypto_elliptic.Unmarshal
Api["crypto/elliptic"]["UnmarshalCompressed"] = crypto_elliptic.UnmarshalCompressed
Api["crypto/elliptic"]["P224"] = crypto_elliptic.P224
Api["crypto/elliptic"]["P256"] = crypto_elliptic.P256
Api["crypto/elliptic"]["P384"] = crypto_elliptic.P384
Api["crypto/elliptic"]["P521"] = crypto_elliptic.P521
}
