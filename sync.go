package ninjascript
import (
sync "sync"
)
func init() {if _, ok := Api["sync"]; !ok {
   Api["sync"] = map[string]interface{}{}
}
Api["sync"]["Cond"] = sync.Cond{}
Api["sync"]["NewCond"] = sync.NewCond
Api["sync"]["Map"] = sync.Map{}
Api["sync"]["Mutex"] = sync.Mutex{}
Api["sync"]["Once"] = sync.Once{}
Api["sync"]["Pool"] = sync.Pool{}
Api["sync"]["RWMutex"] = sync.RWMutex{}
Api["sync"]["WaitGroup"] = sync.WaitGroup{}
Api["sync"]["NewCond"] = sync.NewCond

}