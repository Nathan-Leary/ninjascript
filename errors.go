package ninjascript
import (
errors "errors"
)
func init() {if _, ok := Api["errors"]; !ok {
   Api["errors"] = map[string]interface{}{}
}
Api["errors"]["As"] = errors.As
Api["errors"]["Is"] = errors.Is
Api["errors"]["Join"] = errors.Join
Api["errors"]["New"] = errors.New
Api["errors"]["Unwrap"] = errors.Unwrap
Api["errors"]["As"] = errors.As
Api["errors"]["Is"] = errors.Is
Api["errors"]["Join"] = errors.Join
Api["errors"]["New"] = errors.New
Api["errors"]["Unwrap"] = errors.Unwrap

}