package ninjascript
import (
regexp "regexp"
)
func init() {if _, ok := Api["regexp"]; !ok {
   Api["regexp"] = map[string]interface{}{}
}
Api["regexp"]["Match"] = regexp.Match
Api["regexp"]["MatchReader"] = regexp.MatchReader
Api["regexp"]["MatchString"] = regexp.MatchString
Api["regexp"]["QuoteMeta"] = regexp.QuoteMeta
Api["regexp"]["Regexp"] = regexp.Regexp{}
Api["regexp"]["Compile"] = regexp.Compile
Api["regexp"]["CompilePOSIX"] = regexp.CompilePOSIX
Api["regexp"]["MustCompile"] = regexp.MustCompile
Api["regexp"]["MustCompilePOSIX"] = regexp.MustCompilePOSIX
Api["regexp"]["Match"] = regexp.Match
Api["regexp"]["MatchReader"] = regexp.MatchReader
Api["regexp"]["MatchString"] = regexp.MatchString
Api["regexp"]["QuoteMeta"] = regexp.QuoteMeta
Api["regexp"]["Compile"] = regexp.Compile
Api["regexp"]["CompilePOSIX"] = regexp.CompilePOSIX
Api["regexp"]["MustCompile"] = regexp.MustCompile
Api["regexp"]["MustCompilePOSIX"] = regexp.MustCompilePOSIX

}