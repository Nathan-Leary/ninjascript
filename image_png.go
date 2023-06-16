package ninjascript
import (
image_png "image/png"
)
func init() {if _, ok := Api["image/png"]; !ok {
   Api["image/png"] = map[string]interface{}{}
}
Api["image/png"]["Decode"] = image_png.Decode
Api["image/png"]["DecodeConfig"] = image_png.DecodeConfig
Api["image/png"]["Encode"] = image_png.Encode
Api["image/png"]["DefaultCompression"] = image_png.DefaultCompression
Api["image/png"]["NoCompression"] = image_png.NoCompression
Api["image/png"]["BestSpeed"] = image_png.BestSpeed
Api["image/png"]["BestCompression"] = image_png.BestCompression
Api["image/png"]["Encoder"] = image_png.Encoder{}
Api["image/png"]["Decode"] = image_png.Decode
Api["image/png"]["DecodeConfig"] = image_png.DecodeConfig
Api["image/png"]["Encode"] = image_png.Encode

}