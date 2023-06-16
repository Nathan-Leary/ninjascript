package md5
import (
crypto_md5 "crypto/md5"
)
func Load() {
if _, ok := Api["crypto/md5"]; !ok {
   Api["crypto/md5"] = map[string]interface{}{}
}
Api["crypto/md5"]["New"] = crypto_md5.New
Api["crypto/md5"]["Sum"] = crypto_md5.Sum
Api["crypto/md5"]["New"] = crypto_md5.New
Api["crypto/md5"]["Sum"] = crypto_md5.Sum
}
