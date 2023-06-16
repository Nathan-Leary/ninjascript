package lzw
import (
compress_lzw "compress/lzw"
)
func Load() {
if _, ok := Api["compress/lzw"]; !ok {
   Api["compress/lzw"] = map[string]interface{}{}
}
Api["compress/lzw"]["NewReader"] = compress_lzw.NewReader
Api["compress/lzw"]["NewWriter"] = compress_lzw.NewWriter
Api["compress/lzw"]["LSB"] = compress_lzw.LSB
Api["compress/lzw"]["MSB"] = compress_lzw.MSB
Api["compress/lzw"]["Reader"] = compress_lzw.Reader{}
Api["compress/lzw"]["Writer"] = compress_lzw.Writer{}
Api["compress/lzw"]["NewReader"] = compress_lzw.NewReader
Api["compress/lzw"]["NewWriter"] = compress_lzw.NewWriter
}
