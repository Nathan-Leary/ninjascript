package ninjascript
import (
text_scanner "text/scanner"
)
func init() {if _, ok := Api["text/scanner"]; !ok {
   Api["text/scanner"] = map[string]interface{}{}
}
Api["text/scanner"]["ScanIdents"] = text_scanner.ScanIdents
Api["text/scanner"]["ScanInts"] = text_scanner.ScanInts
Api["text/scanner"]["ScanFloats"] = text_scanner.ScanFloats
Api["text/scanner"]["ScanChars"] = text_scanner.ScanChars
Api["text/scanner"]["ScanStrings"] = text_scanner.ScanStrings
Api["text/scanner"]["ScanRawStrings"] = text_scanner.ScanRawStrings
Api["text/scanner"]["ScanComments"] = text_scanner.ScanComments
Api["text/scanner"]["SkipComments"] = text_scanner.SkipComments
Api["text/scanner"]["GoTokens"] = text_scanner.GoTokens
Api["text/scanner"]["EOF"] = text_scanner.EOF
Api["text/scanner"]["Ident"] = text_scanner.Ident
Api["text/scanner"]["Int"] = text_scanner.Int
Api["text/scanner"]["Float"] = text_scanner.Float
Api["text/scanner"]["Char"] = text_scanner.Char
Api["text/scanner"]["String"] = text_scanner.String
Api["text/scanner"]["RawString"] = text_scanner.RawString
Api["text/scanner"]["Comment"] = text_scanner.Comment
Api["text/scanner"]["TokenString"] = text_scanner.TokenString
Api["text/scanner"]["Position"] = text_scanner.Position{}
Api["text/scanner"]["Scanner"] = text_scanner.Scanner{}
Api["text/scanner"]["TokenString"] = text_scanner.TokenString

}