package ninjascript
import (
image_draw "image/draw"
)
func init() {if _, ok := Api["image/draw"]; !ok {
   Api["image/draw"] = map[string]interface{}{}
}
Api["image/draw"]["Draw"] = image_draw.Draw
Api["image/draw"]["DrawMask"] = image_draw.DrawMask
Api["image/draw"]["FloydSteinberg"] = image_draw.FloydSteinberg
Api["image/draw"]["Over"] = image_draw.Over
Api["image/draw"]["Src"] = image_draw.Src
Api["image/draw"]["Draw"] = image_draw.Draw
Api["image/draw"]["DrawMask"] = image_draw.DrawMask

}