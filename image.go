package ninjascript
import (
image "image"
)
func init() {if _, ok := Api["image"]; !ok {
   Api["image"] = map[string]interface{}{}
}
Api["image"]["ErrFormat"] = image.ErrFormat
Api["image"]["RegisterFormat"] = image.RegisterFormat
Api["image"]["Alpha"] = image.Alpha{}
Api["image"]["NewAlpha"] = image.NewAlpha
Api["image"]["Alpha16"] = image.Alpha16{}
Api["image"]["NewAlpha16"] = image.NewAlpha16
Api["image"]["CMYK"] = image.CMYK{}
Api["image"]["NewCMYK"] = image.NewCMYK
Api["image"]["Config"] = image.Config{}
Api["image"]["DecodeConfig"] = image.DecodeConfig
Api["image"]["Gray"] = image.Gray{}
Api["image"]["NewGray"] = image.NewGray
Api["image"]["Gray16"] = image.Gray16{}
Api["image"]["NewGray16"] = image.NewGray16
Api["image"]["Decode"] = image.Decode
Api["image"]["NRGBA"] = image.NRGBA{}
Api["image"]["NewNRGBA"] = image.NewNRGBA
Api["image"]["NRGBA64"] = image.NRGBA64{}
Api["image"]["NewNRGBA64"] = image.NewNRGBA64
Api["image"]["NYCbCrA"] = image.NYCbCrA{}
Api["image"]["NewNYCbCrA"] = image.NewNYCbCrA
Api["image"]["Paletted"] = image.Paletted{}
Api["image"]["NewPaletted"] = image.NewPaletted
Api["image"]["Point"] = image.Point{}
Api["image"]["ZP"] = image.ZP
Api["image"]["Pt"] = image.Pt
Api["image"]["RGBA"] = image.RGBA{}
Api["image"]["NewRGBA"] = image.NewRGBA
Api["image"]["RGBA64"] = image.RGBA64{}
Api["image"]["NewRGBA64"] = image.NewRGBA64
Api["image"]["Rectangle"] = image.Rectangle{}
Api["image"]["ZR"] = image.ZR
Api["image"]["Rect"] = image.Rect
Api["image"]["Uniform"] = image.Uniform{}
Api["image"]["NewUniform"] = image.NewUniform
Api["image"]["YCbCr"] = image.YCbCr{}
Api["image"]["NewYCbCr"] = image.NewYCbCr
Api["image"]["YCbCrSubsampleRatio444"] = image.YCbCrSubsampleRatio444
Api["image"]["YCbCrSubsampleRatio422"] = image.YCbCrSubsampleRatio422
Api["image"]["YCbCrSubsampleRatio420"] = image.YCbCrSubsampleRatio420
Api["image"]["YCbCrSubsampleRatio440"] = image.YCbCrSubsampleRatio440
Api["image"]["YCbCrSubsampleRatio411"] = image.YCbCrSubsampleRatio411
Api["image"]["YCbCrSubsampleRatio410"] = image.YCbCrSubsampleRatio410
Api["image"]["RegisterFormat"] = image.RegisterFormat
Api["image"]["NewAlpha"] = image.NewAlpha
Api["image"]["NewAlpha16"] = image.NewAlpha16
Api["image"]["NewCMYK"] = image.NewCMYK
Api["image"]["DecodeConfig"] = image.DecodeConfig
Api["image"]["NewGray"] = image.NewGray
Api["image"]["NewGray16"] = image.NewGray16
Api["image"]["Decode"] = image.Decode
Api["image"]["NewNRGBA"] = image.NewNRGBA
Api["image"]["NewNRGBA64"] = image.NewNRGBA64
Api["image"]["NewNYCbCrA"] = image.NewNYCbCrA
Api["image"]["NewPaletted"] = image.NewPaletted
Api["image"]["Pt"] = image.Pt
Api["image"]["NewRGBA"] = image.NewRGBA
Api["image"]["NewRGBA64"] = image.NewRGBA64
Api["image"]["Rect"] = image.Rect
Api["image"]["NewUniform"] = image.NewUniform
Api["image"]["NewYCbCr"] = image.NewYCbCr

}