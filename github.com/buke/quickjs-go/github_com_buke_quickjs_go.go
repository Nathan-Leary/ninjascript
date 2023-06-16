package quickjs-go
import (
github_com_buke_quickjs_go "github.com/buke/quickjs-go"
)
func Load() {
if _, ok := DefaultGojaInterfaces["github.com/buke/quickjs-go"]; !ok {
   DefaultGojaInterfaces["github.com/buke/quickjs-go"] = map[string]interface{}{}
}
DefaultGojaInterfaces["github.com/buke/quickjs-go"]["Atom"] = github_com_buke_quickjs_go.Atom{}
DefaultGojaInterfaces["github.com/buke/quickjs-go"]["Context"] = github_com_buke_quickjs_go.Context{}
DefaultGojaInterfaces["github.com/buke/quickjs-go"]["Error"] = github_com_buke_quickjs_go.Error{}
DefaultGojaInterfaces["github.com/buke/quickjs-go"]["Loop"] = github_com_buke_quickjs_go.Loop{}
DefaultGojaInterfaces["github.com/buke/quickjs-go"]["NewLoop"] = github_com_buke_quickjs_go.NewLoop
DefaultGojaInterfaces["github.com/buke/quickjs-go"]["Runtime"] = github_com_buke_quickjs_go.Runtime{}
DefaultGojaInterfaces["github.com/buke/quickjs-go"]["NewRuntime"] = github_com_buke_quickjs_go.NewRuntime
DefaultGojaInterfaces["github.com/buke/quickjs-go"]["Value"] = github_com_buke_quickjs_go.Value{}
DefaultGojaInterfaces["github.com/buke/quickjs-go"]["NewLoop"] = github_com_buke_quickjs_go.NewLoop
DefaultGojaInterfaces["github.com/buke/quickjs-go"]["NewRuntime"] = github_com_buke_quickjs_go.NewRuntime
}
