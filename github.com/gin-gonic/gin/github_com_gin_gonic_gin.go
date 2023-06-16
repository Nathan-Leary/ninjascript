package gin
import (
github_com_gin_gonic_gin "github.com/gin-gonic/gin"
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
if _, ok := DefaultGojaInterfaces["gin"]; !ok {
   DefaultGojaInterfaces["gin"] = map[string]interface{}{}
}
DefaultGojaInterfaces["gin"]["MIMEJSON"] = github_com_gin_gonic_gin.MIMEJSON
DefaultGojaInterfaces["gin"]["MIMEHTML"] = github_com_gin_gonic_gin.MIMEHTML
DefaultGojaInterfaces["gin"]["MIMEXML"] = github_com_gin_gonic_gin.MIMEXML
DefaultGojaInterfaces["gin"]["MIMEXML2"] = github_com_gin_gonic_gin.MIMEXML2
DefaultGojaInterfaces["gin"]["MIMEPlain"] = github_com_gin_gonic_gin.MIMEPlain
DefaultGojaInterfaces["gin"]["MIMEPOSTForm"] = github_com_gin_gonic_gin.MIMEPOSTForm
DefaultGojaInterfaces["gin"]["MIMEMultipartPOSTForm"] = github_com_gin_gonic_gin.MIMEMultipartPOSTForm
DefaultGojaInterfaces["gin"]["MIMEYAML"] = github_com_gin_gonic_gin.MIMEYAML
DefaultGojaInterfaces["gin"]["MIMETOML"] = github_com_gin_gonic_gin.MIMETOML
DefaultGojaInterfaces["gin"]["PlatformGoogleAppEngine"] = github_com_gin_gonic_gin.PlatformGoogleAppEngine
DefaultGojaInterfaces["gin"]["PlatformCloudflare"] = github_com_gin_gonic_gin.PlatformCloudflare
DefaultGojaInterfaces["gin"]["DebugMode"] = github_com_gin_gonic_gin.DebugMode
DefaultGojaInterfaces["gin"]["ReleaseMode"] = github_com_gin_gonic_gin.ReleaseMode
DefaultGojaInterfaces["gin"]["TestMode"] = github_com_gin_gonic_gin.TestMode
DefaultGojaInterfaces["gin"]["DebugPrintRouteFunc"] = github_com_gin_gonic_gin.DebugPrintRouteFunc
DefaultGojaInterfaces["gin"]["DefaultErrorWriter"] = github_com_gin_gonic_gin.DefaultErrorWriter
DefaultGojaInterfaces["gin"]["DefaultWriter"] = github_com_gin_gonic_gin.DefaultWriter
DefaultGojaInterfaces["gin"]["CreateTestContext"] = github_com_gin_gonic_gin.CreateTestContext
DefaultGojaInterfaces["gin"]["Dir"] = github_com_gin_gonic_gin.Dir
DefaultGojaInterfaces["gin"]["DisableBindValidation"] = github_com_gin_gonic_gin.DisableBindValidation
DefaultGojaInterfaces["gin"]["DisableConsoleColor"] = github_com_gin_gonic_gin.DisableConsoleColor
DefaultGojaInterfaces["gin"]["EnableJsonDecoderDisallowUnknownFields"] = github_com_gin_gonic_gin.EnableJsonDecoderDisallowUnknownFields
DefaultGojaInterfaces["gin"]["EnableJsonDecoderUseNumber"] = github_com_gin_gonic_gin.EnableJsonDecoderUseNumber
DefaultGojaInterfaces["gin"]["ForceConsoleColor"] = github_com_gin_gonic_gin.ForceConsoleColor
DefaultGojaInterfaces["gin"]["IsDebugging"] = github_com_gin_gonic_gin.IsDebugging
DefaultGojaInterfaces["gin"]["Mode"] = github_com_gin_gonic_gin.Mode
DefaultGojaInterfaces["gin"]["SetMode"] = github_com_gin_gonic_gin.SetMode
DefaultGojaInterfaces["gin"]["Context"] = github_com_gin_gonic_gin.Context{}
DefaultGojaInterfaces["gin"]["CreateTestContextOnly"] = github_com_gin_gonic_gin.CreateTestContextOnly
DefaultGojaInterfaces["gin"]["Engine"] = github_com_gin_gonic_gin.Engine{}
DefaultGojaInterfaces["gin"]["Default"] = github_com_gin_gonic_gin.Default
DefaultGojaInterfaces["gin"]["New"] = github_com_gin_gonic_gin.New
DefaultGojaInterfaces["gin"]["Error"] = github_com_gin_gonic_gin.Error{}
DefaultGojaInterfaces["gin"]["ErrorTypeBind"] = github_com_gin_gonic_gin.ErrorTypeBind
DefaultGojaInterfaces["gin"]["ErrorTypeRender"] = github_com_gin_gonic_gin.ErrorTypeRender
DefaultGojaInterfaces["gin"]["ErrorTypePrivate"] = github_com_gin_gonic_gin.ErrorTypePrivate
DefaultGojaInterfaces["gin"]["ErrorTypePublic"] = github_com_gin_gonic_gin.ErrorTypePublic
DefaultGojaInterfaces["gin"]["ErrorTypeAny"] = github_com_gin_gonic_gin.ErrorTypeAny
DefaultGojaInterfaces["gin"]["ErrorTypeNu"] = github_com_gin_gonic_gin.ErrorTypeNu
DefaultGojaInterfaces["gin"]["BasicAuth"] = github_com_gin_gonic_gin.BasicAuth
DefaultGojaInterfaces["gin"]["BasicAuthForRealm"] = github_com_gin_gonic_gin.BasicAuthForRealm
DefaultGojaInterfaces["gin"]["Bind"] = github_com_gin_gonic_gin.Bind
DefaultGojaInterfaces["gin"]["CustomRecovery"] = github_com_gin_gonic_gin.CustomRecovery
DefaultGojaInterfaces["gin"]["CustomRecoveryWithWriter"] = github_com_gin_gonic_gin.CustomRecoveryWithWriter
DefaultGojaInterfaces["gin"]["ErrorLogger"] = github_com_gin_gonic_gin.ErrorLogger
DefaultGojaInterfaces["gin"]["ErrorLoggerT"] = github_com_gin_gonic_gin.ErrorLoggerT
DefaultGojaInterfaces["gin"]["Logger"] = github_com_gin_gonic_gin.Logger
DefaultGojaInterfaces["gin"]["LoggerWithConfig"] = github_com_gin_gonic_gin.LoggerWithConfig
DefaultGojaInterfaces["gin"]["LoggerWithFormatter"] = github_com_gin_gonic_gin.LoggerWithFormatter
DefaultGojaInterfaces["gin"]["LoggerWithWriter"] = github_com_gin_gonic_gin.LoggerWithWriter
DefaultGojaInterfaces["gin"]["Recovery"] = github_com_gin_gonic_gin.Recovery
DefaultGojaInterfaces["gin"]["RecoveryWithWriter"] = github_com_gin_gonic_gin.RecoveryWithWriter
DefaultGojaInterfaces["gin"]["WrapF"] = github_com_gin_gonic_gin.WrapF
DefaultGojaInterfaces["gin"]["WrapH"] = github_com_gin_gonic_gin.WrapH
DefaultGojaInterfaces["gin"]["LogFormatterParams"] = github_com_gin_gonic_gin.LogFormatterParams{}
DefaultGojaInterfaces["gin"]["LoggerConfig"] = github_com_gin_gonic_gin.LoggerConfig{}
DefaultGojaInterfaces["gin"]["Negotiate"] = github_com_gin_gonic_gin.Negotiate{}
DefaultGojaInterfaces["gin"]["Param"] = github_com_gin_gonic_gin.Param{}
DefaultGojaInterfaces["gin"]["RouteInfo"] = github_com_gin_gonic_gin.RouteInfo{}
DefaultGojaInterfaces["gin"]["RouterGroup"] = github_com_gin_gonic_gin.RouterGroup{}
DefaultGojaInterfaces["gin"]["CreateTestContext"] = github_com_gin_gonic_gin.CreateTestContext
DefaultGojaInterfaces["gin"]["Dir"] = github_com_gin_gonic_gin.Dir
DefaultGojaInterfaces["gin"]["DisableBindValidation"] = github_com_gin_gonic_gin.DisableBindValidation
DefaultGojaInterfaces["gin"]["DisableConsoleColor"] = github_com_gin_gonic_gin.DisableConsoleColor
DefaultGojaInterfaces["gin"]["EnableJsonDecoderDisallowUnknownFields"] = github_com_gin_gonic_gin.EnableJsonDecoderDisallowUnknownFields
DefaultGojaInterfaces["gin"]["EnableJsonDecoderUseNumber"] = github_com_gin_gonic_gin.EnableJsonDecoderUseNumber
DefaultGojaInterfaces["gin"]["ForceConsoleColor"] = github_com_gin_gonic_gin.ForceConsoleColor
DefaultGojaInterfaces["gin"]["IsDebugging"] = github_com_gin_gonic_gin.IsDebugging
DefaultGojaInterfaces["gin"]["Mode"] = github_com_gin_gonic_gin.Mode
DefaultGojaInterfaces["gin"]["SetMode"] = github_com_gin_gonic_gin.SetMode
DefaultGojaInterfaces["gin"]["CreateTestContextOnly"] = github_com_gin_gonic_gin.CreateTestContextOnly
DefaultGojaInterfaces["gin"]["Default"] = github_com_gin_gonic_gin.Default
DefaultGojaInterfaces["gin"]["New"] = github_com_gin_gonic_gin.New
DefaultGojaInterfaces["gin"]["BasicAuth"] = github_com_gin_gonic_gin.BasicAuth
DefaultGojaInterfaces["gin"]["BasicAuthForRealm"] = github_com_gin_gonic_gin.BasicAuthForRealm
DefaultGojaInterfaces["gin"]["Bind"] = github_com_gin_gonic_gin.Bind
DefaultGojaInterfaces["gin"]["CustomRecovery"] = github_com_gin_gonic_gin.CustomRecovery
DefaultGojaInterfaces["gin"]["CustomRecoveryWithWriter"] = github_com_gin_gonic_gin.CustomRecoveryWithWriter
DefaultGojaInterfaces["gin"]["ErrorLogger"] = github_com_gin_gonic_gin.ErrorLogger
DefaultGojaInterfaces["gin"]["ErrorLoggerT"] = github_com_gin_gonic_gin.ErrorLoggerT
DefaultGojaInterfaces["gin"]["Logger"] = github_com_gin_gonic_gin.Logger
DefaultGojaInterfaces["gin"]["LoggerWithConfig"] = github_com_gin_gonic_gin.LoggerWithConfig
DefaultGojaInterfaces["gin"]["LoggerWithFormatter"] = github_com_gin_gonic_gin.LoggerWithFormatter
DefaultGojaInterfaces["gin"]["LoggerWithWriter"] = github_com_gin_gonic_gin.LoggerWithWriter
DefaultGojaInterfaces["gin"]["Recovery"] = github_com_gin_gonic_gin.Recovery
DefaultGojaInterfaces["gin"]["RecoveryWithWriter"] = github_com_gin_gonic_gin.RecoveryWithWriter
DefaultGojaInterfaces["gin"]["WrapF"] = github_com_gin_gonic_gin.WrapF
DefaultGojaInterfaces["gin"]["WrapH"] = github_com_gin_gonic_gin.WrapH

			 			}	
			 		}
			 	}
			 }
		}
	}
}