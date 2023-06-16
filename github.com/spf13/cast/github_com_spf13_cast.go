package cast
import (
github_com_spf13_cast "github.com/spf13/cast"
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
if _, ok := DefaultGojaInterfaces["cast"]; !ok {
   DefaultGojaInterfaces["cast"] = map[string]interface{}{}
}
DefaultGojaInterfaces["cast"]["StringToDate"] = github_com_spf13_cast.StringToDate
DefaultGojaInterfaces["cast"]["StringToDateInDefaultLocation"] = github_com_spf13_cast.StringToDateInDefaultLocation
DefaultGojaInterfaces["cast"]["ToBool"] = github_com_spf13_cast.ToBool
DefaultGojaInterfaces["cast"]["ToBoolE"] = github_com_spf13_cast.ToBoolE
DefaultGojaInterfaces["cast"]["ToBoolSlice"] = github_com_spf13_cast.ToBoolSlice
DefaultGojaInterfaces["cast"]["ToBoolSliceE"] = github_com_spf13_cast.ToBoolSliceE
DefaultGojaInterfaces["cast"]["ToDuration"] = github_com_spf13_cast.ToDuration
DefaultGojaInterfaces["cast"]["ToDurationE"] = github_com_spf13_cast.ToDurationE
DefaultGojaInterfaces["cast"]["ToDurationSlice"] = github_com_spf13_cast.ToDurationSlice
DefaultGojaInterfaces["cast"]["ToDurationSliceE"] = github_com_spf13_cast.ToDurationSliceE
DefaultGojaInterfaces["cast"]["ToFloat32"] = github_com_spf13_cast.ToFloat32
DefaultGojaInterfaces["cast"]["ToFloat32E"] = github_com_spf13_cast.ToFloat32E
DefaultGojaInterfaces["cast"]["ToFloat64"] = github_com_spf13_cast.ToFloat64
DefaultGojaInterfaces["cast"]["ToFloat64E"] = github_com_spf13_cast.ToFloat64E
DefaultGojaInterfaces["cast"]["ToInt"] = github_com_spf13_cast.ToInt
DefaultGojaInterfaces["cast"]["ToInt16"] = github_com_spf13_cast.ToInt16
DefaultGojaInterfaces["cast"]["ToInt16E"] = github_com_spf13_cast.ToInt16E
DefaultGojaInterfaces["cast"]["ToInt32"] = github_com_spf13_cast.ToInt32
DefaultGojaInterfaces["cast"]["ToInt32E"] = github_com_spf13_cast.ToInt32E
DefaultGojaInterfaces["cast"]["ToInt64"] = github_com_spf13_cast.ToInt64
DefaultGojaInterfaces["cast"]["ToInt64E"] = github_com_spf13_cast.ToInt64E
DefaultGojaInterfaces["cast"]["ToInt8"] = github_com_spf13_cast.ToInt8
DefaultGojaInterfaces["cast"]["ToInt8E"] = github_com_spf13_cast.ToInt8E
DefaultGojaInterfaces["cast"]["ToIntE"] = github_com_spf13_cast.ToIntE
DefaultGojaInterfaces["cast"]["ToIntSlice"] = github_com_spf13_cast.ToIntSlice
DefaultGojaInterfaces["cast"]["ToIntSliceE"] = github_com_spf13_cast.ToIntSliceE
DefaultGojaInterfaces["cast"]["ToSlice"] = github_com_spf13_cast.ToSlice
DefaultGojaInterfaces["cast"]["ToSliceE"] = github_com_spf13_cast.ToSliceE
DefaultGojaInterfaces["cast"]["ToString"] = github_com_spf13_cast.ToString
DefaultGojaInterfaces["cast"]["ToStringE"] = github_com_spf13_cast.ToStringE
DefaultGojaInterfaces["cast"]["ToStringMap"] = github_com_spf13_cast.ToStringMap
DefaultGojaInterfaces["cast"]["ToStringMapBool"] = github_com_spf13_cast.ToStringMapBool
DefaultGojaInterfaces["cast"]["ToStringMapBoolE"] = github_com_spf13_cast.ToStringMapBoolE
DefaultGojaInterfaces["cast"]["ToStringMapE"] = github_com_spf13_cast.ToStringMapE
DefaultGojaInterfaces["cast"]["ToStringMapInt"] = github_com_spf13_cast.ToStringMapInt
DefaultGojaInterfaces["cast"]["ToStringMapInt64"] = github_com_spf13_cast.ToStringMapInt64
DefaultGojaInterfaces["cast"]["ToStringMapInt64E"] = github_com_spf13_cast.ToStringMapInt64E
DefaultGojaInterfaces["cast"]["ToStringMapIntE"] = github_com_spf13_cast.ToStringMapIntE
DefaultGojaInterfaces["cast"]["ToStringMapString"] = github_com_spf13_cast.ToStringMapString
DefaultGojaInterfaces["cast"]["ToStringMapStringE"] = github_com_spf13_cast.ToStringMapStringE
DefaultGojaInterfaces["cast"]["ToStringMapStringSlice"] = github_com_spf13_cast.ToStringMapStringSlice
DefaultGojaInterfaces["cast"]["ToStringMapStringSliceE"] = github_com_spf13_cast.ToStringMapStringSliceE
DefaultGojaInterfaces["cast"]["ToStringSlice"] = github_com_spf13_cast.ToStringSlice
DefaultGojaInterfaces["cast"]["ToStringSliceE"] = github_com_spf13_cast.ToStringSliceE
DefaultGojaInterfaces["cast"]["ToTime"] = github_com_spf13_cast.ToTime
DefaultGojaInterfaces["cast"]["ToTimeE"] = github_com_spf13_cast.ToTimeE
DefaultGojaInterfaces["cast"]["ToTimeInDefaultLocation"] = github_com_spf13_cast.ToTimeInDefaultLocation
DefaultGojaInterfaces["cast"]["ToTimeInDefaultLocationE"] = github_com_spf13_cast.ToTimeInDefaultLocationE
DefaultGojaInterfaces["cast"]["ToUint"] = github_com_spf13_cast.ToUint
DefaultGojaInterfaces["cast"]["ToUint16"] = github_com_spf13_cast.ToUint16
DefaultGojaInterfaces["cast"]["ToUint16E"] = github_com_spf13_cast.ToUint16E
DefaultGojaInterfaces["cast"]["ToUint32"] = github_com_spf13_cast.ToUint32
DefaultGojaInterfaces["cast"]["ToUint32E"] = github_com_spf13_cast.ToUint32E
DefaultGojaInterfaces["cast"]["ToUint64"] = github_com_spf13_cast.ToUint64
DefaultGojaInterfaces["cast"]["ToUint64E"] = github_com_spf13_cast.ToUint64E
DefaultGojaInterfaces["cast"]["ToUint8"] = github_com_spf13_cast.ToUint8
DefaultGojaInterfaces["cast"]["ToUint8E"] = github_com_spf13_cast.ToUint8E
DefaultGojaInterfaces["cast"]["ToUintE"] = github_com_spf13_cast.ToUintE
DefaultGojaInterfaces["cast"]["StringToDate"] = github_com_spf13_cast.StringToDate
DefaultGojaInterfaces["cast"]["StringToDateInDefaultLocation"] = github_com_spf13_cast.StringToDateInDefaultLocation
DefaultGojaInterfaces["cast"]["ToBool"] = github_com_spf13_cast.ToBool
DefaultGojaInterfaces["cast"]["ToBoolE"] = github_com_spf13_cast.ToBoolE
DefaultGojaInterfaces["cast"]["ToBoolSlice"] = github_com_spf13_cast.ToBoolSlice
DefaultGojaInterfaces["cast"]["ToBoolSliceE"] = github_com_spf13_cast.ToBoolSliceE
DefaultGojaInterfaces["cast"]["ToDuration"] = github_com_spf13_cast.ToDuration
DefaultGojaInterfaces["cast"]["ToDurationE"] = github_com_spf13_cast.ToDurationE
DefaultGojaInterfaces["cast"]["ToDurationSlice"] = github_com_spf13_cast.ToDurationSlice
DefaultGojaInterfaces["cast"]["ToDurationSliceE"] = github_com_spf13_cast.ToDurationSliceE
DefaultGojaInterfaces["cast"]["ToFloat32"] = github_com_spf13_cast.ToFloat32
DefaultGojaInterfaces["cast"]["ToFloat32E"] = github_com_spf13_cast.ToFloat32E
DefaultGojaInterfaces["cast"]["ToFloat64"] = github_com_spf13_cast.ToFloat64
DefaultGojaInterfaces["cast"]["ToFloat64E"] = github_com_spf13_cast.ToFloat64E
DefaultGojaInterfaces["cast"]["ToInt"] = github_com_spf13_cast.ToInt
DefaultGojaInterfaces["cast"]["ToInt16"] = github_com_spf13_cast.ToInt16
DefaultGojaInterfaces["cast"]["ToInt16E"] = github_com_spf13_cast.ToInt16E
DefaultGojaInterfaces["cast"]["ToInt32"] = github_com_spf13_cast.ToInt32
DefaultGojaInterfaces["cast"]["ToInt32E"] = github_com_spf13_cast.ToInt32E
DefaultGojaInterfaces["cast"]["ToInt64"] = github_com_spf13_cast.ToInt64
DefaultGojaInterfaces["cast"]["ToInt64E"] = github_com_spf13_cast.ToInt64E
DefaultGojaInterfaces["cast"]["ToInt8"] = github_com_spf13_cast.ToInt8
DefaultGojaInterfaces["cast"]["ToInt8E"] = github_com_spf13_cast.ToInt8E
DefaultGojaInterfaces["cast"]["ToIntE"] = github_com_spf13_cast.ToIntE
DefaultGojaInterfaces["cast"]["ToIntSlice"] = github_com_spf13_cast.ToIntSlice
DefaultGojaInterfaces["cast"]["ToIntSliceE"] = github_com_spf13_cast.ToIntSliceE
DefaultGojaInterfaces["cast"]["ToSlice"] = github_com_spf13_cast.ToSlice
DefaultGojaInterfaces["cast"]["ToSliceE"] = github_com_spf13_cast.ToSliceE
DefaultGojaInterfaces["cast"]["ToString"] = github_com_spf13_cast.ToString
DefaultGojaInterfaces["cast"]["ToStringE"] = github_com_spf13_cast.ToStringE
DefaultGojaInterfaces["cast"]["ToStringMap"] = github_com_spf13_cast.ToStringMap
DefaultGojaInterfaces["cast"]["ToStringMapBool"] = github_com_spf13_cast.ToStringMapBool
DefaultGojaInterfaces["cast"]["ToStringMapBoolE"] = github_com_spf13_cast.ToStringMapBoolE
DefaultGojaInterfaces["cast"]["ToStringMapE"] = github_com_spf13_cast.ToStringMapE
DefaultGojaInterfaces["cast"]["ToStringMapInt"] = github_com_spf13_cast.ToStringMapInt
DefaultGojaInterfaces["cast"]["ToStringMapInt64"] = github_com_spf13_cast.ToStringMapInt64
DefaultGojaInterfaces["cast"]["ToStringMapInt64E"] = github_com_spf13_cast.ToStringMapInt64E
DefaultGojaInterfaces["cast"]["ToStringMapIntE"] = github_com_spf13_cast.ToStringMapIntE
DefaultGojaInterfaces["cast"]["ToStringMapString"] = github_com_spf13_cast.ToStringMapString
DefaultGojaInterfaces["cast"]["ToStringMapStringE"] = github_com_spf13_cast.ToStringMapStringE
DefaultGojaInterfaces["cast"]["ToStringMapStringSlice"] = github_com_spf13_cast.ToStringMapStringSlice
DefaultGojaInterfaces["cast"]["ToStringMapStringSliceE"] = github_com_spf13_cast.ToStringMapStringSliceE
DefaultGojaInterfaces["cast"]["ToStringSlice"] = github_com_spf13_cast.ToStringSlice
DefaultGojaInterfaces["cast"]["ToStringSliceE"] = github_com_spf13_cast.ToStringSliceE
DefaultGojaInterfaces["cast"]["ToTime"] = github_com_spf13_cast.ToTime
DefaultGojaInterfaces["cast"]["ToTimeE"] = github_com_spf13_cast.ToTimeE
DefaultGojaInterfaces["cast"]["ToTimeInDefaultLocation"] = github_com_spf13_cast.ToTimeInDefaultLocation
DefaultGojaInterfaces["cast"]["ToTimeInDefaultLocationE"] = github_com_spf13_cast.ToTimeInDefaultLocationE
DefaultGojaInterfaces["cast"]["ToUint"] = github_com_spf13_cast.ToUint
DefaultGojaInterfaces["cast"]["ToUint16"] = github_com_spf13_cast.ToUint16
DefaultGojaInterfaces["cast"]["ToUint16E"] = github_com_spf13_cast.ToUint16E
DefaultGojaInterfaces["cast"]["ToUint32"] = github_com_spf13_cast.ToUint32
DefaultGojaInterfaces["cast"]["ToUint32E"] = github_com_spf13_cast.ToUint32E
DefaultGojaInterfaces["cast"]["ToUint64"] = github_com_spf13_cast.ToUint64
DefaultGojaInterfaces["cast"]["ToUint64E"] = github_com_spf13_cast.ToUint64E
DefaultGojaInterfaces["cast"]["ToUint8"] = github_com_spf13_cast.ToUint8
DefaultGojaInterfaces["cast"]["ToUint8E"] = github_com_spf13_cast.ToUint8E
DefaultGojaInterfaces["cast"]["ToUintE"] = github_com_spf13_cast.ToUintE

			 			}	
			 		}
			 	}
			 }
		}
	}
}