package ninjascript
import (
net_mail "net/mail"
)
func init() {if _, ok := Api["net/mail"]; !ok {
   Api["net/mail"] = map[string]interface{}{}
}
Api["net/mail"]["ErrHeaderNotPresent"] = net_mail.ErrHeaderNotPresent
Api["net/mail"]["ParseDate"] = net_mail.ParseDate
Api["net/mail"]["Address"] = net_mail.Address{}
Api["net/mail"]["ParseAddress"] = net_mail.ParseAddress
Api["net/mail"]["ParseAddressList"] = net_mail.ParseAddressList
Api["net/mail"]["AddressParser"] = net_mail.AddressParser{}
Api["net/mail"]["Message"] = net_mail.Message{}
Api["net/mail"]["ReadMessage"] = net_mail.ReadMessage
Api["net/mail"]["ParseDate"] = net_mail.ParseDate
Api["net/mail"]["ParseAddress"] = net_mail.ParseAddress
Api["net/mail"]["ParseAddressList"] = net_mail.ParseAddressList
Api["net/mail"]["ReadMessage"] = net_mail.ReadMessage

}