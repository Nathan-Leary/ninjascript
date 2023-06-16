package ninjascript

import(
"github.com/dop251/goja"
)

var Api = map[string]map[string]interface{}{}

func New() *goja.Runtime {
  var vm *goja.Runtime
  return goja.New()
}

func Install() {
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

	b, _ := ioutil.ReadFile(filepath.Dir(os.Args[0]) + "/init.js")
	script += "\n" + string(b)

	_, err := Vm.RunString(script)

	if err != nil {
		fmt.Println(err)
	}

}
