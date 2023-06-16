package ninjascript
import (
plugin "plugin"
)
func init() {if _, ok := Api["plugin"]; !ok {
   Api["plugin"] = map[string]interface{}{}
}
Api["plugin"]["Plugin"] = plugin.Plugin{}
Api["plugin"]["Open"] = plugin.Open
Api["plugin"]["Open"] = plugin.Open

}