package ninjascript
import (
flag "flag"
)
func init() {if _, ok := Api["flag"]; !ok {
   Api["flag"] = map[string]interface{}{}
}
Api["flag"]["CommandLine"] = flag.CommandLine
Api["flag"]["ErrHelp"] = flag.ErrHelp
Api["flag"]["Usage"] = flag.Usage
Api["flag"]["Arg"] = flag.Arg
Api["flag"]["Args"] = flag.Args
Api["flag"]["Bool"] = flag.Bool
Api["flag"]["BoolVar"] = flag.BoolVar
Api["flag"]["Duration"] = flag.Duration
Api["flag"]["DurationVar"] = flag.DurationVar
Api["flag"]["Float64"] = flag.Float64
Api["flag"]["Float64Var"] = flag.Float64Var
Api["flag"]["Func"] = flag.Func
Api["flag"]["Int"] = flag.Int
Api["flag"]["Int64"] = flag.Int64
Api["flag"]["Int64Var"] = flag.Int64Var
Api["flag"]["IntVar"] = flag.IntVar
Api["flag"]["NArg"] = flag.NArg
Api["flag"]["NFlag"] = flag.NFlag
Api["flag"]["Parse"] = flag.Parse
Api["flag"]["Parsed"] = flag.Parsed
Api["flag"]["PrintDefaults"] = flag.PrintDefaults
Api["flag"]["Set"] = flag.Set
Api["flag"]["String"] = flag.String
Api["flag"]["StringVar"] = flag.StringVar
Api["flag"]["TextVar"] = flag.TextVar
Api["flag"]["Uint"] = flag.Uint
Api["flag"]["Uint64"] = flag.Uint64
Api["flag"]["Uint64Var"] = flag.Uint64Var
Api["flag"]["UintVar"] = flag.UintVar
Api["flag"]["UnquoteUsage"] = flag.UnquoteUsage
Api["flag"]["Var"] = flag.Var
Api["flag"]["Visit"] = flag.Visit
Api["flag"]["VisitAll"] = flag.VisitAll
Api["flag"]["ContinueOnError"] = flag.ContinueOnError
Api["flag"]["ExitOnError"] = flag.ExitOnError
Api["flag"]["PanicOnError"] = flag.PanicOnError
Api["flag"]["Flag"] = flag.Flag{}
Api["flag"]["Lookup"] = flag.Lookup
Api["flag"]["FlagSet"] = flag.FlagSet{}
Api["flag"]["NewFlagSet"] = flag.NewFlagSet
Api["flag"]["Arg"] = flag.Arg
Api["flag"]["Args"] = flag.Args
Api["flag"]["Bool"] = flag.Bool
Api["flag"]["BoolVar"] = flag.BoolVar
Api["flag"]["Duration"] = flag.Duration
Api["flag"]["DurationVar"] = flag.DurationVar
Api["flag"]["Float64"] = flag.Float64
Api["flag"]["Float64Var"] = flag.Float64Var
Api["flag"]["Func"] = flag.Func
Api["flag"]["Int"] = flag.Int
Api["flag"]["Int64"] = flag.Int64
Api["flag"]["Int64Var"] = flag.Int64Var
Api["flag"]["IntVar"] = flag.IntVar
Api["flag"]["NArg"] = flag.NArg
Api["flag"]["NFlag"] = flag.NFlag
Api["flag"]["Parse"] = flag.Parse
Api["flag"]["Parsed"] = flag.Parsed
Api["flag"]["PrintDefaults"] = flag.PrintDefaults
Api["flag"]["Set"] = flag.Set
Api["flag"]["String"] = flag.String
Api["flag"]["StringVar"] = flag.StringVar
Api["flag"]["TextVar"] = flag.TextVar
Api["flag"]["Uint"] = flag.Uint
Api["flag"]["Uint64"] = flag.Uint64
Api["flag"]["Uint64Var"] = flag.Uint64Var
Api["flag"]["UintVar"] = flag.UintVar
Api["flag"]["UnquoteUsage"] = flag.UnquoteUsage
Api["flag"]["Var"] = flag.Var
Api["flag"]["Visit"] = flag.Visit
Api["flag"]["VisitAll"] = flag.VisitAll
Api["flag"]["Lookup"] = flag.Lookup
Api["flag"]["NewFlagSet"] = flag.NewFlagSet

}