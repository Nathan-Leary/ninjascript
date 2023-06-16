package ninjascript
import (
image_gif "image/gif"
)
func init() {if _, ok := Api["image/gif"]; !ok {
   Api["image/gif"] = map[string]interface{}{}
}
Api["image/gif"]["DisposalNone"] = image_gif.DisposalNone
Api["image/gif"]["DisposalBackground"] = image_gif.DisposalBackground
Api["image/gif"]["DisposalPrevious"] = image_gif.DisposalPrevious
Api["image/gif"]["Decode"] = image_gif.Decode
Api["image/gif"]["DecodeConfig"] = image_gif.DecodeConfig
Api["image/gif"]["Encode"] = image_gif.Encode
Api["image/gif"]["EncodeAll"] = image_gif.EncodeAll
Api["image/gif"]["GIF"] = image_gif.GIF{}
Api["image/gif"]["DecodeAll"] = image_gif.DecodeAll
Api["image/gif"]["Options"] = image_gif.Options{}
Api["image/gif"]["Decode"] = image_gif.Decode
Api["image/gif"]["DecodeConfig"] = image_gif.DecodeConfig
Api["image/gif"]["Encode"] = image_gif.Encode
Api["image/gif"]["EncodeAll"] = image_gif.EncodeAll
Api["image/gif"]["DecodeAll"] = image_gif.DecodeAll

}