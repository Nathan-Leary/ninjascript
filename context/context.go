package context
import (
context "context"
)
func Load() {
if _, ok := Api["context"]; !ok {
   Api["context"] = map[string]interface{}{}
}
Api["context"]["Canceled"] = context.Canceled
Api["context"]["DeadlineExceeded"] = context.DeadlineExceeded
Api["context"]["Cause"] = context.Cause
Api["context"]["WithCancel"] = context.WithCancel
Api["context"]["WithCancelCause"] = context.WithCancelCause
Api["context"]["WithDeadline"] = context.WithDeadline
Api["context"]["WithTimeout"] = context.WithTimeout
Api["context"]["Context"] = new(context.Context)
Api["context"]["Background"] = context.Background
Api["context"]["TODO"] = context.TODO
Api["context"]["WithValue"] = context.WithValue
Api["context"]["Cause"] = context.Cause
Api["context"]["WithCancel"] = context.WithCancel
Api["context"]["WithCancelCause"] = context.WithCancelCause
Api["context"]["WithDeadline"] = context.WithDeadline
Api["context"]["WithTimeout"] = context.WithTimeout
Api["context"]["Background"] = context.Background
Api["context"]["TODO"] = context.TODO
Api["context"]["WithValue"] = context.WithValue
}
