package crypto
import (
crypto "crypto"
)
func Load() {
if _, ok := Api["crypto"]; !ok {
   Api["crypto"] = map[string]interface{}{}
}
Api["crypto"]["RegisterHash"] = crypto.RegisterHash
Api["crypto"]["MD4"] = crypto.MD4
Api["crypto"]["MD5"] = crypto.MD5
Api["crypto"]["SHA1"] = crypto.SHA1
Api["crypto"]["SHA224"] = crypto.SHA224
Api["crypto"]["SHA256"] = crypto.SHA256
Api["crypto"]["SHA384"] = crypto.SHA384
Api["crypto"]["SHA512"] = crypto.SHA512
Api["crypto"]["MD5SHA1"] = crypto.MD5SHA1
Api["crypto"]["RIPEMD160"] = crypto.RIPEMD160
Api["crypto"]["SHA3_224"] = crypto.SHA3_224
Api["crypto"]["SHA3_256"] = crypto.SHA3_256
Api["crypto"]["SHA3_384"] = crypto.SHA3_384
Api["crypto"]["SHA3_512"] = crypto.SHA3_512
Api["crypto"]["SHA512_224"] = crypto.SHA512_224
Api["crypto"]["SHA512_256"] = crypto.SHA512_256
Api["crypto"]["BLAKE2s_256"] = crypto.BLAKE2s_256
Api["crypto"]["BLAKE2b_256"] = crypto.BLAKE2b_256
Api["crypto"]["BLAKE2b_384"] = crypto.BLAKE2b_384
Api["crypto"]["BLAKE2b_512"] = crypto.BLAKE2b_512
Api["crypto"]["Signer"] = new(crypto.Signer)
Api["crypto"]["RegisterHash"] = crypto.RegisterHash
}
