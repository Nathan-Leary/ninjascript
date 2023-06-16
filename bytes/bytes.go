package bytes
import (
bytes "bytes"
)
func Load(Interfaces ...interface{}) {
	if len(Interfaces) > 5 {
		if i, ok := Interfaces[0].(map[string]map[string]interface{}) {
			 DefaultGojaInterfaces: = i
			 if i, ok := Interfaces[1].(*goja.Runtime) {
				vm := i
				if i, ok := Interfaces[2].(*quickjs.Context) {
					ctx := i	
					if i, ok := Interfaces[3].(func(*quickjs.Context, interface{}) (quickjs.Value, string)) {
						ConvertInterfaceToQuickJS := i
						if i, ok := Interfaces[4].(func(interface{}, ...interface{}) interface{}) {
							ExecuteFunction := i
if _, ok := DefaultGojaInterfaces["bytes"]; !ok {
   DefaultGojaInterfaces["bytes"] = map[string]interface{}{}
}
DefaultGojaInterfaces["bytes"]["ErrTooLarge"] = bytes.ErrTooLarge
DefaultGojaInterfaces["bytes"]["Clone"] = bytes.Clone
DefaultGojaInterfaces["bytes"]["Compare"] = bytes.Compare
DefaultGojaInterfaces["bytes"]["Contains"] = bytes.Contains
DefaultGojaInterfaces["bytes"]["ContainsAny"] = bytes.ContainsAny
DefaultGojaInterfaces["bytes"]["ContainsRune"] = bytes.ContainsRune
DefaultGojaInterfaces["bytes"]["Count"] = bytes.Count
DefaultGojaInterfaces["bytes"]["Cut"] = bytes.Cut
DefaultGojaInterfaces["bytes"]["CutPrefix"] = bytes.CutPrefix
DefaultGojaInterfaces["bytes"]["CutSuffix"] = bytes.CutSuffix
DefaultGojaInterfaces["bytes"]["Equal"] = bytes.Equal
DefaultGojaInterfaces["bytes"]["EqualFold"] = bytes.EqualFold
DefaultGojaInterfaces["bytes"]["Fields"] = bytes.Fields
DefaultGojaInterfaces["bytes"]["FieldsFunc"] = bytes.FieldsFunc
DefaultGojaInterfaces["bytes"]["HasPrefix"] = bytes.HasPrefix
DefaultGojaInterfaces["bytes"]["HasSuffix"] = bytes.HasSuffix
DefaultGojaInterfaces["bytes"]["Index"] = bytes.Index
DefaultGojaInterfaces["bytes"]["IndexAny"] = bytes.IndexAny
DefaultGojaInterfaces["bytes"]["IndexByte"] = bytes.IndexByte
DefaultGojaInterfaces["bytes"]["IndexFunc"] = bytes.IndexFunc
DefaultGojaInterfaces["bytes"]["IndexRune"] = bytes.IndexRune
DefaultGojaInterfaces["bytes"]["Join"] = bytes.Join
DefaultGojaInterfaces["bytes"]["LastIndex"] = bytes.LastIndex
DefaultGojaInterfaces["bytes"]["LastIndexAny"] = bytes.LastIndexAny
DefaultGojaInterfaces["bytes"]["LastIndexByte"] = bytes.LastIndexByte
DefaultGojaInterfaces["bytes"]["LastIndexFunc"] = bytes.LastIndexFunc
DefaultGojaInterfaces["bytes"]["Map"] = bytes.Map
DefaultGojaInterfaces["bytes"]["Repeat"] = bytes.Repeat
DefaultGojaInterfaces["bytes"]["Replace"] = bytes.Replace
DefaultGojaInterfaces["bytes"]["ReplaceAll"] = bytes.ReplaceAll
DefaultGojaInterfaces["bytes"]["Runes"] = bytes.Runes
DefaultGojaInterfaces["bytes"]["Split"] = bytes.Split
DefaultGojaInterfaces["bytes"]["SplitAfter"] = bytes.SplitAfter
DefaultGojaInterfaces["bytes"]["SplitAfterN"] = bytes.SplitAfterN
DefaultGojaInterfaces["bytes"]["SplitN"] = bytes.SplitN
DefaultGojaInterfaces["bytes"]["Title"] = bytes.Title
DefaultGojaInterfaces["bytes"]["ToLower"] = bytes.ToLower
DefaultGojaInterfaces["bytes"]["ToLowerSpecial"] = bytes.ToLowerSpecial
DefaultGojaInterfaces["bytes"]["ToTitle"] = bytes.ToTitle
DefaultGojaInterfaces["bytes"]["ToTitleSpecial"] = bytes.ToTitleSpecial
DefaultGojaInterfaces["bytes"]["ToUpper"] = bytes.ToUpper
DefaultGojaInterfaces["bytes"]["ToUpperSpecial"] = bytes.ToUpperSpecial
DefaultGojaInterfaces["bytes"]["ToValidUTF8"] = bytes.ToValidUTF8
DefaultGojaInterfaces["bytes"]["Trim"] = bytes.Trim
DefaultGojaInterfaces["bytes"]["TrimFunc"] = bytes.TrimFunc
DefaultGojaInterfaces["bytes"]["TrimLeft"] = bytes.TrimLeft
DefaultGojaInterfaces["bytes"]["TrimLeftFunc"] = bytes.TrimLeftFunc
DefaultGojaInterfaces["bytes"]["TrimPrefix"] = bytes.TrimPrefix
DefaultGojaInterfaces["bytes"]["TrimRight"] = bytes.TrimRight
DefaultGojaInterfaces["bytes"]["TrimRightFunc"] = bytes.TrimRightFunc
DefaultGojaInterfaces["bytes"]["TrimSpace"] = bytes.TrimSpace
DefaultGojaInterfaces["bytes"]["TrimSuffix"] = bytes.TrimSuffix
DefaultGojaInterfaces["bytes"]["Buffer"] = bytes.Buffer{}
DefaultGojaInterfaces["bytes"]["NewBuffer"] = bytes.NewBuffer
DefaultGojaInterfaces["bytes"]["NewBufferString"] = bytes.NewBufferString
DefaultGojaInterfaces["bytes"]["Reader"] = bytes.Reader{}
DefaultGojaInterfaces["bytes"]["NewReader"] = bytes.NewReader
DefaultGojaInterfaces["bytes"]["Clone"] = bytes.Clone
DefaultGojaInterfaces["bytes"]["Compare"] = bytes.Compare
DefaultGojaInterfaces["bytes"]["Contains"] = bytes.Contains
DefaultGojaInterfaces["bytes"]["ContainsAny"] = bytes.ContainsAny
DefaultGojaInterfaces["bytes"]["ContainsRune"] = bytes.ContainsRune
DefaultGojaInterfaces["bytes"]["Count"] = bytes.Count
DefaultGojaInterfaces["bytes"]["Cut"] = bytes.Cut
DefaultGojaInterfaces["bytes"]["CutPrefix"] = bytes.CutPrefix
DefaultGojaInterfaces["bytes"]["CutSuffix"] = bytes.CutSuffix
DefaultGojaInterfaces["bytes"]["Equal"] = bytes.Equal
DefaultGojaInterfaces["bytes"]["EqualFold"] = bytes.EqualFold
DefaultGojaInterfaces["bytes"]["Fields"] = bytes.Fields
DefaultGojaInterfaces["bytes"]["FieldsFunc"] = bytes.FieldsFunc
DefaultGojaInterfaces["bytes"]["HasPrefix"] = bytes.HasPrefix
DefaultGojaInterfaces["bytes"]["HasSuffix"] = bytes.HasSuffix
DefaultGojaInterfaces["bytes"]["Index"] = bytes.Index
DefaultGojaInterfaces["bytes"]["IndexAny"] = bytes.IndexAny
DefaultGojaInterfaces["bytes"]["IndexByte"] = bytes.IndexByte
DefaultGojaInterfaces["bytes"]["IndexFunc"] = bytes.IndexFunc
DefaultGojaInterfaces["bytes"]["IndexRune"] = bytes.IndexRune
DefaultGojaInterfaces["bytes"]["Join"] = bytes.Join
DefaultGojaInterfaces["bytes"]["LastIndex"] = bytes.LastIndex
DefaultGojaInterfaces["bytes"]["LastIndexAny"] = bytes.LastIndexAny
DefaultGojaInterfaces["bytes"]["LastIndexByte"] = bytes.LastIndexByte
DefaultGojaInterfaces["bytes"]["LastIndexFunc"] = bytes.LastIndexFunc
DefaultGojaInterfaces["bytes"]["Map"] = bytes.Map
DefaultGojaInterfaces["bytes"]["Repeat"] = bytes.Repeat
DefaultGojaInterfaces["bytes"]["Replace"] = bytes.Replace
DefaultGojaInterfaces["bytes"]["ReplaceAll"] = bytes.ReplaceAll
DefaultGojaInterfaces["bytes"]["Runes"] = bytes.Runes
DefaultGojaInterfaces["bytes"]["Split"] = bytes.Split
DefaultGojaInterfaces["bytes"]["SplitAfter"] = bytes.SplitAfter
DefaultGojaInterfaces["bytes"]["SplitAfterN"] = bytes.SplitAfterN
DefaultGojaInterfaces["bytes"]["SplitN"] = bytes.SplitN
DefaultGojaInterfaces["bytes"]["Title"] = bytes.Title
DefaultGojaInterfaces["bytes"]["ToLower"] = bytes.ToLower
DefaultGojaInterfaces["bytes"]["ToLowerSpecial"] = bytes.ToLowerSpecial
DefaultGojaInterfaces["bytes"]["ToTitle"] = bytes.ToTitle
DefaultGojaInterfaces["bytes"]["ToTitleSpecial"] = bytes.ToTitleSpecial
DefaultGojaInterfaces["bytes"]["ToUpper"] = bytes.ToUpper
DefaultGojaInterfaces["bytes"]["ToUpperSpecial"] = bytes.ToUpperSpecial
DefaultGojaInterfaces["bytes"]["ToValidUTF8"] = bytes.ToValidUTF8
DefaultGojaInterfaces["bytes"]["Trim"] = bytes.Trim
DefaultGojaInterfaces["bytes"]["TrimFunc"] = bytes.TrimFunc
DefaultGojaInterfaces["bytes"]["TrimLeft"] = bytes.TrimLeft
DefaultGojaInterfaces["bytes"]["TrimLeftFunc"] = bytes.TrimLeftFunc
DefaultGojaInterfaces["bytes"]["TrimPrefix"] = bytes.TrimPrefix
DefaultGojaInterfaces["bytes"]["TrimRight"] = bytes.TrimRight
DefaultGojaInterfaces["bytes"]["TrimRightFunc"] = bytes.TrimRightFunc
DefaultGojaInterfaces["bytes"]["TrimSpace"] = bytes.TrimSpace
DefaultGojaInterfaces["bytes"]["TrimSuffix"] = bytes.TrimSuffix
DefaultGojaInterfaces["bytes"]["NewBuffer"] = bytes.NewBuffer
DefaultGojaInterfaces["bytes"]["NewBufferString"] = bytes.NewBufferString
DefaultGojaInterfaces["bytes"]["NewReader"] = bytes.NewReader

			 			}	
			 		}
			 	}
			 }
		}
	}
}