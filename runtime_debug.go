package ninjascript
import (
runtime_debug "runtime/debug"
)
func init() {if _, ok := Api["runtime/debug"]; !ok {
   Api["runtime/debug"] = map[string]interface{}{}
}
Api["runtime/debug"]["FreeOSMemory"] = runtime_debug.FreeOSMemory
Api["runtime/debug"]["PrintStack"] = runtime_debug.PrintStack
Api["runtime/debug"]["ReadGCStats"] = runtime_debug.ReadGCStats
Api["runtime/debug"]["SetGCPercent"] = runtime_debug.SetGCPercent
Api["runtime/debug"]["SetMaxStack"] = runtime_debug.SetMaxStack
Api["runtime/debug"]["SetMaxThreads"] = runtime_debug.SetMaxThreads
Api["runtime/debug"]["SetMemoryLimit"] = runtime_debug.SetMemoryLimit
Api["runtime/debug"]["SetPanicOnFault"] = runtime_debug.SetPanicOnFault
Api["runtime/debug"]["SetTraceback"] = runtime_debug.SetTraceback
Api["runtime/debug"]["Stack"] = runtime_debug.Stack
Api["runtime/debug"]["WriteHeapDump"] = runtime_debug.WriteHeapDump
Api["runtime/debug"]["BuildInfo"] = runtime_debug.BuildInfo{}
Api["runtime/debug"]["ParseBuildInfo"] = runtime_debug.ParseBuildInfo
Api["runtime/debug"]["ReadBuildInfo"] = runtime_debug.ReadBuildInfo
Api["runtime/debug"]["BuildSetting"] = runtime_debug.BuildSetting{}
Api["runtime/debug"]["GCStats"] = runtime_debug.GCStats{}
Api["runtime/debug"]["Module"] = runtime_debug.Module{}
Api["runtime/debug"]["FreeOSMemory"] = runtime_debug.FreeOSMemory
Api["runtime/debug"]["PrintStack"] = runtime_debug.PrintStack
Api["runtime/debug"]["ReadGCStats"] = runtime_debug.ReadGCStats
Api["runtime/debug"]["SetGCPercent"] = runtime_debug.SetGCPercent
Api["runtime/debug"]["SetMaxStack"] = runtime_debug.SetMaxStack
Api["runtime/debug"]["SetMaxThreads"] = runtime_debug.SetMaxThreads
Api["runtime/debug"]["SetMemoryLimit"] = runtime_debug.SetMemoryLimit
Api["runtime/debug"]["SetPanicOnFault"] = runtime_debug.SetPanicOnFault
Api["runtime/debug"]["SetTraceback"] = runtime_debug.SetTraceback
Api["runtime/debug"]["Stack"] = runtime_debug.Stack
Api["runtime/debug"]["WriteHeapDump"] = runtime_debug.WriteHeapDump
Api["runtime/debug"]["ParseBuildInfo"] = runtime_debug.ParseBuildInfo
Api["runtime/debug"]["ReadBuildInfo"] = runtime_debug.ReadBuildInfo

}