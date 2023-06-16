package dsa
import (
crypto_dsa "crypto/dsa"
)
func Load() {
if _, ok := Api["crypto/dsa"]; !ok {
   Api["crypto/dsa"] = map[string]interface{}{}
}
Api["crypto/dsa"]["GenerateKey"] = crypto_dsa.GenerateKey
Api["crypto/dsa"]["GenerateParameters"] = crypto_dsa.GenerateParameters
Api["crypto/dsa"]["Sign"] = crypto_dsa.Sign
Api["crypto/dsa"]["Verify"] = crypto_dsa.Verify
Api["crypto/dsa"]["L1024N160"] = crypto_dsa.L1024N160
Api["crypto/dsa"]["L2048N224"] = crypto_dsa.L2048N224
Api["crypto/dsa"]["L2048N256"] = crypto_dsa.L2048N256
Api["crypto/dsa"]["L3072N256"] = crypto_dsa.L3072N256
Api["crypto/dsa"]["Parameters"] = crypto_dsa.Parameters{}
Api["crypto/dsa"]["PrivateKey"] = crypto_dsa.PrivateKey{}
Api["crypto/dsa"]["PublicKey"] = crypto_dsa.PublicKey{}
Api["crypto/dsa"]["GenerateKey"] = crypto_dsa.GenerateKey
Api["crypto/dsa"]["GenerateParameters"] = crypto_dsa.GenerateParameters
Api["crypto/dsa"]["Sign"] = crypto_dsa.Sign
Api["crypto/dsa"]["Verify"] = crypto_dsa.Verify
}
