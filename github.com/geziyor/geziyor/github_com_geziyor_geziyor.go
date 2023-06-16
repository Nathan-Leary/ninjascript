package geziyor
import (
github_com_geziyor_geziyor "github.com/geziyor/geziyor"
)
func Load() {
if _, ok := DefaultGojaInterfaces["github.com/geziyor/geziyor"]; !ok {
   DefaultGojaInterfaces["github.com/geziyor/geziyor"] = map[string]interface{}{}
}
DefaultGojaInterfaces["github.com/geziyor/geziyor"]["Geziyor"] = github_com_geziyor_geziyor.Geziyor{}
DefaultGojaInterfaces["github.com/geziyor/geziyor"]["NewGeziyor"] = github_com_geziyor_geziyor.NewGeziyor
DefaultGojaInterfaces["github.com/geziyor/geziyor"]["Options"] = github_com_geziyor_geziyor.Options{}
DefaultGojaInterfaces["github.com/geziyor/geziyor"]["NewGeziyor"] = github_com_geziyor_geziyor.NewGeziyor
}
