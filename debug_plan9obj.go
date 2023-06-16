package ninjascript
import (
debug_plan9obj "debug/plan9obj"
)
func init() {if _, ok := Api["debug/plan9obj"]; !ok {
   Api["debug/plan9obj"] = map[string]interface{}{}
}
Api["debug/plan9obj"]["Magic64"] = debug_plan9obj.Magic64
Api["debug/plan9obj"]["Magic386"] = debug_plan9obj.Magic386
Api["debug/plan9obj"]["MagicAMD64"] = debug_plan9obj.MagicAMD64
Api["debug/plan9obj"]["MagicARM"] = debug_plan9obj.MagicARM
Api["debug/plan9obj"]["ErrNoSymbols"] = debug_plan9obj.ErrNoSymbols
Api["debug/plan9obj"]["File"] = debug_plan9obj.File{}
Api["debug/plan9obj"]["NewFile"] = debug_plan9obj.NewFile
Api["debug/plan9obj"]["Open"] = debug_plan9obj.Open
Api["debug/plan9obj"]["FileHeader"] = debug_plan9obj.FileHeader{}
Api["debug/plan9obj"]["Section"] = debug_plan9obj.Section{}
Api["debug/plan9obj"]["SectionHeader"] = debug_plan9obj.SectionHeader{}
Api["debug/plan9obj"]["Sym"] = debug_plan9obj.Sym{}
Api["debug/plan9obj"]["NewFile"] = debug_plan9obj.NewFile
Api["debug/plan9obj"]["Open"] = debug_plan9obj.Open

}