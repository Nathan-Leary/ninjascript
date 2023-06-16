package ninjascript
import (
image_jpeg "image/jpeg"
)
func init() {if _, ok := Api["image/jpeg"]; !ok {
   Api["image/jpeg"] = map[string]interface{}{}
}
Api["image/jpeg"]["Decode"] = image_jpeg.Decode
Api["image/jpeg"]["DecodeConfig"] = image_jpeg.DecodeConfig
Api["image/jpeg"]["Encode"] = image_jpeg.Encode
Api["image/jpeg"]["Options"] = image_jpeg.Options{}
Api["image/jpeg"]["Decode"] = image_jpeg.Decode
Api["image/jpeg"]["DecodeConfig"] = image_jpeg.DecodeConfig
Api["image/jpeg"]["Encode"] = image_jpeg.Encode

}