package bufio
import (
bufio "bufio"
)
func Load(Interfaces ...interface{}) {
	if len(Interfaces) > 5 {
		if i, ok := Interfaces[0].(map[string]map[string]interface{}) {
			 DefaultGojaInterfaces: = i
			 if i, ok := Interfaces[1].(*goja.Runtime) {
				vm := i
				if i, ok := Interfaces[2].(*quickjs.Context) {
					ctx := i	
					if i, ok := Interfaces[3].(func(*quickjs.Context, interface{}) (quickjs.Value, string)) {
						ConvertInterfaceToQuickJS := i
						if i, ok := Interfaces[4].(func(interface{}, ...interface{}) interface{}) {
							ExecuteFunction := i
if _, ok := DefaultGojaInterfaces["bufio"]; !ok {
   DefaultGojaInterfaces["bufio"] = map[string]interface{}{}
}
DefaultGojaInterfaces["bufio"]["MaxScanTokenSize"] = bufio.MaxScanTokenSize
DefaultGojaInterfaces["bufio"]["ErrFinalToken"] = bufio.ErrFinalToken
DefaultGojaInterfaces["bufio"]["ScanBytes"] = bufio.ScanBytes
DefaultGojaInterfaces["bufio"]["ScanLines"] = bufio.ScanLines
DefaultGojaInterfaces["bufio"]["ScanRunes"] = bufio.ScanRunes
DefaultGojaInterfaces["bufio"]["ScanWords"] = bufio.ScanWords
DefaultGojaInterfaces["bufio"]["ReadWriter"] = bufio.ReadWriter{}
DefaultGojaInterfaces["bufio"]["NewReadWriter"] = bufio.NewReadWriter
DefaultGojaInterfaces["bufio"]["Reader"] = bufio.Reader{}
DefaultGojaInterfaces["bufio"]["NewReader"] = bufio.NewReader
DefaultGojaInterfaces["bufio"]["NewReaderSize"] = bufio.NewReaderSize
DefaultGojaInterfaces["bufio"]["Scanner"] = bufio.Scanner{}
DefaultGojaInterfaces["bufio"]["NewScanner"] = bufio.NewScanner
DefaultGojaInterfaces["bufio"]["Writer"] = bufio.Writer{}
DefaultGojaInterfaces["bufio"]["NewWriter"] = bufio.NewWriter
DefaultGojaInterfaces["bufio"]["NewWriterSize"] = bufio.NewWriterSize
DefaultGojaInterfaces["bufio"]["ScanBytes"] = bufio.ScanBytes
DefaultGojaInterfaces["bufio"]["ScanLines"] = bufio.ScanLines
DefaultGojaInterfaces["bufio"]["ScanRunes"] = bufio.ScanRunes
DefaultGojaInterfaces["bufio"]["ScanWords"] = bufio.ScanWords
DefaultGojaInterfaces["bufio"]["NewReadWriter"] = bufio.NewReadWriter
DefaultGojaInterfaces["bufio"]["NewReader"] = bufio.NewReader
DefaultGojaInterfaces["bufio"]["NewReaderSize"] = bufio.NewReaderSize
DefaultGojaInterfaces["bufio"]["NewScanner"] = bufio.NewScanner
DefaultGojaInterfaces["bufio"]["NewWriter"] = bufio.NewWriter
DefaultGojaInterfaces["bufio"]["NewWriterSize"] = bufio.NewWriterSize

			 			}	
			 		}
			 	}
			 }
		}
	}
}