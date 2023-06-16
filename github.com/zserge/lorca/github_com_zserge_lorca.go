package lorca
import (
github_com_zserge_lorca "github.com/zserge/lorca"
)
func Load() {
if _, ok := DefaultGojaInterfaces["github.com/zserge/lorca"]; !ok {
   DefaultGojaInterfaces["github.com/zserge/lorca"] = map[string]interface{}{}
}
DefaultGojaInterfaces["github.com/zserge/lorca"]["PageA4Width"] = github_com_zserge_lorca.PageA4Width
DefaultGojaInterfaces["github.com/zserge/lorca"]["PageA4Height"] = github_com_zserge_lorca.PageA4Height
DefaultGojaInterfaces["github.com/zserge/lorca"]["ChromeExecutable"] = github_com_zserge_lorca.ChromeExecutable
DefaultGojaInterfaces["github.com/zserge/lorca"]["Embed"] = github_com_zserge_lorca.Embed
DefaultGojaInterfaces["github.com/zserge/lorca"]["LocateChrome"] = github_com_zserge_lorca.LocateChrome
DefaultGojaInterfaces["github.com/zserge/lorca"]["PDF"] = github_com_zserge_lorca.PDF
DefaultGojaInterfaces["github.com/zserge/lorca"]["PNG"] = github_com_zserge_lorca.PNG
DefaultGojaInterfaces["github.com/zserge/lorca"]["PromptDownload"] = github_com_zserge_lorca.PromptDownload
DefaultGojaInterfaces["github.com/zserge/lorca"]["Bounds"] = github_com_zserge_lorca.Bounds{}
DefaultGojaInterfaces["github.com/zserge/lorca"]["UI"] = new(github_com_zserge_lorca.UI)
DefaultGojaInterfaces["github.com/zserge/lorca"]["New"] = github_com_zserge_lorca.New
DefaultGojaInterfaces["github.com/zserge/lorca"]["WindowStateNormal"] = github_com_zserge_lorca.WindowStateNormal
DefaultGojaInterfaces["github.com/zserge/lorca"]["WindowStateMaximized"] = github_com_zserge_lorca.WindowStateMaximized
DefaultGojaInterfaces["github.com/zserge/lorca"]["WindowStateMinimized"] = github_com_zserge_lorca.WindowStateMinimized
DefaultGojaInterfaces["github.com/zserge/lorca"]["WindowStateFullscreen"] = github_com_zserge_lorca.WindowStateFullscreen
DefaultGojaInterfaces["github.com/zserge/lorca"]["Embed"] = github_com_zserge_lorca.Embed
DefaultGojaInterfaces["github.com/zserge/lorca"]["LocateChrome"] = github_com_zserge_lorca.LocateChrome
DefaultGojaInterfaces["github.com/zserge/lorca"]["PDF"] = github_com_zserge_lorca.PDF
DefaultGojaInterfaces["github.com/zserge/lorca"]["PNG"] = github_com_zserge_lorca.PNG
DefaultGojaInterfaces["github.com/zserge/lorca"]["PromptDownload"] = github_com_zserge_lorca.PromptDownload
DefaultGojaInterfaces["github.com/zserge/lorca"]["New"] = github_com_zserge_lorca.New
}
