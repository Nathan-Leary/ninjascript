package ninjascript
import (
runtime_metrics "runtime/metrics"
)
func init() {if _, ok := Api["runtime/metrics"]; !ok {
   Api["runtime/metrics"] = map[string]interface{}{}
}
Api["runtime/metrics"]["Read"] = runtime_metrics.Read
Api["runtime/metrics"]["Description"] = runtime_metrics.Description{}
Api["runtime/metrics"]["All"] = runtime_metrics.All
Api["runtime/metrics"]["Float64Histogram"] = runtime_metrics.Float64Histogram{}
Api["runtime/metrics"]["Sample"] = runtime_metrics.Sample{}
Api["runtime/metrics"]["Value"] = runtime_metrics.Value{}
Api["runtime/metrics"]["KindBad"] = runtime_metrics.KindBad
Api["runtime/metrics"]["KindUint64"] = runtime_metrics.KindUint64
Api["runtime/metrics"]["KindFloat64"] = runtime_metrics.KindFloat64
Api["runtime/metrics"]["KindFloat64Histogram"] = runtime_metrics.KindFloat64Histogram
Api["runtime/metrics"]["Read"] = runtime_metrics.Read
Api["runtime/metrics"]["All"] = runtime_metrics.All

}