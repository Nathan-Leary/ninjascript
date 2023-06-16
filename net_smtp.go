package ninjascript
import (
net_smtp "net/smtp"
)
func init() {if _, ok := Api["net/smtp"]; !ok {
   Api["net/smtp"] = map[string]interface{}{}
}
Api["net/smtp"]["SendMail"] = net_smtp.SendMail
Api["net/smtp"]["CRAMMD5Auth"] = net_smtp.CRAMMD5Auth
Api["net/smtp"]["PlainAuth"] = net_smtp.PlainAuth
Api["net/smtp"]["Client"] = net_smtp.Client{}
Api["net/smtp"]["Dial"] = net_smtp.Dial
Api["net/smtp"]["NewClient"] = net_smtp.NewClient
Api["net/smtp"]["ServerInfo"] = net_smtp.ServerInfo{}
Api["net/smtp"]["SendMail"] = net_smtp.SendMail
Api["net/smtp"]["CRAMMD5Auth"] = net_smtp.CRAMMD5Auth
Api["net/smtp"]["PlainAuth"] = net_smtp.PlainAuth
Api["net/smtp"]["Dial"] = net_smtp.Dial
Api["net/smtp"]["NewClient"] = net_smtp.NewClient

}