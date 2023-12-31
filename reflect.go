package ninjascript
import (
reflect "reflect"
)
func init() {if _, ok := Api["reflect"]; !ok {
   Api["reflect"] = map[string]interface{}{}
}
Api["reflect"]["Copy"] = reflect.Copy
Api["reflect"]["DeepEqual"] = reflect.DeepEqual
Api["reflect"]["Swapper"] = reflect.Swapper
Api["reflect"]["RecvDir"] = reflect.RecvDir
Api["reflect"]["SendDir"] = reflect.SendDir
Api["reflect"]["BothDir"] = reflect.BothDir
Api["reflect"]["Invalid"] = reflect.Invalid
Api["reflect"]["Bool"] = reflect.Bool
Api["reflect"]["Int"] = reflect.Int
Api["reflect"]["Int8"] = reflect.Int8
Api["reflect"]["Int16"] = reflect.Int16
Api["reflect"]["Int32"] = reflect.Int32
Api["reflect"]["Int64"] = reflect.Int64
Api["reflect"]["Uint"] = reflect.Uint
Api["reflect"]["Uint8"] = reflect.Uint8
Api["reflect"]["Uint16"] = reflect.Uint16
Api["reflect"]["Uint32"] = reflect.Uint32
Api["reflect"]["Uint64"] = reflect.Uint64
Api["reflect"]["Uintptr"] = reflect.Uintptr
Api["reflect"]["Float32"] = reflect.Float32
Api["reflect"]["Float64"] = reflect.Float64
Api["reflect"]["Complex64"] = reflect.Complex64
Api["reflect"]["Complex128"] = reflect.Complex128
Api["reflect"]["Array"] = reflect.Array
Api["reflect"]["Chan"] = reflect.Chan
Api["reflect"]["Func"] = reflect.Func
Api["reflect"]["Interface"] = reflect.Interface
Api["reflect"]["Map"] = reflect.Map
Api["reflect"]["Pointer"] = reflect.Pointer
Api["reflect"]["Slice"] = reflect.Slice
Api["reflect"]["String"] = reflect.String
Api["reflect"]["Struct"] = reflect.Struct
Api["reflect"]["UnsafePointer"] = reflect.UnsafePointer
Api["reflect"]["MapIter"] = reflect.MapIter{}
Api["reflect"]["Method"] = reflect.Method{}
Api["reflect"]["SelectCase"] = reflect.SelectCase{}
Api["reflect"]["SelectSend"] = reflect.SelectSend
Api["reflect"]["SelectRecv"] = reflect.SelectRecv
Api["reflect"]["SelectDefault"] = reflect.SelectDefault
Api["reflect"]["SliceHeader"] = reflect.SliceHeader{}
Api["reflect"]["StringHeader"] = reflect.StringHeader{}
Api["reflect"]["StructField"] = reflect.StructField{}
Api["reflect"]["VisibleFields"] = reflect.VisibleFields
Api["reflect"]["Type"] = new(reflect.Type)
Api["reflect"]["ArrayOf"] = reflect.ArrayOf
Api["reflect"]["ChanOf"] = reflect.ChanOf
Api["reflect"]["FuncOf"] = reflect.FuncOf
Api["reflect"]["MapOf"] = reflect.MapOf
Api["reflect"]["PointerTo"] = reflect.PointerTo
Api["reflect"]["PtrTo"] = reflect.PtrTo
Api["reflect"]["SliceOf"] = reflect.SliceOf
Api["reflect"]["StructOf"] = reflect.StructOf
Api["reflect"]["TypeOf"] = reflect.TypeOf
Api["reflect"]["Value"] = reflect.Value{}
Api["reflect"]["Append"] = reflect.Append
Api["reflect"]["AppendSlice"] = reflect.AppendSlice
Api["reflect"]["Indirect"] = reflect.Indirect
Api["reflect"]["MakeChan"] = reflect.MakeChan
Api["reflect"]["MakeFunc"] = reflect.MakeFunc
Api["reflect"]["MakeMap"] = reflect.MakeMap
Api["reflect"]["MakeMapWithSize"] = reflect.MakeMapWithSize
Api["reflect"]["MakeSlice"] = reflect.MakeSlice
Api["reflect"]["New"] = reflect.New
Api["reflect"]["NewAt"] = reflect.NewAt
Api["reflect"]["Select"] = reflect.Select
Api["reflect"]["ValueOf"] = reflect.ValueOf
Api["reflect"]["Zero"] = reflect.Zero
Api["reflect"]["ValueError"] = reflect.ValueError{}
Api["reflect"]["Copy"] = reflect.Copy
Api["reflect"]["DeepEqual"] = reflect.DeepEqual
Api["reflect"]["Swapper"] = reflect.Swapper
Api["reflect"]["VisibleFields"] = reflect.VisibleFields
Api["reflect"]["ArrayOf"] = reflect.ArrayOf
Api["reflect"]["ChanOf"] = reflect.ChanOf
Api["reflect"]["FuncOf"] = reflect.FuncOf
Api["reflect"]["MapOf"] = reflect.MapOf
Api["reflect"]["PointerTo"] = reflect.PointerTo
Api["reflect"]["PtrTo"] = reflect.PtrTo
Api["reflect"]["SliceOf"] = reflect.SliceOf
Api["reflect"]["StructOf"] = reflect.StructOf
Api["reflect"]["TypeOf"] = reflect.TypeOf
Api["reflect"]["Append"] = reflect.Append
Api["reflect"]["AppendSlice"] = reflect.AppendSlice
Api["reflect"]["Indirect"] = reflect.Indirect
Api["reflect"]["MakeChan"] = reflect.MakeChan
Api["reflect"]["MakeFunc"] = reflect.MakeFunc
Api["reflect"]["MakeMap"] = reflect.MakeMap
Api["reflect"]["MakeMapWithSize"] = reflect.MakeMapWithSize
Api["reflect"]["MakeSlice"] = reflect.MakeSlice
Api["reflect"]["New"] = reflect.New
Api["reflect"]["NewAt"] = reflect.NewAt
Api["reflect"]["Select"] = reflect.Select
Api["reflect"]["ValueOf"] = reflect.ValueOf
Api["reflect"]["Zero"] = reflect.Zero

}