package heap
import (
container_heap "container/heap"
)
func Load() {
if _, ok := Api["container/heap"]; !ok {
   Api["container/heap"] = map[string]interface{}{}
}
Api["container/heap"]["Fix"] = container_heap.Fix
Api["container/heap"]["Init"] = container_heap.Init
Api["container/heap"]["Pop"] = container_heap.Pop
Api["container/heap"]["Push"] = container_heap.Push
Api["container/heap"]["Remove"] = container_heap.Remove
Api["container/heap"]["Fix"] = container_heap.Fix
Api["container/heap"]["Init"] = container_heap.Init
Api["container/heap"]["Pop"] = container_heap.Pop
Api["container/heap"]["Push"] = container_heap.Push
Api["container/heap"]["Remove"] = container_heap.Remove
}
