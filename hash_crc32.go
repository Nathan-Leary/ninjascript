package ninjascript
import (
hash_crc32 "hash/crc32"
)
func init() {if _, ok := Api["hash/crc32"]; !ok {
   Api["hash/crc32"] = map[string]interface{}{}
}
Api["hash/crc32"]["IEEE"] = hash_crc32.IEEE
Api["hash/crc32"]["Castagnoli"] = hash_crc32.Castagnoli
Api["hash/crc32"]["Koopman"] = hash_crc32.Koopman
Api["hash/crc32"]["IEEETable"] = hash_crc32.IEEETable
Api["hash/crc32"]["Checksum"] = hash_crc32.Checksum
Api["hash/crc32"]["ChecksumIEEE"] = hash_crc32.ChecksumIEEE
Api["hash/crc32"]["New"] = hash_crc32.New
Api["hash/crc32"]["NewIEEE"] = hash_crc32.NewIEEE
Api["hash/crc32"]["Update"] = hash_crc32.Update
Api["hash/crc32"]["MakeTable"] = hash_crc32.MakeTable
Api["hash/crc32"]["Checksum"] = hash_crc32.Checksum
Api["hash/crc32"]["ChecksumIEEE"] = hash_crc32.ChecksumIEEE
Api["hash/crc32"]["New"] = hash_crc32.New
Api["hash/crc32"]["NewIEEE"] = hash_crc32.NewIEEE
Api["hash/crc32"]["Update"] = hash_crc32.Update
Api["hash/crc32"]["MakeTable"] = hash_crc32.MakeTable

}