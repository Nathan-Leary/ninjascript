package quickjs-go-polyfill
import (
github_com_buke_quickjs_go_polyfill "github.com/buke/quickjs-go-polyfill"
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
if _, ok := DefaultGojaInterfaces["quickjsgopolyfill"]; !ok {
   DefaultGojaInterfaces["quickjsgopolyfill"] = map[string]interface{}{}
}
DefaultGojaInterfaces["quickjsgopolyfill"]["InjectAll"] = github_com_buke_quickjs_go_polyfill.InjectAll
DefaultGojaInterfaces["quickjsgopolyfill"]["InjectAll"] = github_com_buke_quickjs_go_polyfill.InjectAll

			 			}	
			 		}
			 	}
			 }
		}
	}
}