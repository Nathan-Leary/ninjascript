package hmac
import (
crypto_hmac "crypto/hmac"
)
func Load() {
if _, ok := Api["crypto/hmac"]; !ok {
   Api["crypto/hmac"] = map[string]interface{}{}
}
Api["crypto/hmac"]["Equal"] = crypto_hmac.Equal
Api["crypto/hmac"]["New"] = crypto_hmac.New
Api["crypto/hmac"]["Equal"] = crypto_hmac.Equal
Api["crypto/hmac"]["New"] = crypto_hmac.New
}
