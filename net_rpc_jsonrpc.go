package ninjascript
import (
net_rpc_jsonrpc "net/rpc/jsonrpc"
)
func init() {if _, ok := Api["net/rpc/jsonrpc"]; !ok {
   Api["net/rpc/jsonrpc"] = map[string]interface{}{}
}
Api["net/rpc/jsonrpc"]["Dial"] = net_rpc_jsonrpc.Dial
Api["net/rpc/jsonrpc"]["NewClient"] = net_rpc_jsonrpc.NewClient
Api["net/rpc/jsonrpc"]["NewClientCodec"] = net_rpc_jsonrpc.NewClientCodec
Api["net/rpc/jsonrpc"]["NewServerCodec"] = net_rpc_jsonrpc.NewServerCodec
Api["net/rpc/jsonrpc"]["ServeConn"] = net_rpc_jsonrpc.ServeConn
Api["net/rpc/jsonrpc"]["Dial"] = net_rpc_jsonrpc.Dial
Api["net/rpc/jsonrpc"]["NewClient"] = net_rpc_jsonrpc.NewClient
Api["net/rpc/jsonrpc"]["NewClientCodec"] = net_rpc_jsonrpc.NewClientCodec
Api["net/rpc/jsonrpc"]["NewServerCodec"] = net_rpc_jsonrpc.NewServerCodec
Api["net/rpc/jsonrpc"]["ServeConn"] = net_rpc_jsonrpc.ServeConn

}