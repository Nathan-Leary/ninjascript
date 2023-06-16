package ninjascript
import (
encoding_csv "encoding/csv"
)
func init() {if _, ok := Api["encoding/csv"]; !ok {
   Api["encoding/csv"] = map[string]interface{}{}
}
Api["encoding/csv"]["ParseError"] = encoding_csv.ParseError{}
Api["encoding/csv"]["Reader"] = encoding_csv.Reader{}
Api["encoding/csv"]["NewReader"] = encoding_csv.NewReader
Api["encoding/csv"]["Writer"] = encoding_csv.Writer{}
Api["encoding/csv"]["NewWriter"] = encoding_csv.NewWriter
Api["encoding/csv"]["NewReader"] = encoding_csv.NewReader
Api["encoding/csv"]["NewWriter"] = encoding_csv.NewWriter

}