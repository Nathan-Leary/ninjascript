package aes
import (
crypto_aes "crypto/aes"
)
func Load() {
if _, ok := Api["crypto/aes"]; !ok {
   Api["crypto/aes"] = map[string]interface{}{}
}
Api["crypto/aes"]["NewCipher"] = crypto_aes.NewCipher
Api["crypto/aes"]["NewCipher"] = crypto_aes.NewCipher
}
