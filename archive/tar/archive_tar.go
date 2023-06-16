package tar
import (
archive_tar "archive/tar"
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
if _, ok := DefaultGojaInterfaces["tar"]; !ok {
   DefaultGojaInterfaces["tar"] = map[string]interface{}{}
}
DefaultGojaInterfaces["tar"]["TypeReg"] = archive_tar.TypeReg
DefaultGojaInterfaces["tar"]["TypeRegA"] = archive_tar.TypeRegA
DefaultGojaInterfaces["tar"]["TypeLink"] = archive_tar.TypeLink
DefaultGojaInterfaces["tar"]["TypeSymlink"] = archive_tar.TypeSymlink
DefaultGojaInterfaces["tar"]["TypeChar"] = archive_tar.TypeChar
DefaultGojaInterfaces["tar"]["TypeBlock"] = archive_tar.TypeBlock
DefaultGojaInterfaces["tar"]["TypeDir"] = archive_tar.TypeDir
DefaultGojaInterfaces["tar"]["TypeFifo"] = archive_tar.TypeFifo
DefaultGojaInterfaces["tar"]["TypeCont"] = archive_tar.TypeCont
DefaultGojaInterfaces["tar"]["TypeXHeader"] = archive_tar.TypeXHeader
DefaultGojaInterfaces["tar"]["TypeXGlobalHeader"] = archive_tar.TypeXGlobalHeader
DefaultGojaInterfaces["tar"]["TypeGNUSparse"] = archive_tar.TypeGNUSparse
DefaultGojaInterfaces["tar"]["TypeGNULongName"] = archive_tar.TypeGNULongName
DefaultGojaInterfaces["tar"]["TypeGNULongLink"] = archive_tar.TypeGNULongLink
DefaultGojaInterfaces["tar"]["FormatUnknown"] = archive_tar.FormatUnknown
DefaultGojaInterfaces["tar"]["FormatUSTAR"] = archive_tar.FormatUSTAR
DefaultGojaInterfaces["tar"]["FormatPAX"] = archive_tar.FormatPAX
DefaultGojaInterfaces["tar"]["FormatGNU"] = archive_tar.FormatGNU
DefaultGojaInterfaces["tar"]["FileInfoHeader"] = archive_tar.FileInfoHeader
DefaultGojaInterfaces["tar"]["Reader"] = archive_tar.Reader{}
DefaultGojaInterfaces["tar"]["NewReader"] = archive_tar.NewReader
DefaultGojaInterfaces["tar"]["Writer"] = archive_tar.Writer{}
DefaultGojaInterfaces["tar"]["NewWriter"] = archive_tar.NewWriter
DefaultGojaInterfaces["tar"]["FileInfoHeader"] = archive_tar.FileInfoHeader
DefaultGojaInterfaces["tar"]["NewReader"] = archive_tar.NewReader
DefaultGojaInterfaces["tar"]["NewWriter"] = archive_tar.NewWriter

			 			}	
			 		}
			 	}
			 }
		}
	}
}