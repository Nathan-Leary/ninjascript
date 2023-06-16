package ninjascript
import (
hash_maphash "hash/maphash"
)
func init() {if _, ok := Api["hash/maphash"]; !ok {
   Api["hash/maphash"] = map[string]interface{}{}
}
Api["hash/maphash"]["Bytes"] = hash_maphash.Bytes
Api["hash/maphash"]["String"] = hash_maphash.String
Api["hash/maphash"]["Hash"] = hash_maphash.Hash{}
Api["hash/maphash"]["Seed"] = hash_maphash.Seed{}
Api["hash/maphash"]["MakeSeed"] = hash_maphash.MakeSeed
Api["hash/maphash"]["Bytes"] = hash_maphash.Bytes
Api["hash/maphash"]["String"] = hash_maphash.String
Api["hash/maphash"]["MakeSeed"] = hash_maphash.MakeSeed

}