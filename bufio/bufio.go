package bufio
import (
bufio "bufio"
)
func Load() {
if _, ok := Api["bufio"]; !ok {
   Api["bufio"] = map[string]interface{}{}
}
Api["bufio"]["MaxScanTokenSize"] = bufio.MaxScanTokenSize
Api["bufio"]["ErrFinalToken"] = bufio.ErrFinalToken
Api["bufio"]["ScanBytes"] = bufio.ScanBytes
Api["bufio"]["ScanLines"] = bufio.ScanLines
Api["bufio"]["ScanRunes"] = bufio.ScanRunes
Api["bufio"]["ScanWords"] = bufio.ScanWords
Api["bufio"]["ReadWriter"] = bufio.ReadWriter{}
Api["bufio"]["NewReadWriter"] = bufio.NewReadWriter
Api["bufio"]["Reader"] = bufio.Reader{}
Api["bufio"]["NewReader"] = bufio.NewReader
Api["bufio"]["NewReaderSize"] = bufio.NewReaderSize
Api["bufio"]["Scanner"] = bufio.Scanner{}
Api["bufio"]["NewScanner"] = bufio.NewScanner
Api["bufio"]["Writer"] = bufio.Writer{}
Api["bufio"]["NewWriter"] = bufio.NewWriter
Api["bufio"]["NewWriterSize"] = bufio.NewWriterSize
Api["bufio"]["ScanBytes"] = bufio.ScanBytes
Api["bufio"]["ScanLines"] = bufio.ScanLines
Api["bufio"]["ScanRunes"] = bufio.ScanRunes
Api["bufio"]["ScanWords"] = bufio.ScanWords
Api["bufio"]["NewReadWriter"] = bufio.NewReadWriter
Api["bufio"]["NewReader"] = bufio.NewReader
Api["bufio"]["NewReaderSize"] = bufio.NewReaderSize
Api["bufio"]["NewScanner"] = bufio.NewScanner
Api["bufio"]["NewWriter"] = bufio.NewWriter
Api["bufio"]["NewWriterSize"] = bufio.NewWriterSize
}
