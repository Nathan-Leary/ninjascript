package ninjascript
import (
image_color "image/color"
)
func init() {if _, ok := Api["image/color"]; !ok {
   Api["image/color"] = map[string]interface{}{}
}
Api["image/color"]["CMYKToRGB"] = image_color.CMYKToRGB
Api["image/color"]["RGBToCMYK"] = image_color.RGBToCMYK
Api["image/color"]["RGBToYCbCr"] = image_color.RGBToYCbCr
Api["image/color"]["YCbCrToRGB"] = image_color.YCbCrToRGB
Api["image/color"]["Alpha"] = image_color.Alpha{}
Api["image/color"]["Alpha16"] = image_color.Alpha16{}
Api["image/color"]["CMYK"] = image_color.CMYK{}
Api["image/color"]["Gray"] = image_color.Gray{}
Api["image/color"]["Gray16"] = image_color.Gray16{}
Api["image/color"]["CMYKModel"] = image_color.CMYKModel
Api["image/color"]["NYCbCrAModel"] = image_color.NYCbCrAModel
Api["image/color"]["YCbCrModel"] = image_color.YCbCrModel
Api["image/color"]["ModelFunc"] = image_color.ModelFunc
Api["image/color"]["NRGBA"] = image_color.NRGBA{}
Api["image/color"]["NRGBA64"] = image_color.NRGBA64{}
Api["image/color"]["NYCbCrA"] = image_color.NYCbCrA{}
Api["image/color"]["RGBA"] = image_color.RGBA{}
Api["image/color"]["RGBA64"] = image_color.RGBA64{}
Api["image/color"]["YCbCr"] = image_color.YCbCr{}
Api["image/color"]["CMYKToRGB"] = image_color.CMYKToRGB
Api["image/color"]["RGBToCMYK"] = image_color.RGBToCMYK
Api["image/color"]["RGBToYCbCr"] = image_color.RGBToYCbCr
Api["image/color"]["YCbCrToRGB"] = image_color.YCbCrToRGB
Api["image/color"]["ModelFunc"] = image_color.ModelFunc

}