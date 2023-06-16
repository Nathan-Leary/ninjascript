package quickjs-go-polyfill
import (
github_com_buke_quickjs_go_polyfill "github.com/buke/quickjs-go-polyfill"
)
func Load() {
if _, ok := Api["github.com/buke/quickjs-go-polyfill"]; !ok {
   Api["github.com/buke/quickjs-go-polyfill"] = map[string]interface{}{}
}
Api["github.com/buke/quickjs-go-polyfill"]["InjectAll"] = github_com_buke_quickjs_go_polyfill.InjectAll
Api["github.com/buke/quickjs-go-polyfill"]["InjectAll"] = github_com_buke_quickjs_go_polyfill.InjectAll
}
