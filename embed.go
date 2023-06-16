package ninjascript
import (
embed "embed"
)
func init() {if _, ok := Api["embed"]; !ok {
   Api["embed"] = map[string]interface{}{}
}
Api["embed"]["FS"] = embed.FS{}

}