package rand
import (
crypto_rand "crypto/rand"
)
func Load() {
if _, ok := Api["crypto/rand"]; !ok {
   Api["crypto/rand"] = map[string]interface{}{}
}
Api["crypto/rand"]["Reader"] = crypto_rand.Reader
Api["crypto/rand"]["Int"] = crypto_rand.Int
Api["crypto/rand"]["Prime"] = crypto_rand.Prime
Api["crypto/rand"]["Read"] = crypto_rand.Read
Api["crypto/rand"]["Int"] = crypto_rand.Int
Api["crypto/rand"]["Prime"] = crypto_rand.Prime
Api["crypto/rand"]["Read"] = crypto_rand.Read
}
