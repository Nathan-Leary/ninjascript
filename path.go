package ninjascript
import (
path "path"
)
func init() {if _, ok := Api["path"]; !ok {
   Api["path"] = map[string]interface{}{}
}
Api["path"]["ErrBadPattern"] = path.ErrBadPattern
Api["path"]["Base"] = path.Base
Api["path"]["Clean"] = path.Clean
Api["path"]["Dir"] = path.Dir
Api["path"]["Ext"] = path.Ext
Api["path"]["IsAbs"] = path.IsAbs
Api["path"]["Join"] = path.Join
Api["path"]["Match"] = path.Match
Api["path"]["Split"] = path.Split
Api["path"]["Base"] = path.Base
Api["path"]["Clean"] = path.Clean
Api["path"]["Dir"] = path.Dir
Api["path"]["Ext"] = path.Ext
Api["path"]["IsAbs"] = path.IsAbs
Api["path"]["Join"] = path.Join
Api["path"]["Match"] = path.Match
Api["path"]["Split"] = path.Split

}