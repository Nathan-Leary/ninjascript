package ninjascript
import (
testing_quick "testing/quick"
)
func init() {if _, ok := Api["testing/quick"]; !ok {
   Api["testing/quick"] = map[string]interface{}{}
}
Api["testing/quick"]["Check"] = testing_quick.Check
Api["testing/quick"]["CheckEqual"] = testing_quick.CheckEqual
Api["testing/quick"]["Value"] = testing_quick.Value
Api["testing/quick"]["CheckEqualError"] = testing_quick.CheckEqualError{}
Api["testing/quick"]["CheckError"] = testing_quick.CheckError{}
Api["testing/quick"]["Config"] = testing_quick.Config{}
Api["testing/quick"]["Check"] = testing_quick.Check
Api["testing/quick"]["CheckEqual"] = testing_quick.CheckEqual
Api["testing/quick"]["Value"] = testing_quick.Value

}