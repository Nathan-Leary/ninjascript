package ninjascript

import (
	runtime "runtime"
)

func init() {
	if _, ok := Api["runtime"]; !ok {
		Api["runtime"] = map[string]interface{}{}
	}
	Api["runtime"]["MemProfileRate"] = runtime.MemProfileRate
	Api["runtime"]["BlockProfile"] = runtime.BlockProfile
	Api["runtime"]["Breakpoint"] = runtime.Breakpoint
	Api["runtime"]["CPUProfile"] = runtime.CPUProfile
	Api["runtime"]["Caller"] = runtime.Caller
	Api["runtime"]["Callers"] = runtime.Callers
	Api["runtime"]["GC"] = runtime.GC
	Api["runtime"]["GOMAXPROCS"] = runtime.GOMAXPROCS
	Api["runtime"]["GOROOT"] = runtime.GOROOT
	Api["runtime"]["Goexit"] = runtime.Goexit
	Api["runtime"]["GoroutineProfile"] = runtime.GoroutineProfile
	Api["runtime"]["Gosched"] = runtime.Gosched
	Api["runtime"]["KeepAlive"] = runtime.KeepAlive
	// Api["runtime"]["File"] = runtime.File{}
	Api["runtime"]["LockOSThread"] = runtime.LockOSThread
	Api["runtime"]["MemProfile"] = runtime.MemProfile
	Api["runtime"]["MutexProfile"] = runtime.MutexProfile
	Api["runtime"]["NumCPU"] = runtime.NumCPU
	Api["runtime"]["NumCgoCall"] = runtime.NumCgoCall
	Api["runtime"]["NumGoroutine"] = runtime.NumGoroutine
	Api["runtime"]["ReadMemStats"] = runtime.ReadMemStats
	Api["runtime"]["ReadTrace"] = runtime.ReadTrace
	Api["runtime"]["SetBlockProfileRate"] = runtime.SetBlockProfileRate
	Api["runtime"]["SetCPUProfileRate"] = runtime.SetCPUProfileRate
	Api["runtime"]["SetCgoTraceback"] = runtime.SetCgoTraceback
	Api["runtime"]["SetFinalizer"] = runtime.SetFinalizer
	Api["runtime"]["SetMutexProfileFraction"] = runtime.SetMutexProfileFraction
	Api["runtime"]["Stack"] = runtime.Stack
	Api["runtime"]["StartTrace"] = runtime.StartTrace
	Api["runtime"]["StopTrace"] = runtime.StopTrace
	Api["runtime"]["ThreadCreateProfile"] = runtime.ThreadCreateProfile
	Api["runtime"]["UnlockOSThread"] = runtime.UnlockOSThread
	Api["runtime"]["Version"] = runtime.Version
	Api["runtime"]["BlockProfileRecord"] = runtime.BlockProfileRecord{}
	Api["runtime"]["Frame"] = runtime.Frame{}
	Api["runtime"]["Frames"] = runtime.Frames{}
	Api["runtime"]["CallersFrames"] = runtime.CallersFrames
	Api["runtime"]["Func"] = runtime.Func{}
	Api["runtime"]["FuncForPC"] = runtime.FuncForPC
	Api["runtime"]["MemProfileRecord"] = runtime.MemProfileRecord{}
	Api["runtime"]["MemStats"] = runtime.MemStats{}
	Api["runtime"]["StackRecord"] = runtime.StackRecord{}
	Api["runtime"]["TypeAssertionError"] = runtime.TypeAssertionError{}
	Api["runtime"]["BlockProfile"] = runtime.BlockProfile
	Api["runtime"]["Breakpoint"] = runtime.Breakpoint
	Api["runtime"]["CPUProfile"] = runtime.CPUProfile
	Api["runtime"]["Caller"] = runtime.Caller
	Api["runtime"]["Callers"] = runtime.Callers
	Api["runtime"]["GC"] = runtime.GC
	Api["runtime"]["GOMAXPROCS"] = runtime.GOMAXPROCS
	Api["runtime"]["GOROOT"] = runtime.GOROOT
	Api["runtime"]["Goexit"] = runtime.Goexit
	Api["runtime"]["GoroutineProfile"] = runtime.GoroutineProfile
	Api["runtime"]["Gosched"] = runtime.Gosched
	Api["runtime"]["KeepAlive"] = runtime.KeepAlive
	Api["runtime"]["LockOSThread"] = runtime.LockOSThread
	Api["runtime"]["MemProfile"] = runtime.MemProfile
	Api["runtime"]["MutexProfile"] = runtime.MutexProfile
	Api["runtime"]["NumCPU"] = runtime.NumCPU
	Api["runtime"]["NumCgoCall"] = runtime.NumCgoCall
	Api["runtime"]["NumGoroutine"] = runtime.NumGoroutine
	Api["runtime"]["ReadMemStats"] = runtime.ReadMemStats
	Api["runtime"]["ReadTrace"] = runtime.ReadTrace
	Api["runtime"]["SetBlockProfileRate"] = runtime.SetBlockProfileRate
	Api["runtime"]["SetCPUProfileRate"] = runtime.SetCPUProfileRate
	Api["runtime"]["SetCgoTraceback"] = runtime.SetCgoTraceback
	Api["runtime"]["SetFinalizer"] = runtime.SetFinalizer
	Api["runtime"]["SetMutexProfileFraction"] = runtime.SetMutexProfileFraction
	Api["runtime"]["Stack"] = runtime.Stack
	Api["runtime"]["StartTrace"] = runtime.StartTrace
	Api["runtime"]["StopTrace"] = runtime.StopTrace
	Api["runtime"]["ThreadCreateProfile"] = runtime.ThreadCreateProfile
	Api["runtime"]["UnlockOSThread"] = runtime.UnlockOSThread
	Api["runtime"]["Version"] = runtime.Version
	Api["runtime"]["CallersFrames"] = runtime.CallersFrames
	Api["runtime"]["FuncForPC"] = runtime.FuncForPC

}
