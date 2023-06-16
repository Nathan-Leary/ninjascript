package ninjascript
import (
runtime_coverage "runtime/coverage"
)
func init() {if _, ok := Api["runtime/coverage"]; !ok {
   Api["runtime/coverage"] = map[string]interface{}{}
}
Api["runtime/coverage"]["ClearCounters"] = runtime_coverage.ClearCounters
Api["runtime/coverage"]["WriteCounters"] = runtime_coverage.WriteCounters
Api["runtime/coverage"]["WriteCountersDir"] = runtime_coverage.WriteCountersDir
Api["runtime/coverage"]["WriteMeta"] = runtime_coverage.WriteMeta
Api["runtime/coverage"]["WriteMetaDir"] = runtime_coverage.WriteMetaDir
Api["runtime/coverage"]["ClearCounters"] = runtime_coverage.ClearCounters
Api["runtime/coverage"]["WriteCounters"] = runtime_coverage.WriteCounters
Api["runtime/coverage"]["WriteCountersDir"] = runtime_coverage.WriteCountersDir
Api["runtime/coverage"]["WriteMeta"] = runtime_coverage.WriteMeta
Api["runtime/coverage"]["WriteMetaDir"] = runtime_coverage.WriteMetaDir

}