package quickjs-go
import (
github_com_buke_quickjs_go "github.com/buke/quickjs-go"
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
if _, ok := DefaultGojaInterfaces["quickjsgo"]; !ok {
   DefaultGojaInterfaces["quickjsgo"] = map[string]interface{}{}
}
DefaultGojaInterfaces["quickjsgo"]["Atom"] = github_com_buke_quickjs_go.Atom{}
DefaultGojaInterfaces["quickjsgo"]["Context"] = github_com_buke_quickjs_go.Context{}
DefaultGojaInterfaces["quickjsgo"]["Error"] = github_com_buke_quickjs_go.Error{}
DefaultGojaInterfaces["quickjsgo"]["Loop"] = github_com_buke_quickjs_go.Loop{}
DefaultGojaInterfaces["quickjsgo"]["NewLoop"] = github_com_buke_quickjs_go.NewLoop
DefaultGojaInterfaces["quickjsgo"]["Runtime"] = github_com_buke_quickjs_go.Runtime{}
DefaultGojaInterfaces["quickjsgo"]["NewRuntime"] = github_com_buke_quickjs_go.NewRuntime
DefaultGojaInterfaces["quickjsgo"]["Value"] = github_com_buke_quickjs_go.Value{}
DefaultGojaInterfaces["quickjsgo"]["NewLoop"] = github_com_buke_quickjs_go.NewLoop
DefaultGojaInterfaces["quickjsgo"]["NewRuntime"] = github_com_buke_quickjs_go.NewRuntime

			 			}	
			 		}
			 	}
			 }
		}
	}
}