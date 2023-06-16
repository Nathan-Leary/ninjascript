package ninjascript
import (
runtime_pprof "runtime/pprof"
)
func init() {if _, ok := Api["runtime/pprof"]; !ok {
   Api["runtime/pprof"] = map[string]interface{}{}
}
Api["runtime/pprof"]["Do"] = runtime_pprof.Do
Api["runtime/pprof"]["ForLabels"] = runtime_pprof.ForLabels
Api["runtime/pprof"]["Label"] = runtime_pprof.Label
Api["runtime/pprof"]["SetGoroutineLabels"] = runtime_pprof.SetGoroutineLabels
Api["runtime/pprof"]["StartCPUProfile"] = runtime_pprof.StartCPUProfile
Api["runtime/pprof"]["StopCPUProfile"] = runtime_pprof.StopCPUProfile
Api["runtime/pprof"]["WithLabels"] = runtime_pprof.WithLabels
Api["runtime/pprof"]["WriteHeapProfile"] = runtime_pprof.WriteHeapProfile
Api["runtime/pprof"]["LabelSet"] = runtime_pprof.LabelSet{}
Api["runtime/pprof"]["Labels"] = runtime_pprof.Labels
Api["runtime/pprof"]["Profile"] = runtime_pprof.Profile{}
Api["runtime/pprof"]["Lookup"] = runtime_pprof.Lookup
Api["runtime/pprof"]["NewProfile"] = runtime_pprof.NewProfile
Api["runtime/pprof"]["Profiles"] = runtime_pprof.Profiles
Api["runtime/pprof"]["Do"] = runtime_pprof.Do
Api["runtime/pprof"]["ForLabels"] = runtime_pprof.ForLabels
Api["runtime/pprof"]["Label"] = runtime_pprof.Label
Api["runtime/pprof"]["SetGoroutineLabels"] = runtime_pprof.SetGoroutineLabels
Api["runtime/pprof"]["StartCPUProfile"] = runtime_pprof.StartCPUProfile
Api["runtime/pprof"]["StopCPUProfile"] = runtime_pprof.StopCPUProfile
Api["runtime/pprof"]["WithLabels"] = runtime_pprof.WithLabels
Api["runtime/pprof"]["WriteHeapProfile"] = runtime_pprof.WriteHeapProfile
Api["runtime/pprof"]["Labels"] = runtime_pprof.Labels
Api["runtime/pprof"]["Lookup"] = runtime_pprof.Lookup
Api["runtime/pprof"]["NewProfile"] = runtime_pprof.NewProfile
Api["runtime/pprof"]["Profiles"] = runtime_pprof.Profiles

}