package ninjascript

import (
	hash_crc64 "hash/crc64"
)

func init() {
	if _, ok := Api["hash/crc64"]; !ok {
		Api["hash/crc64"] = map[string]interface{}{}
	}
	// Api["hash/crc64"]["ISO"] = hash_crc64.ISO
	// Api["hash/crc64"]["ECMA"] = hash_crc64.ECMA
	Api["hash/crc64"]["Checksum"] = hash_crc64.Checksum
	Api["hash/crc64"]["New"] = hash_crc64.New
	Api["hash/crc64"]["Update"] = hash_crc64.Update
	Api["hash/crc64"]["MakeTable"] = hash_crc64.MakeTable
	Api["hash/crc64"]["Checksum"] = hash_crc64.Checksum
	Api["hash/crc64"]["New"] = hash_crc64.New
	Api["hash/crc64"]["Update"] = hash_crc64.Update
	Api["hash/crc64"]["MakeTable"] = hash_crc64.MakeTable

}
