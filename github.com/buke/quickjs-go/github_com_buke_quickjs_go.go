package quickjs-go
import (
github_com_buke_quickjs_go "github.com/buke/quickjs-go"
)
func Load() {
if _, ok := Api["github.com/buke/quickjs-go"]; !ok {
   Api["github.com/buke/quickjs-go"] = map[string]interface{}{}
}
Api["github.com/buke/quickjs-go"]["Atom"] = github_com_buke_quickjs_go.Atom{}
Api["github.com/buke/quickjs-go"]["Context"] = github_com_buke_quickjs_go.Context{}
Api["github.com/buke/quickjs-go"]["Error"] = github_com_buke_quickjs_go.Error{}
Api["github.com/buke/quickjs-go"]["Loop"] = github_com_buke_quickjs_go.Loop{}
Api["github.com/buke/quickjs-go"]["NewLoop"] = github_com_buke_quickjs_go.NewLoop
Api["github.com/buke/quickjs-go"]["Runtime"] = github_com_buke_quickjs_go.Runtime{}
Api["github.com/buke/quickjs-go"]["NewRuntime"] = github_com_buke_quickjs_go.NewRuntime
Api["github.com/buke/quickjs-go"]["Value"] = github_com_buke_quickjs_go.Value{}
Api["github.com/buke/quickjs-go"]["NewLoop"] = github_com_buke_quickjs_go.NewLoop
Api["github.com/buke/quickjs-go"]["NewRuntime"] = github_com_buke_quickjs_go.NewRuntime
}
