package ninjascript
import (
io_ioutil "io/ioutil"
)
func init() {if _, ok := Api["io/ioutil"]; !ok {
   Api["io/ioutil"] = map[string]interface{}{}
}
Api["io/ioutil"]["Discard"] = io_ioutil.Discard
Api["io/ioutil"]["NopCloser"] = io_ioutil.NopCloser
Api["io/ioutil"]["ReadAll"] = io_ioutil.ReadAll
Api["io/ioutil"]["ReadDir"] = io_ioutil.ReadDir
Api["io/ioutil"]["ReadFile"] = io_ioutil.ReadFile
Api["io/ioutil"]["TempDir"] = io_ioutil.TempDir
Api["io/ioutil"]["TempFile"] = io_ioutil.TempFile
Api["io/ioutil"]["WriteFile"] = io_ioutil.WriteFile
Api["io/ioutil"]["NopCloser"] = io_ioutil.NopCloser
Api["io/ioutil"]["ReadAll"] = io_ioutil.ReadAll
Api["io/ioutil"]["ReadDir"] = io_ioutil.ReadDir
Api["io/ioutil"]["ReadFile"] = io_ioutil.ReadFile
Api["io/ioutil"]["TempDir"] = io_ioutil.TempDir
Api["io/ioutil"]["TempFile"] = io_ioutil.TempFile
Api["io/ioutil"]["WriteFile"] = io_ioutil.WriteFile

}