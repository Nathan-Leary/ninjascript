package des
import (
crypto_des "crypto/des"
)
func Load() {
if _, ok := Api["crypto/des"]; !ok {
   Api["crypto/des"] = map[string]interface{}{}
}
Api["crypto/des"]["NewCipher"] = crypto_des.NewCipher
Api["crypto/des"]["NewTripleDESCipher"] = crypto_des.NewTripleDESCipher
Api["crypto/des"]["NewCipher"] = crypto_des.NewCipher
Api["crypto/des"]["NewTripleDESCipher"] = crypto_des.NewTripleDESCipher
}
