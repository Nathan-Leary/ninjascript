package ninjascript
import (
os_exec "os/exec"
)
func init() {if _, ok := Api["os/exec"]; !ok {
   Api["os/exec"] = map[string]interface{}{}
}
Api["os/exec"]["ErrDot"] = os_exec.ErrDot
Api["os/exec"]["ErrNotFound"] = os_exec.ErrNotFound
Api["os/exec"]["LookPath"] = os_exec.LookPath
Api["os/exec"]["Cmd"] = os_exec.Cmd{}
Api["os/exec"]["Command"] = os_exec.Command
Api["os/exec"]["CommandContext"] = os_exec.CommandContext
Api["os/exec"]["Error"] = os_exec.Error{}
Api["os/exec"]["ExitError"] = os_exec.ExitError{}
Api["os/exec"]["LookPath"] = os_exec.LookPath
Api["os/exec"]["Command"] = os_exec.Command
Api["os/exec"]["CommandContext"] = os_exec.CommandContext

}