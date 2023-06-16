package lorca
import (
github_com_zserge_lorca "github.com/zserge/lorca"
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
if _, ok := DefaultGojaInterfaces["lorca"]; !ok {
   DefaultGojaInterfaces["lorca"] = map[string]interface{}{}
}
DefaultGojaInterfaces["lorca"]["PageA4Width"] = github_com_zserge_lorca.PageA4Width
DefaultGojaInterfaces["lorca"]["PageA4Height"] = github_com_zserge_lorca.PageA4Height
DefaultGojaInterfaces["lorca"]["ChromeExecutable"] = github_com_zserge_lorca.ChromeExecutable
DefaultGojaInterfaces["lorca"]["Embed"] = github_com_zserge_lorca.Embed
DefaultGojaInterfaces["lorca"]["LocateChrome"] = github_com_zserge_lorca.LocateChrome
DefaultGojaInterfaces["lorca"]["PDF"] = github_com_zserge_lorca.PDF
DefaultGojaInterfaces["lorca"]["PNG"] = github_com_zserge_lorca.PNG
DefaultGojaInterfaces["lorca"]["PromptDownload"] = github_com_zserge_lorca.PromptDownload
DefaultGojaInterfaces["lorca"]["Bounds"] = github_com_zserge_lorca.Bounds{}
DefaultGojaInterfaces["lorca"]["UI"] = new(github_com_zserge_lorca.UI)
DefaultGojaInterfaces["lorca"]["New"] = github_com_zserge_lorca.New
DefaultGojaInterfaces["lorca"]["WindowStateNormal"] = github_com_zserge_lorca.WindowStateNormal
DefaultGojaInterfaces["lorca"]["WindowStateMaximized"] = github_com_zserge_lorca.WindowStateMaximized
DefaultGojaInterfaces["lorca"]["WindowStateMinimized"] = github_com_zserge_lorca.WindowStateMinimized
DefaultGojaInterfaces["lorca"]["WindowStateFullscreen"] = github_com_zserge_lorca.WindowStateFullscreen
DefaultGojaInterfaces["lorca"]["Embed"] = github_com_zserge_lorca.Embed
DefaultGojaInterfaces["lorca"]["LocateChrome"] = github_com_zserge_lorca.LocateChrome
DefaultGojaInterfaces["lorca"]["PDF"] = github_com_zserge_lorca.PDF
DefaultGojaInterfaces["lorca"]["PNG"] = github_com_zserge_lorca.PNG
DefaultGojaInterfaces["lorca"]["PromptDownload"] = github_com_zserge_lorca.PromptDownload
DefaultGojaInterfaces["lorca"]["New"] = github_com_zserge_lorca.New

			 			}	
			 		}
			 	}
			 }
		}
	}
}