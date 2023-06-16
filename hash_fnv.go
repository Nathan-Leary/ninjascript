package ninjascript
import (
hash_fnv "hash/fnv"
)
func init() {if _, ok := Api["hash/fnv"]; !ok {
   Api["hash/fnv"] = map[string]interface{}{}
}
Api["hash/fnv"]["New128"] = hash_fnv.New128
Api["hash/fnv"]["New128a"] = hash_fnv.New128a
Api["hash/fnv"]["New32"] = hash_fnv.New32
Api["hash/fnv"]["New32a"] = hash_fnv.New32a
Api["hash/fnv"]["New64"] = hash_fnv.New64
Api["hash/fnv"]["New64a"] = hash_fnv.New64a
Api["hash/fnv"]["New128"] = hash_fnv.New128
Api["hash/fnv"]["New128a"] = hash_fnv.New128a
Api["hash/fnv"]["New32"] = hash_fnv.New32
Api["hash/fnv"]["New32a"] = hash_fnv.New32a
Api["hash/fnv"]["New64"] = hash_fnv.New64
Api["hash/fnv"]["New64a"] = hash_fnv.New64a

}