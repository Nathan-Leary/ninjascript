package ninjascript
import (
image_color_palette "image/color/palette"
)
func init() {if _, ok := Api["image/color/palette"]; !ok {
   Api["image/color/palette"] = map[string]interface{}{}
}
Api["image/color/palette"]["Plan9"] = image_color_palette.Plan9
Api["image/color/palette"]["WebSafe"] = image_color_palette.WebSafe

}