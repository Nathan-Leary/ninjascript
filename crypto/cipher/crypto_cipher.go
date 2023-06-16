package cipher
import (
crypto_cipher "crypto/cipher"
)
func Load() {
if _, ok := Api["crypto/cipher"]; !ok {
   Api["crypto/cipher"] = map[string]interface{}{}
}
Api["crypto/cipher"]["NewGCM"] = crypto_cipher.NewGCM
Api["crypto/cipher"]["NewGCMWithNonceSize"] = crypto_cipher.NewGCMWithNonceSize
Api["crypto/cipher"]["NewGCMWithTagSize"] = crypto_cipher.NewGCMWithTagSize
Api["crypto/cipher"]["NewCBCDecrypter"] = crypto_cipher.NewCBCDecrypter
Api["crypto/cipher"]["NewCBCEncrypter"] = crypto_cipher.NewCBCEncrypter
Api["crypto/cipher"]["NewCFBDecrypter"] = crypto_cipher.NewCFBDecrypter
Api["crypto/cipher"]["NewCFBEncrypter"] = crypto_cipher.NewCFBEncrypter
Api["crypto/cipher"]["NewCTR"] = crypto_cipher.NewCTR
Api["crypto/cipher"]["NewOFB"] = crypto_cipher.NewOFB
Api["crypto/cipher"]["StreamReader"] = crypto_cipher.StreamReader{}
Api["crypto/cipher"]["StreamWriter"] = crypto_cipher.StreamWriter{}
Api["crypto/cipher"]["NewGCM"] = crypto_cipher.NewGCM
Api["crypto/cipher"]["NewGCMWithNonceSize"] = crypto_cipher.NewGCMWithNonceSize
Api["crypto/cipher"]["NewGCMWithTagSize"] = crypto_cipher.NewGCMWithTagSize
Api["crypto/cipher"]["NewCBCDecrypter"] = crypto_cipher.NewCBCDecrypter
Api["crypto/cipher"]["NewCBCEncrypter"] = crypto_cipher.NewCBCEncrypter
Api["crypto/cipher"]["NewCFBDecrypter"] = crypto_cipher.NewCFBDecrypter
Api["crypto/cipher"]["NewCFBEncrypter"] = crypto_cipher.NewCFBEncrypter
Api["crypto/cipher"]["NewCTR"] = crypto_cipher.NewCTR
Api["crypto/cipher"]["NewOFB"] = crypto_cipher.NewOFB
}
