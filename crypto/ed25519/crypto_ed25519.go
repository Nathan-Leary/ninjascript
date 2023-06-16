package ed25519
import (
crypto_ed25519 "crypto/ed25519"
)
func Load() {
if _, ok := Api["crypto/ed25519"]; !ok {
   Api["crypto/ed25519"] = map[string]interface{}{}
}
Api["crypto/ed25519"]["PublicKeySize"] = crypto_ed25519.PublicKeySize
Api["crypto/ed25519"]["PrivateKeySize"] = crypto_ed25519.PrivateKeySize
Api["crypto/ed25519"]["SignatureSize"] = crypto_ed25519.SignatureSize
Api["crypto/ed25519"]["SeedSize"] = crypto_ed25519.SeedSize
Api["crypto/ed25519"]["GenerateKey"] = crypto_ed25519.GenerateKey
Api["crypto/ed25519"]["Sign"] = crypto_ed25519.Sign
Api["crypto/ed25519"]["Verify"] = crypto_ed25519.Verify
Api["crypto/ed25519"]["VerifyWithOptions"] = crypto_ed25519.VerifyWithOptions
Api["crypto/ed25519"]["Options"] = crypto_ed25519.Options{}
Api["crypto/ed25519"]["NewKeyFromSeed"] = crypto_ed25519.NewKeyFromSeed
Api["crypto/ed25519"]["GenerateKey"] = crypto_ed25519.GenerateKey
Api["crypto/ed25519"]["Sign"] = crypto_ed25519.Sign
Api["crypto/ed25519"]["Verify"] = crypto_ed25519.Verify
Api["crypto/ed25519"]["VerifyWithOptions"] = crypto_ed25519.VerifyWithOptions
Api["crypto/ed25519"]["NewKeyFromSeed"] = crypto_ed25519.NewKeyFromSeed
}
