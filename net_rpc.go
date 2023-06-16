package ninjascript
import (
net_rpc "net/rpc"
)
func init() {if _, ok := Api["net/rpc"]; !ok {
   Api["net/rpc"] = map[string]interface{}{}
}
Api["net/rpc"]["DefaultRPCPath"] = net_rpc.DefaultRPCPath
Api["net/rpc"]["DefaultDebugPath"] = net_rpc.DefaultDebugPath
Api["net/rpc"]["DefaultServer"] = net_rpc.DefaultServer
Api["net/rpc"]["ErrShutdown"] = net_rpc.ErrShutdown
Api["net/rpc"]["Accept"] = net_rpc.Accept
Api["net/rpc"]["HandleHTTP"] = net_rpc.HandleHTTP
Api["net/rpc"]["Register"] = net_rpc.Register
Api["net/rpc"]["RegisterName"] = net_rpc.RegisterName
Api["net/rpc"]["ServeCodec"] = net_rpc.ServeCodec
Api["net/rpc"]["ServeConn"] = net_rpc.ServeConn
Api["net/rpc"]["ServeRequest"] = net_rpc.ServeRequest
Api["net/rpc"]["Call"] = net_rpc.Call{}
Api["net/rpc"]["Client"] = net_rpc.Client{}
Api["net/rpc"]["Dial"] = net_rpc.Dial
Api["net/rpc"]["DialHTTP"] = net_rpc.DialHTTP
Api["net/rpc"]["DialHTTPPath"] = net_rpc.DialHTTPPath
Api["net/rpc"]["NewClient"] = net_rpc.NewClient
Api["net/rpc"]["NewClientWithCodec"] = net_rpc.NewClientWithCodec
Api["net/rpc"]["Request"] = net_rpc.Request{}
Api["net/rpc"]["Response"] = net_rpc.Response{}
Api["net/rpc"]["Server"] = net_rpc.Server{}
Api["net/rpc"]["NewServer"] = net_rpc.NewServer
Api["net/rpc"]["Accept"] = net_rpc.Accept
Api["net/rpc"]["HandleHTTP"] = net_rpc.HandleHTTP
Api["net/rpc"]["Register"] = net_rpc.Register
Api["net/rpc"]["RegisterName"] = net_rpc.RegisterName
Api["net/rpc"]["ServeCodec"] = net_rpc.ServeCodec
Api["net/rpc"]["ServeConn"] = net_rpc.ServeConn
Api["net/rpc"]["ServeRequest"] = net_rpc.ServeRequest
Api["net/rpc"]["Dial"] = net_rpc.Dial
Api["net/rpc"]["DialHTTP"] = net_rpc.DialHTTP
Api["net/rpc"]["DialHTTPPath"] = net_rpc.DialHTTPPath
Api["net/rpc"]["NewClient"] = net_rpc.NewClient
Api["net/rpc"]["NewClientWithCodec"] = net_rpc.NewClientWithCodec
Api["net/rpc"]["NewServer"] = net_rpc.NewServer

}