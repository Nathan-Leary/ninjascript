package zip
import (
archive_zip "archive/zip"
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
if _, ok := DefaultGojaInterfaces["zip"]; !ok {
   DefaultGojaInterfaces["zip"] = map[string]interface{}{}
}
DefaultGojaInterfaces["zip"]["Store"] = archive_zip.Store
DefaultGojaInterfaces["zip"]["Deflate"] = archive_zip.Deflate
DefaultGojaInterfaces["zip"]["RegisterCompressor"] = archive_zip.RegisterCompressor
DefaultGojaInterfaces["zip"]["RegisterDecompressor"] = archive_zip.RegisterDecompressor
DefaultGojaInterfaces["zip"]["File"] = archive_zip.File{}
DefaultGojaInterfaces["zip"]["FileHeader"] = archive_zip.FileHeader{}
DefaultGojaInterfaces["zip"]["FileInfoHeader"] = archive_zip.FileInfoHeader
DefaultGojaInterfaces["zip"]["ReadCloser"] = archive_zip.ReadCloser{}
DefaultGojaInterfaces["zip"]["OpenReader"] = archive_zip.OpenReader
DefaultGojaInterfaces["zip"]["Reader"] = archive_zip.Reader{}
DefaultGojaInterfaces["zip"]["NewReader"] = archive_zip.NewReader
DefaultGojaInterfaces["zip"]["Writer"] = archive_zip.Writer{}
DefaultGojaInterfaces["zip"]["NewWriter"] = archive_zip.NewWriter
DefaultGojaInterfaces["zip"]["RegisterCompressor"] = archive_zip.RegisterCompressor
DefaultGojaInterfaces["zip"]["RegisterDecompressor"] = archive_zip.RegisterDecompressor
DefaultGojaInterfaces["zip"]["FileInfoHeader"] = archive_zip.FileInfoHeader
DefaultGojaInterfaces["zip"]["OpenReader"] = archive_zip.OpenReader
DefaultGojaInterfaces["zip"]["NewReader"] = archive_zip.NewReader
DefaultGojaInterfaces["zip"]["NewWriter"] = archive_zip.NewWriter

			 			}	
			 		}
			 	}
			 }
		}
	}
}