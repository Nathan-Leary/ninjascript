package gomponents
import (
github_com_maragudk_gomponents "github.com/maragudk/gomponents"
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
if _, ok := DefaultGojaInterfaces["gomponents"]; !ok {
   DefaultGojaInterfaces["gomponents"] = map[string]interface{}{}
}
DefaultGojaInterfaces["gomponents"]["ElementType"] = github_com_maragudk_gomponents.ElementType
DefaultGojaInterfaces["gomponents"]["AttributeType"] = github_com_maragudk_gomponents.AttributeType
DefaultGojaInterfaces["gomponents"]["Attr"] = github_com_maragudk_gomponents.Attr
DefaultGojaInterfaces["gomponents"]["El"] = github_com_maragudk_gomponents.El
DefaultGojaInterfaces["gomponents"]["Group"] = github_com_maragudk_gomponents.Group
DefaultGojaInterfaces["gomponents"]["If"] = github_com_maragudk_gomponents.If
DefaultGojaInterfaces["gomponents"]["Map"] = github_com_maragudk_gomponents.Map
DefaultGojaInterfaces["gomponents"]["Raw"] = github_com_maragudk_gomponents.Raw
DefaultGojaInterfaces["gomponents"]["Rawf"] = github_com_maragudk_gomponents.Rawf
DefaultGojaInterfaces["gomponents"]["Text"] = github_com_maragudk_gomponents.Text
DefaultGojaInterfaces["gomponents"]["Textf"] = github_com_maragudk_gomponents.Textf
DefaultGojaInterfaces["gomponents"]["Attr"] = github_com_maragudk_gomponents.Attr
DefaultGojaInterfaces["gomponents"]["El"] = github_com_maragudk_gomponents.El
DefaultGojaInterfaces["gomponents"]["Group"] = github_com_maragudk_gomponents.Group
DefaultGojaInterfaces["gomponents"]["If"] = github_com_maragudk_gomponents.If
DefaultGojaInterfaces["gomponents"]["Map"] = github_com_maragudk_gomponents.Map
DefaultGojaInterfaces["gomponents"]["Raw"] = github_com_maragudk_gomponents.Raw
DefaultGojaInterfaces["gomponents"]["Rawf"] = github_com_maragudk_gomponents.Rawf
DefaultGojaInterfaces["gomponents"]["Text"] = github_com_maragudk_gomponents.Text
DefaultGojaInterfaces["gomponents"]["Textf"] = github_com_maragudk_gomponents.Textf

			 			}	
			 		}
			 	}
			 }
		}
	}
}