package ninjascript
import (
index_suffixarray "index/suffixarray"
)
func init() {if _, ok := Api["index/suffixarray"]; !ok {
   Api["index/suffixarray"] = map[string]interface{}{}
}
Api["index/suffixarray"]["Index"] = index_suffixarray.Index{}
Api["index/suffixarray"]["New"] = index_suffixarray.New
Api["index/suffixarray"]["New"] = index_suffixarray.New

}