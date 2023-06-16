package ninjascript
import (
runtime_cgo "runtime/cgo"
)
func init() {if _, ok := Api["runtime/cgo"]; !ok {
   Api["runtime/cgo"] = map[string]interface{}{}
}
Api["runtime/cgo"]["NewHandle"] = runtime_cgo.NewHandle
//Api["runtime/cgo"]["Incomplete"] = runtime_cgo.Incomplete{}
Api["runtime/cgo"]["NewHandle"] = runtime_cgo.NewHandle

}