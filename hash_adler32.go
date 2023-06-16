package ninjascript
import (
hash_adler32 "hash/adler32"
)
func init() {if _, ok := Api["hash/adler32"]; !ok {
   Api["hash/adler32"] = map[string]interface{}{}
}
Api["hash/adler32"]["Checksum"] = hash_adler32.Checksum
Api["hash/adler32"]["New"] = hash_adler32.New
Api["hash/adler32"]["Checksum"] = hash_adler32.Checksum
Api["hash/adler32"]["New"] = hash_adler32.New

}