package ninjascript

import (
	"fmt"
	"strings"

	"github.com/dop251/goja"
)

var Api map[string]map[string]interface{} = map[string]map[string]interface{}{}

func init() {

}

var Vm *goja.Runtime

var uniqueMap = map[string]bool{}

func Import(ImportedApi map[string]map[string]interface{}) {
	for k, v := range ImportedApi {
		Api[k] = v
	}
}

func New() *goja.Runtime {

	Vm = goja.New()

	script := ""
	for k, v := range Api {

		nameSplit := strings.Split(strings.ReplaceAll(k, ".", "/"), "/")

		concatName := ""
		for x := 1; x <= len(nameSplit); x++ {
			key := strings.Join(nameSplit[:x], "/")
			if x == 1 {
				concatName = nameSplit[x-1]
			} else {
				concatName += "." + nameSplit[x-1]
			}
			concatName = strings.ReplaceAll(concatName, "-", "_")

			if _, ok := uniqueMap[key]; !ok {
				script += (concatName + ` = {};`)
				uniqueMap[key] = true
			}

		}

		name := strings.ReplaceAll(strings.Join(nameSplit, "_"), "-", "_")

		for k2, v2 := range v {
			Vm.Set("__"+name+"_"+k2, v2)
			script += concatName + "." + k2 + " = " + "__" + name + "_" + k2 + ";"
		}

	}

	_, err := Vm.RunString(script)

	if err != nil {
		fmt.Println(err)
	}

	return Vm

}
