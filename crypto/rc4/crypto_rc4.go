package rc4
import (
crypto_rc4 "crypto/rc4"
)
func Load() {
if _, ok := Api["crypto/rc4"]; !ok {
   Api["crypto/rc4"] = map[string]interface{}{}
}
Api["crypto/rc4"]["Cipher"] = crypto_rc4.Cipher{}
Api["crypto/rc4"]["NewCipher"] = crypto_rc4.NewCipher
Api["crypto/rc4"]["NewCipher"] = crypto_rc4.NewCipher
}
