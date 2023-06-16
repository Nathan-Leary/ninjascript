package ninjascript
import (
testing_fstest "testing/fstest"
)
func init() {if _, ok := Api["testing/fstest"]; !ok {
   Api["testing/fstest"] = map[string]interface{}{}
}
Api["testing/fstest"]["TestFS"] = testing_fstest.TestFS
Api["testing/fstest"]["MapFile"] = testing_fstest.MapFile{}
Api["testing/fstest"]["TestFS"] = testing_fstest.TestFS

}