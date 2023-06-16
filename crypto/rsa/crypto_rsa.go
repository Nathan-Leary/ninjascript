package rsa
import (
crypto_rsa "crypto/rsa"
)
func Load() {
if _, ok := Api["crypto/rsa"]; !ok {
   Api["crypto/rsa"] = map[string]interface{}{}
}
Api["crypto/rsa"]["PSSSaltLengthAuto"] = crypto_rsa.PSSSaltLengthAuto
Api["crypto/rsa"]["PSSSaltLengthEqualsHash"] = crypto_rsa.PSSSaltLengthEqualsHash
Api["crypto/rsa"]["DecryptOAEP"] = crypto_rsa.DecryptOAEP
Api["crypto/rsa"]["DecryptPKCS1v15"] = crypto_rsa.DecryptPKCS1v15
Api["crypto/rsa"]["DecryptPKCS1v15SessionKey"] = crypto_rsa.DecryptPKCS1v15SessionKey
Api["crypto/rsa"]["EncryptOAEP"] = crypto_rsa.EncryptOAEP
Api["crypto/rsa"]["EncryptPKCS1v15"] = crypto_rsa.EncryptPKCS1v15
Api["crypto/rsa"]["SignPKCS1v15"] = crypto_rsa.SignPKCS1v15
Api["crypto/rsa"]["SignPSS"] = crypto_rsa.SignPSS
Api["crypto/rsa"]["VerifyPKCS1v15"] = crypto_rsa.VerifyPKCS1v15
Api["crypto/rsa"]["VerifyPSS"] = crypto_rsa.VerifyPSS
Api["crypto/rsa"]["CRTValue"] = crypto_rsa.CRTValue{}
Api["crypto/rsa"]["OAEPOptions"] = crypto_rsa.OAEPOptions{}
Api["crypto/rsa"]["PKCS1v15DecryptOptions"] = crypto_rsa.PKCS1v15DecryptOptions{}
Api["crypto/rsa"]["PSSOptions"] = crypto_rsa.PSSOptions{}
Api["crypto/rsa"]["PrecomputedValues"] = crypto_rsa.PrecomputedValues{}
Api["crypto/rsa"]["PrivateKey"] = crypto_rsa.PrivateKey{}
Api["crypto/rsa"]["GenerateKey"] = crypto_rsa.GenerateKey
Api["crypto/rsa"]["GenerateMultiPrimeKey"] = crypto_rsa.GenerateMultiPrimeKey
Api["crypto/rsa"]["PublicKey"] = crypto_rsa.PublicKey{}
Api["crypto/rsa"]["DecryptOAEP"] = crypto_rsa.DecryptOAEP
Api["crypto/rsa"]["DecryptPKCS1v15"] = crypto_rsa.DecryptPKCS1v15
Api["crypto/rsa"]["DecryptPKCS1v15SessionKey"] = crypto_rsa.DecryptPKCS1v15SessionKey
Api["crypto/rsa"]["EncryptOAEP"] = crypto_rsa.EncryptOAEP
Api["crypto/rsa"]["EncryptPKCS1v15"] = crypto_rsa.EncryptPKCS1v15
Api["crypto/rsa"]["SignPKCS1v15"] = crypto_rsa.SignPKCS1v15
Api["crypto/rsa"]["SignPSS"] = crypto_rsa.SignPSS
Api["crypto/rsa"]["VerifyPKCS1v15"] = crypto_rsa.VerifyPKCS1v15
Api["crypto/rsa"]["VerifyPSS"] = crypto_rsa.VerifyPSS
Api["crypto/rsa"]["GenerateKey"] = crypto_rsa.GenerateKey
Api["crypto/rsa"]["GenerateMultiPrimeKey"] = crypto_rsa.GenerateMultiPrimeKey
}
