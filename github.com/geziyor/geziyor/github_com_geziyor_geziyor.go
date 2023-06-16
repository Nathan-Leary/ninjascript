package geziyor
import (
github_com_geziyor_geziyor "github.com/geziyor/geziyor"
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
if _, ok := DefaultGojaInterfaces["geziyor"]; !ok {
   DefaultGojaInterfaces["geziyor"] = map[string]interface{}{}
}
DefaultGojaInterfaces["geziyor"]["Geziyor"] = github_com_geziyor_geziyor.Geziyor{}
DefaultGojaInterfaces["geziyor"]["NewGeziyor"] = github_com_geziyor_geziyor.NewGeziyor
DefaultGojaInterfaces["geziyor"]["Options"] = github_com_geziyor_geziyor.Options{}
DefaultGojaInterfaces["geziyor"]["NewGeziyor"] = github_com_geziyor_geziyor.NewGeziyor

			 			}	
			 		}
			 	}
			 }
		}
	}
}