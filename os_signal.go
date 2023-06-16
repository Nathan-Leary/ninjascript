package ninjascript
import (
os_signal "os/signal"
)
func init() {if _, ok := Api["os/signal"]; !ok {
   Api["os/signal"] = map[string]interface{}{}
}
Api["os/signal"]["Ignore"] = os_signal.Ignore
Api["os/signal"]["Ignored"] = os_signal.Ignored
Api["os/signal"]["Notify"] = os_signal.Notify
Api["os/signal"]["NotifyContext"] = os_signal.NotifyContext
Api["os/signal"]["Reset"] = os_signal.Reset
Api["os/signal"]["Stop"] = os_signal.Stop
Api["os/signal"]["Ignore"] = os_signal.Ignore
Api["os/signal"]["Ignored"] = os_signal.Ignored
Api["os/signal"]["Notify"] = os_signal.Notify
Api["os/signal"]["NotifyContext"] = os_signal.NotifyContext
Api["os/signal"]["Reset"] = os_signal.Reset
Api["os/signal"]["Stop"] = os_signal.Stop

}