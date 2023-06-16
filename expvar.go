package ninjascript
import (
expvar "expvar"
)
func init() {if _, ok := Api["expvar"]; !ok {
   Api["expvar"] = map[string]interface{}{}
}
Api["expvar"]["Do"] = expvar.Do
Api["expvar"]["Handler"] = expvar.Handler
Api["expvar"]["Publish"] = expvar.Publish
Api["expvar"]["Float"] = expvar.Float{}
Api["expvar"]["NewFloat"] = expvar.NewFloat
Api["expvar"]["Int"] = expvar.Int{}
Api["expvar"]["NewInt"] = expvar.NewInt
Api["expvar"]["KeyValue"] = expvar.KeyValue{}
Api["expvar"]["Map"] = expvar.Map{}
Api["expvar"]["NewMap"] = expvar.NewMap
Api["expvar"]["String"] = expvar.String{}
Api["expvar"]["NewString"] = expvar.NewString
Api["expvar"]["Get"] = expvar.Get
Api["expvar"]["Do"] = expvar.Do
Api["expvar"]["Handler"] = expvar.Handler
Api["expvar"]["Publish"] = expvar.Publish
Api["expvar"]["NewFloat"] = expvar.NewFloat
Api["expvar"]["NewInt"] = expvar.NewInt
Api["expvar"]["NewMap"] = expvar.NewMap
Api["expvar"]["NewString"] = expvar.NewString
Api["expvar"]["Get"] = expvar.Get

}