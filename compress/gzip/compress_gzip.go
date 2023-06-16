package gzip
import (
compress_gzip "compress/gzip"
)
func Load() {
if _, ok := Api["compress/gzip"]; !ok {
   Api["compress/gzip"] = map[string]interface{}{}
}
Api["compress/gzip"]["NoCompression"] = compress_gzip.NoCompression
Api["compress/gzip"]["BestSpeed"] = compress_gzip.BestSpeed
Api["compress/gzip"]["BestCompression"] = compress_gzip.BestCompression
Api["compress/gzip"]["DefaultCompression"] = compress_gzip.DefaultCompression
Api["compress/gzip"]["HuffmanOnly"] = compress_gzip.HuffmanOnly
Api["compress/gzip"]["Header"] = compress_gzip.Header{}
Api["compress/gzip"]["Reader"] = compress_gzip.Reader{}
Api["compress/gzip"]["NewReader"] = compress_gzip.NewReader
Api["compress/gzip"]["Writer"] = compress_gzip.Writer{}
Api["compress/gzip"]["NewWriter"] = compress_gzip.NewWriter
Api["compress/gzip"]["NewWriterLevel"] = compress_gzip.NewWriterLevel
Api["compress/gzip"]["NewReader"] = compress_gzip.NewReader
Api["compress/gzip"]["NewWriter"] = compress_gzip.NewWriter
Api["compress/gzip"]["NewWriterLevel"] = compress_gzip.NewWriterLevel
}
