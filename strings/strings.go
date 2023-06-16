package strings
import (
strings "strings"
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
if _, ok := DefaultGojaInterfaces["strings"]; !ok {
   DefaultGojaInterfaces["strings"] = map[string]interface{}{}
}
DefaultGojaInterfaces["strings"]["Clone"] = strings.Clone
DefaultGojaInterfaces["strings"]["Compare"] = strings.Compare
DefaultGojaInterfaces["strings"]["Contains"] = strings.Contains
DefaultGojaInterfaces["strings"]["ContainsAny"] = strings.ContainsAny
DefaultGojaInterfaces["strings"]["ContainsRune"] = strings.ContainsRune
DefaultGojaInterfaces["strings"]["Count"] = strings.Count
DefaultGojaInterfaces["strings"]["Cut"] = strings.Cut
DefaultGojaInterfaces["strings"]["CutPrefix"] = strings.CutPrefix
DefaultGojaInterfaces["strings"]["CutSuffix"] = strings.CutSuffix
DefaultGojaInterfaces["strings"]["EqualFold"] = strings.EqualFold
DefaultGojaInterfaces["strings"]["Fields"] = strings.Fields
DefaultGojaInterfaces["strings"]["FieldsFunc"] = strings.FieldsFunc
DefaultGojaInterfaces["strings"]["HasPrefix"] = strings.HasPrefix
DefaultGojaInterfaces["strings"]["HasSuffix"] = strings.HasSuffix
DefaultGojaInterfaces["strings"]["Index"] = strings.Index
DefaultGojaInterfaces["strings"]["IndexAny"] = strings.IndexAny
DefaultGojaInterfaces["strings"]["IndexByte"] = strings.IndexByte
DefaultGojaInterfaces["strings"]["IndexFunc"] = strings.IndexFunc
DefaultGojaInterfaces["strings"]["IndexRune"] = strings.IndexRune
DefaultGojaInterfaces["strings"]["Join"] = strings.Join
DefaultGojaInterfaces["strings"]["LastIndex"] = strings.LastIndex
DefaultGojaInterfaces["strings"]["LastIndexAny"] = strings.LastIndexAny
DefaultGojaInterfaces["strings"]["LastIndexByte"] = strings.LastIndexByte
DefaultGojaInterfaces["strings"]["LastIndexFunc"] = strings.LastIndexFunc
DefaultGojaInterfaces["strings"]["Map"] = strings.Map
DefaultGojaInterfaces["strings"]["Repeat"] = strings.Repeat
DefaultGojaInterfaces["strings"]["Replace"] = strings.Replace
DefaultGojaInterfaces["strings"]["ReplaceAll"] = strings.ReplaceAll
DefaultGojaInterfaces["strings"]["Split"] = strings.Split
DefaultGojaInterfaces["strings"]["SplitAfter"] = strings.SplitAfter
DefaultGojaInterfaces["strings"]["SplitAfterN"] = strings.SplitAfterN
DefaultGojaInterfaces["strings"]["SplitN"] = strings.SplitN
DefaultGojaInterfaces["strings"]["Title"] = strings.Title
DefaultGojaInterfaces["strings"]["ToLower"] = strings.ToLower
DefaultGojaInterfaces["strings"]["ToLowerSpecial"] = strings.ToLowerSpecial
DefaultGojaInterfaces["strings"]["ToTitle"] = strings.ToTitle
DefaultGojaInterfaces["strings"]["ToTitleSpecial"] = strings.ToTitleSpecial
DefaultGojaInterfaces["strings"]["ToUpper"] = strings.ToUpper
DefaultGojaInterfaces["strings"]["ToUpperSpecial"] = strings.ToUpperSpecial
DefaultGojaInterfaces["strings"]["ToValidUTF8"] = strings.ToValidUTF8
DefaultGojaInterfaces["strings"]["Trim"] = strings.Trim
DefaultGojaInterfaces["strings"]["TrimFunc"] = strings.TrimFunc
DefaultGojaInterfaces["strings"]["TrimLeft"] = strings.TrimLeft
DefaultGojaInterfaces["strings"]["TrimLeftFunc"] = strings.TrimLeftFunc
DefaultGojaInterfaces["strings"]["TrimPrefix"] = strings.TrimPrefix
DefaultGojaInterfaces["strings"]["TrimRight"] = strings.TrimRight
DefaultGojaInterfaces["strings"]["TrimRightFunc"] = strings.TrimRightFunc
DefaultGojaInterfaces["strings"]["TrimSpace"] = strings.TrimSpace
DefaultGojaInterfaces["strings"]["TrimSuffix"] = strings.TrimSuffix
DefaultGojaInterfaces["strings"]["Builder"] = strings.Builder{}
DefaultGojaInterfaces["strings"]["Reader"] = strings.Reader{}
DefaultGojaInterfaces["strings"]["NewReader"] = strings.NewReader
DefaultGojaInterfaces["strings"]["Replacer"] = strings.Replacer{}
DefaultGojaInterfaces["strings"]["NewReplacer"] = strings.NewReplacer
DefaultGojaInterfaces["strings"]["Clone"] = strings.Clone
DefaultGojaInterfaces["strings"]["Compare"] = strings.Compare
DefaultGojaInterfaces["strings"]["Contains"] = strings.Contains
DefaultGojaInterfaces["strings"]["ContainsAny"] = strings.ContainsAny
DefaultGojaInterfaces["strings"]["ContainsRune"] = strings.ContainsRune
DefaultGojaInterfaces["strings"]["Count"] = strings.Count
DefaultGojaInterfaces["strings"]["Cut"] = strings.Cut
DefaultGojaInterfaces["strings"]["CutPrefix"] = strings.CutPrefix
DefaultGojaInterfaces["strings"]["CutSuffix"] = strings.CutSuffix
DefaultGojaInterfaces["strings"]["EqualFold"] = strings.EqualFold
DefaultGojaInterfaces["strings"]["Fields"] = strings.Fields
DefaultGojaInterfaces["strings"]["FieldsFunc"] = strings.FieldsFunc
DefaultGojaInterfaces["strings"]["HasPrefix"] = strings.HasPrefix
DefaultGojaInterfaces["strings"]["HasSuffix"] = strings.HasSuffix
DefaultGojaInterfaces["strings"]["Index"] = strings.Index
DefaultGojaInterfaces["strings"]["IndexAny"] = strings.IndexAny
DefaultGojaInterfaces["strings"]["IndexByte"] = strings.IndexByte
DefaultGojaInterfaces["strings"]["IndexFunc"] = strings.IndexFunc
DefaultGojaInterfaces["strings"]["IndexRune"] = strings.IndexRune
DefaultGojaInterfaces["strings"]["Join"] = strings.Join
DefaultGojaInterfaces["strings"]["LastIndex"] = strings.LastIndex
DefaultGojaInterfaces["strings"]["LastIndexAny"] = strings.LastIndexAny
DefaultGojaInterfaces["strings"]["LastIndexByte"] = strings.LastIndexByte
DefaultGojaInterfaces["strings"]["LastIndexFunc"] = strings.LastIndexFunc
DefaultGojaInterfaces["strings"]["Map"] = strings.Map
DefaultGojaInterfaces["strings"]["Repeat"] = strings.Repeat
DefaultGojaInterfaces["strings"]["Replace"] = strings.Replace
DefaultGojaInterfaces["strings"]["ReplaceAll"] = strings.ReplaceAll
DefaultGojaInterfaces["strings"]["Split"] = strings.Split
DefaultGojaInterfaces["strings"]["SplitAfter"] = strings.SplitAfter
DefaultGojaInterfaces["strings"]["SplitAfterN"] = strings.SplitAfterN
DefaultGojaInterfaces["strings"]["SplitN"] = strings.SplitN
DefaultGojaInterfaces["strings"]["Title"] = strings.Title
DefaultGojaInterfaces["strings"]["ToLower"] = strings.ToLower
DefaultGojaInterfaces["strings"]["ToLowerSpecial"] = strings.ToLowerSpecial
DefaultGojaInterfaces["strings"]["ToTitle"] = strings.ToTitle
DefaultGojaInterfaces["strings"]["ToTitleSpecial"] = strings.ToTitleSpecial
DefaultGojaInterfaces["strings"]["ToUpper"] = strings.ToUpper
DefaultGojaInterfaces["strings"]["ToUpperSpecial"] = strings.ToUpperSpecial
DefaultGojaInterfaces["strings"]["ToValidUTF8"] = strings.ToValidUTF8
DefaultGojaInterfaces["strings"]["Trim"] = strings.Trim
DefaultGojaInterfaces["strings"]["TrimFunc"] = strings.TrimFunc
DefaultGojaInterfaces["strings"]["TrimLeft"] = strings.TrimLeft
DefaultGojaInterfaces["strings"]["TrimLeftFunc"] = strings.TrimLeftFunc
DefaultGojaInterfaces["strings"]["TrimPrefix"] = strings.TrimPrefix
DefaultGojaInterfaces["strings"]["TrimRight"] = strings.TrimRight
DefaultGojaInterfaces["strings"]["TrimRightFunc"] = strings.TrimRightFunc
DefaultGojaInterfaces["strings"]["TrimSpace"] = strings.TrimSpace
DefaultGojaInterfaces["strings"]["TrimSuffix"] = strings.TrimSuffix
DefaultGojaInterfaces["strings"]["NewReader"] = strings.NewReader
DefaultGojaInterfaces["strings"]["NewReplacer"] = strings.NewReplacer

			 			}	
			 		}
			 	}
			 }
		}
	}
}