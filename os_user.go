package ninjascript
import (
os_user "os/user"
)
func init() {if _, ok := Api["os/user"]; !ok {
   Api["os/user"] = map[string]interface{}{}
}
Api["os/user"]["Group"] = os_user.Group{}
Api["os/user"]["LookupGroup"] = os_user.LookupGroup
Api["os/user"]["LookupGroupId"] = os_user.LookupGroupId
Api["os/user"]["User"] = os_user.User{}
Api["os/user"]["Current"] = os_user.Current
Api["os/user"]["Lookup"] = os_user.Lookup
Api["os/user"]["LookupId"] = os_user.LookupId
Api["os/user"]["LookupGroup"] = os_user.LookupGroup
Api["os/user"]["LookupGroupId"] = os_user.LookupGroupId
Api["os/user"]["Current"] = os_user.Current
Api["os/user"]["Lookup"] = os_user.Lookup
Api["os/user"]["LookupId"] = os_user.LookupId

}