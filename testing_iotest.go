package ninjascript
import (
testing_iotest "testing/iotest"
)
func init() {if _, ok := Api["testing/iotest"]; !ok {
   Api["testing/iotest"] = map[string]interface{}{}
}
Api["testing/iotest"]["ErrTimeout"] = testing_iotest.ErrTimeout
Api["testing/iotest"]["DataErrReader"] = testing_iotest.DataErrReader
Api["testing/iotest"]["ErrReader"] = testing_iotest.ErrReader
Api["testing/iotest"]["HalfReader"] = testing_iotest.HalfReader
Api["testing/iotest"]["NewReadLogger"] = testing_iotest.NewReadLogger
Api["testing/iotest"]["NewWriteLogger"] = testing_iotest.NewWriteLogger
Api["testing/iotest"]["OneByteReader"] = testing_iotest.OneByteReader
Api["testing/iotest"]["TestReader"] = testing_iotest.TestReader
Api["testing/iotest"]["TimeoutReader"] = testing_iotest.TimeoutReader
Api["testing/iotest"]["TruncateWriter"] = testing_iotest.TruncateWriter
Api["testing/iotest"]["DataErrReader"] = testing_iotest.DataErrReader
Api["testing/iotest"]["ErrReader"] = testing_iotest.ErrReader
Api["testing/iotest"]["HalfReader"] = testing_iotest.HalfReader
Api["testing/iotest"]["NewReadLogger"] = testing_iotest.NewReadLogger
Api["testing/iotest"]["NewWriteLogger"] = testing_iotest.NewWriteLogger
Api["testing/iotest"]["OneByteReader"] = testing_iotest.OneByteReader
Api["testing/iotest"]["TestReader"] = testing_iotest.TestReader
Api["testing/iotest"]["TimeoutReader"] = testing_iotest.TimeoutReader
Api["testing/iotest"]["TruncateWriter"] = testing_iotest.TruncateWriter

}