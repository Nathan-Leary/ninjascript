package list
import (
container_list "container/list"
)
func Load() {
if _, ok := Api["container/list"]; !ok {
   Api["container/list"] = map[string]interface{}{}
}
Api["container/list"]["Element"] = container_list.Element{}
Api["container/list"]["List"] = container_list.List{}
Api["container/list"]["New"] = container_list.New
Api["container/list"]["New"] = container_list.New
}
