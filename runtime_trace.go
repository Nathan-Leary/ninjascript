package ninjascript
import (
runtime_trace "runtime/trace"
)
func init() {if _, ok := Api["runtime/trace"]; !ok {
   Api["runtime/trace"] = map[string]interface{}{}
}
Api["runtime/trace"]["IsEnabled"] = runtime_trace.IsEnabled
Api["runtime/trace"]["Log"] = runtime_trace.Log
Api["runtime/trace"]["Logf"] = runtime_trace.Logf
Api["runtime/trace"]["Start"] = runtime_trace.Start
Api["runtime/trace"]["Stop"] = runtime_trace.Stop
Api["runtime/trace"]["WithRegion"] = runtime_trace.WithRegion
Api["runtime/trace"]["Region"] = runtime_trace.Region{}
Api["runtime/trace"]["StartRegion"] = runtime_trace.StartRegion
Api["runtime/trace"]["Task"] = runtime_trace.Task{}
Api["runtime/trace"]["NewTask"] = runtime_trace.NewTask
Api["runtime/trace"]["IsEnabled"] = runtime_trace.IsEnabled
Api["runtime/trace"]["Log"] = runtime_trace.Log
Api["runtime/trace"]["Logf"] = runtime_trace.Logf
Api["runtime/trace"]["Start"] = runtime_trace.Start
Api["runtime/trace"]["Stop"] = runtime_trace.Stop
Api["runtime/trace"]["WithRegion"] = runtime_trace.WithRegion
Api["runtime/trace"]["StartRegion"] = runtime_trace.StartRegion
Api["runtime/trace"]["NewTask"] = runtime_trace.NewTask

}