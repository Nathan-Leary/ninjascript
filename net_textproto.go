package ninjascript
import (
net_textproto "net/textproto"
)
func init() {if _, ok := Api["net/textproto"]; !ok {
   Api["net/textproto"] = map[string]interface{}{}
}
Api["net/textproto"]["CanonicalMIMEHeaderKey"] = net_textproto.CanonicalMIMEHeaderKey
Api["net/textproto"]["TrimBytes"] = net_textproto.TrimBytes
Api["net/textproto"]["TrimString"] = net_textproto.TrimString
Api["net/textproto"]["Conn"] = net_textproto.Conn{}
Api["net/textproto"]["Dial"] = net_textproto.Dial
Api["net/textproto"]["NewConn"] = net_textproto.NewConn
Api["net/textproto"]["Error"] = net_textproto.Error{}
Api["net/textproto"]["Pipeline"] = net_textproto.Pipeline{}
Api["net/textproto"]["Reader"] = net_textproto.Reader{}
Api["net/textproto"]["NewReader"] = net_textproto.NewReader
Api["net/textproto"]["Writer"] = net_textproto.Writer{}
Api["net/textproto"]["NewWriter"] = net_textproto.NewWriter
Api["net/textproto"]["CanonicalMIMEHeaderKey"] = net_textproto.CanonicalMIMEHeaderKey
Api["net/textproto"]["TrimBytes"] = net_textproto.TrimBytes
Api["net/textproto"]["TrimString"] = net_textproto.TrimString
Api["net/textproto"]["Dial"] = net_textproto.Dial
Api["net/textproto"]["NewConn"] = net_textproto.NewConn
Api["net/textproto"]["NewReader"] = net_textproto.NewReader
Api["net/textproto"]["NewWriter"] = net_textproto.NewWriter

}