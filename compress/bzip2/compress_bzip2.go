package bzip2
import (
compress_bzip2 "compress/bzip2"
)
func Load() {
if _, ok := Api["compress/bzip2"]; !ok {
   Api["compress/bzip2"] = map[string]interface{}{}
}
Api["compress/bzip2"]["NewReader"] = compress_bzip2.NewReader
Api["compress/bzip2"]["NewReader"] = compress_bzip2.NewReader
}
