package ninjascript
import (
bytes "bytes"
)
func init() {if _, ok := Api["bytes"]; !ok {
   Api["bytes"] = map[string]interface{}{}
}
if _, ok := Api["bytes/new"]; !ok {
	Api["bytes/new"] = map[string]interface{}{}
}
Api["bytes"]["ErrTooLarge"] = bytes.ErrTooLarge
Api["bytes"]["Clone"] = bytes.Clone
Api["bytes"]["Compare"] = bytes.Compare
Api["bytes"]["Contains"] = bytes.Contains
Api["bytes"]["ContainsAny"] = bytes.ContainsAny
Api["bytes"]["ContainsRune"] = bytes.ContainsRune
Api["bytes"]["Count"] = bytes.Count
Api["bytes"]["Cut"] = bytes.Cut
Api["bytes"]["CutPrefix"] = bytes.CutPrefix
Api["bytes"]["CutSuffix"] = bytes.CutSuffix
Api["bytes"]["Equal"] = bytes.Equal
Api["bytes"]["EqualFold"] = bytes.EqualFold
Api["bytes"]["Fields"] = bytes.Fields
Api["bytes"]["FieldsFunc"] = bytes.FieldsFunc
Api["bytes"]["HasPrefix"] = bytes.HasPrefix
Api["bytes"]["HasSuffix"] = bytes.HasSuffix
Api["bytes"]["Index"] = bytes.Index
Api["bytes"]["IndexAny"] = bytes.IndexAny
Api["bytes"]["IndexByte"] = bytes.IndexByte
Api["bytes"]["IndexFunc"] = bytes.IndexFunc
Api["bytes"]["IndexRune"] = bytes.IndexRune
Api["bytes"]["Join"] = bytes.Join
Api["bytes"]["LastIndex"] = bytes.LastIndex
Api["bytes"]["LastIndexAny"] = bytes.LastIndexAny
Api["bytes"]["LastIndexByte"] = bytes.LastIndexByte
Api["bytes"]["LastIndexFunc"] = bytes.LastIndexFunc
Api["bytes"]["Repeat"] = bytes.Repeat
Api["bytes"]["Replace"] = bytes.Replace
Api["bytes"]["ReplaceAll"] = bytes.ReplaceAll
Api["bytes"]["Runes"] = bytes.Runes
Api["bytes"]["Split"] = bytes.Split
Api["bytes"]["SplitAfter"] = bytes.SplitAfter
Api["bytes"]["SplitAfterN"] = bytes.SplitAfterN
Api["bytes"]["SplitN"] = bytes.SplitN
Api["bytes"]["Title"] = bytes.Title
Api["bytes"]["ToLower"] = bytes.ToLower
Api["bytes"]["ToLowerSpecial"] = bytes.ToLowerSpecial
Api["bytes"]["ToTitle"] = bytes.ToTitle
Api["bytes"]["ToTitleSpecial"] = bytes.ToTitleSpecial
Api["bytes"]["ToUpper"] = bytes.ToUpper
Api["bytes"]["ToUpperSpecial"] = bytes.ToUpperSpecial
Api["bytes"]["ToValidUTF8"] = bytes.ToValidUTF8
Api["bytes"]["Trim"] = bytes.Trim
Api["bytes"]["TrimFunc"] = bytes.TrimFunc
Api["bytes"]["TrimLeft"] = bytes.TrimLeft
Api["bytes"]["TrimLeftFunc"] = bytes.TrimLeftFunc
Api["bytes"]["TrimPrefix"] = bytes.TrimPrefix
Api["bytes"]["TrimRight"] = bytes.TrimRight
Api["bytes"]["TrimRightFunc"] = bytes.TrimRightFunc
Api["bytes"]["TrimSpace"] = bytes.TrimSpace
Api["bytes"]["TrimSuffix"] = bytes.TrimSuffix
Api["bytes"]["Buffer"] = bytes.Buffer{}
Api["bytes/new"]["Buffer"] = func(args ...interface{}) interface{} {
					return new(bytes.Buffer)
}
Api["bytes"]["NewBuffer"] = bytes.NewBuffer
Api["bytes"]["NewBufferString"] = bytes.NewBufferString

}
