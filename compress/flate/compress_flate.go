package flate
import (
compress_flate "compress/flate"
)
func Load() {
if _, ok := Api["compress/flate"]; !ok {
   Api["compress/flate"] = map[string]interface{}{}
}
Api["compress/flate"]["NoCompression"] = compress_flate.NoCompression
Api["compress/flate"]["BestSpeed"] = compress_flate.BestSpeed
Api["compress/flate"]["BestCompression"] = compress_flate.BestCompression
Api["compress/flate"]["DefaultCompression"] = compress_flate.DefaultCompression
Api["compress/flate"]["HuffmanOnly"] = compress_flate.HuffmanOnly
Api["compress/flate"]["NewReader"] = compress_flate.NewReader
Api["compress/flate"]["NewReaderDict"] = compress_flate.NewReaderDict
Api["compress/flate"]["ReadError"] = compress_flate.ReadError{}
Api["compress/flate"]["WriteError"] = compress_flate.WriteError{}
Api["compress/flate"]["Writer"] = compress_flate.Writer{}
Api["compress/flate"]["NewWriter"] = compress_flate.NewWriter
Api["compress/flate"]["NewWriterDict"] = compress_flate.NewWriterDict
Api["compress/flate"]["NewReader"] = compress_flate.NewReader
Api["compress/flate"]["NewReaderDict"] = compress_flate.NewReaderDict
Api["compress/flate"]["NewWriter"] = compress_flate.NewWriter
Api["compress/flate"]["NewWriterDict"] = compress_flate.NewWriterDict
}
