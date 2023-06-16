package ninjascript

import (
	testing "testing"
)

func init() {
	if _, ok := Api["testing"]; !ok {
		Api["testing"] = map[string]interface{}{}
	}
	// Api["testing"]["TestXxx"] = testing.TestXxx
	// Api["testing"]["BenchmarkXxx"] = testing.BenchmarkXxx
	// Api["testing"]["FuzzXxx"] = testing.FuzzXxx
	// Api["testing"]["FuzzJSONMarshaling"] = testing.FuzzJSONMarshaling
	// Api["testing"]["TestMain"] = testing.TestMain
	Api["testing"]["AllocsPerRun"] = testing.AllocsPerRun
	Api["testing"]["CoverMode"] = testing.CoverMode
	Api["testing"]["Coverage"] = testing.Coverage
	Api["testing"]["Init"] = testing.Init
	Api["testing"]["Main"] = testing.Main
	Api["testing"]["RegisterCover"] = testing.RegisterCover
	Api["testing"]["RunBenchmarks"] = testing.RunBenchmarks
	Api["testing"]["RunExamples"] = testing.RunExamples
	Api["testing"]["RunTests"] = testing.RunTests
	Api["testing"]["Short"] = testing.Short
	Api["testing"]["Verbose"] = testing.Verbose
	Api["testing"]["B"] = testing.B{}
	Api["testing"]["BenchmarkResult"] = testing.BenchmarkResult{}
	Api["testing"]["Benchmark"] = testing.Benchmark
	Api["testing"]["Cover"] = testing.Cover{}
	Api["testing"]["CoverBlock"] = testing.CoverBlock{}
	Api["testing"]["F"] = testing.F{}
	Api["testing"]["InternalBenchmark"] = testing.InternalBenchmark{}
	Api["testing"]["InternalExample"] = testing.InternalExample{}
	Api["testing"]["InternalFuzzTarget"] = testing.InternalFuzzTarget{}
	Api["testing"]["InternalTest"] = testing.InternalTest{}
	Api["testing"]["M"] = testing.M{}
	Api["testing"]["MainStart"] = testing.MainStart
	Api["testing"]["PB"] = testing.PB{}
	Api["testing"]["T"] = testing.T{}
	Api["testing"]["AllocsPerRun"] = testing.AllocsPerRun
	Api["testing"]["CoverMode"] = testing.CoverMode
	Api["testing"]["Coverage"] = testing.Coverage
	Api["testing"]["Init"] = testing.Init
	Api["testing"]["Main"] = testing.Main
	Api["testing"]["RegisterCover"] = testing.RegisterCover
	Api["testing"]["RunBenchmarks"] = testing.RunBenchmarks
	Api["testing"]["RunExamples"] = testing.RunExamples
	Api["testing"]["RunTests"] = testing.RunTests
	Api["testing"]["Short"] = testing.Short
	Api["testing"]["Verbose"] = testing.Verbose
	Api["testing"]["Benchmark"] = testing.Benchmark
	Api["testing"]["MainStart"] = testing.MainStart

}
