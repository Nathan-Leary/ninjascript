package ring
import (
container_ring "container/ring"
)
func Load() {
if _, ok := Api["container/ring"]; !ok {
   Api["container/ring"] = map[string]interface{}{}
}
Api["container/ring"]["Ring"] = container_ring.Ring{}
Api["container/ring"]["New"] = container_ring.New
Api["container/ring"]["New"] = container_ring.New
}
