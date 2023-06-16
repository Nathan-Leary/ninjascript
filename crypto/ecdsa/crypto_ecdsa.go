package ecdsa
import (
crypto_ecdsa "crypto/ecdsa"
)
func Load() {
if _, ok := Api["crypto/ecdsa"]; !ok {
   Api["crypto/ecdsa"] = map[string]interface{}{}
}
Api["crypto/ecdsa"]["Sign"] = crypto_ecdsa.Sign
Api["crypto/ecdsa"]["SignASN1"] = crypto_ecdsa.SignASN1
Api["crypto/ecdsa"]["Verify"] = crypto_ecdsa.Verify
Api["crypto/ecdsa"]["VerifyASN1"] = crypto_ecdsa.VerifyASN1
Api["crypto/ecdsa"]["PrivateKey"] = crypto_ecdsa.PrivateKey{}
Api["crypto/ecdsa"]["GenerateKey"] = crypto_ecdsa.GenerateKey
Api["crypto/ecdsa"]["PublicKey"] = crypto_ecdsa.PublicKey{}
Api["crypto/ecdsa"]["Sign"] = crypto_ecdsa.Sign
Api["crypto/ecdsa"]["SignASN1"] = crypto_ecdsa.SignASN1
Api["crypto/ecdsa"]["Verify"] = crypto_ecdsa.Verify
Api["crypto/ecdsa"]["VerifyASN1"] = crypto_ecdsa.VerifyASN1
Api["crypto/ecdsa"]["GenerateKey"] = crypto_ecdsa.GenerateKey
}
