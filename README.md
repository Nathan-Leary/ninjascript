# ninjascript

Goja with almost all of the Golang Api installed.

exanple code:

package main

import (
	"github.com/Nathan-Leary/ninjascript"
)

func main() {
	vm := ninjascript.New()
	vm.RunString(`fmt.Println(123);`)
}

// outputs >> 123

more info coming...
