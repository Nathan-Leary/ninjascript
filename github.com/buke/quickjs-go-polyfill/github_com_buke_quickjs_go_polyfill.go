package quickjs-go-polyfill
import (
github_com_buke_quickjs_go_polyfill "github.com/buke/quickjs-go-polyfill"
)
func Load() {
if _, ok := DefaultGojaInterfaces["github.com/buke/quickjs-go-polyfill"]; !ok {
   DefaultGojaInterfaces["github.com/buke/quickjs-go-polyfill"] = map[string]interface{}{}
}
DefaultGojaInterfaces["github.com/buke/quickjs-go-polyfill"]["InjectAll"] = github_com_buke_quickjs_go_polyfill.InjectAll
DefaultGojaInterfaces["github.com/buke/quickjs-go-polyfill"]["InjectAll"] = github_com_buke_quickjs_go_polyfill.InjectAll
}
