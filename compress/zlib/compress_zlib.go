package zlib
import (
compress_zlib "compress/zlib"
)
func Load() {
if _, ok := Api["compress/zlib"]; !ok {
   Api["compress/zlib"] = map[string]interface{}{}
}
Api["compress/zlib"]["NoCompression"] = compress_zlib.NoCompression
Api["compress/zlib"]["BestSpeed"] = compress_zlib.BestSpeed
Api["compress/zlib"]["BestCompression"] = compress_zlib.BestCompression
Api["compress/zlib"]["DefaultCompression"] = compress_zlib.DefaultCompression
Api["compress/zlib"]["HuffmanOnly"] = compress_zlib.HuffmanOnly
Api["compress/zlib"]["NewReader"] = compress_zlib.NewReader
Api["compress/zlib"]["NewReaderDict"] = compress_zlib.NewReaderDict
Api["compress/zlib"]["Writer"] = compress_zlib.Writer{}
Api["compress/zlib"]["NewWriter"] = compress_zlib.NewWriter
Api["compress/zlib"]["NewWriterLevel"] = compress_zlib.NewWriterLevel
Api["compress/zlib"]["NewWriterLevelDict"] = compress_zlib.NewWriterLevelDict
Api["compress/zlib"]["NewReader"] = compress_zlib.NewReader
Api["compress/zlib"]["NewReaderDict"] = compress_zlib.NewReaderDict
Api["compress/zlib"]["NewWriter"] = compress_zlib.NewWriter
Api["compress/zlib"]["NewWriterLevel"] = compress_zlib.NewWriterLevel
Api["compress/zlib"]["NewWriterLevelDict"] = compress_zlib.NewWriterLevelDict
}
