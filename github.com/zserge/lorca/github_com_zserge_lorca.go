package lorca
import (
github_com_zserge_lorca "github.com/zserge/lorca"
)
func Load() {
if _, ok := Api["github.com/zserge/lorca"]; !ok {
   Api["github.com/zserge/lorca"] = map[string]interface{}{}
}
Api["github.com/zserge/lorca"]["PageA4Width"] = github_com_zserge_lorca.PageA4Width
Api["github.com/zserge/lorca"]["PageA4Height"] = github_com_zserge_lorca.PageA4Height
Api["github.com/zserge/lorca"]["ChromeExecutable"] = github_com_zserge_lorca.ChromeExecutable
Api["github.com/zserge/lorca"]["Embed"] = github_com_zserge_lorca.Embed
Api["github.com/zserge/lorca"]["LocateChrome"] = github_com_zserge_lorca.LocateChrome
Api["github.com/zserge/lorca"]["PDF"] = github_com_zserge_lorca.PDF
Api["github.com/zserge/lorca"]["PNG"] = github_com_zserge_lorca.PNG
Api["github.com/zserge/lorca"]["PromptDownload"] = github_com_zserge_lorca.PromptDownload
Api["github.com/zserge/lorca"]["Bounds"] = github_com_zserge_lorca.Bounds{}
Api["github.com/zserge/lorca"]["UI"] = new(github_com_zserge_lorca.UI)
Api["github.com/zserge/lorca"]["New"] = github_com_zserge_lorca.New
Api["github.com/zserge/lorca"]["WindowStateNormal"] = github_com_zserge_lorca.WindowStateNormal
Api["github.com/zserge/lorca"]["WindowStateMaximized"] = github_com_zserge_lorca.WindowStateMaximized
Api["github.com/zserge/lorca"]["WindowStateMinimized"] = github_com_zserge_lorca.WindowStateMinimized
Api["github.com/zserge/lorca"]["WindowStateFullscreen"] = github_com_zserge_lorca.WindowStateFullscreen
Api["github.com/zserge/lorca"]["Embed"] = github_com_zserge_lorca.Embed
Api["github.com/zserge/lorca"]["LocateChrome"] = github_com_zserge_lorca.LocateChrome
Api["github.com/zserge/lorca"]["PDF"] = github_com_zserge_lorca.PDF
Api["github.com/zserge/lorca"]["PNG"] = github_com_zserge_lorca.PNG
Api["github.com/zserge/lorca"]["PromptDownload"] = github_com_zserge_lorca.PromptDownload
Api["github.com/zserge/lorca"]["New"] = github_com_zserge_lorca.New
}
