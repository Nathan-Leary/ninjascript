# ninjascript
 
Goja (Javascript ECMA5 Engine) with almost all of the Golang Api installed.

example code:
```
// ninjascripttest project main.go
package main

import (
	"github.com/Nathan-Leary/ninjascript"
	"github.com/Nathan-Leary/ninjastars/github.com/spf13/cast"
)

func init() {

}

func main() {

	ninjascript.Import(cast.Api)

	vm := ninjascript.New()

	vm.RunString(`
	var number = 123;
	fmt.Println(number, "=>", typeof github.com.spf13.cast.ToString(number))
	`)

}

// outputs >> 123 => string
```
more info coming...
