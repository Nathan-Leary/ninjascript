package sha1
import (
crypto_sha1 "crypto/sha1"
)
func Load() {
if _, ok := Api["crypto/sha1"]; !ok {
   Api["crypto/sha1"] = map[string]interface{}{}
}
Api["crypto/sha1"]["New"] = crypto_sha1.New
Api["crypto/sha1"]["Sum"] = crypto_sha1.Sum
Api["crypto/sha1"]["New"] = crypto_sha1.New
Api["crypto/sha1"]["Sum"] = crypto_sha1.Sum
}
