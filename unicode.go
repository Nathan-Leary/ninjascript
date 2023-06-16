package ninjascript
import (
unicode "unicode"
)
func init() {if _, ok := Api["unicode"]; !ok {
   Api["unicode"] = map[string]interface{}{}
}
Api["unicode"]["MaxRune"] = unicode.MaxRune
Api["unicode"]["ReplacementChar"] = unicode.ReplacementChar
Api["unicode"]["MaxASCII"] = unicode.MaxASCII
Api["unicode"]["MaxLatin1"] = unicode.MaxLatin1
Api["unicode"]["UpperCase"] = unicode.UpperCase
Api["unicode"]["LowerCase"] = unicode.LowerCase
Api["unicode"]["TitleCase"] = unicode.TitleCase
Api["unicode"]["MaxCase"] = unicode.MaxCase
Api["unicode"]["UpperLower"] = unicode.UpperLower
Api["unicode"]["CaseRanges"] = unicode.CaseRanges
Api["unicode"]["Categories"] = unicode.Categories
Api["unicode"]["FoldCategory"] = unicode.FoldCategory
Api["unicode"]["FoldScript"] = unicode.FoldScript
Api["unicode"]["GraphicRanges"] = unicode.GraphicRanges
Api["unicode"]["PrintRanges"] = unicode.PrintRanges
Api["unicode"]["Properties"] = unicode.Properties
Api["unicode"]["Scripts"] = unicode.Scripts
Api["unicode"]["In"] = unicode.In
Api["unicode"]["Is"] = unicode.Is
Api["unicode"]["IsControl"] = unicode.IsControl
Api["unicode"]["IsDigit"] = unicode.IsDigit
Api["unicode"]["IsGraphic"] = unicode.IsGraphic
Api["unicode"]["IsLetter"] = unicode.IsLetter
Api["unicode"]["IsLower"] = unicode.IsLower
Api["unicode"]["IsMark"] = unicode.IsMark
Api["unicode"]["IsNumber"] = unicode.IsNumber
Api["unicode"]["IsOneOf"] = unicode.IsOneOf
Api["unicode"]["IsPrint"] = unicode.IsPrint
Api["unicode"]["IsPunct"] = unicode.IsPunct
Api["unicode"]["IsSpace"] = unicode.IsSpace
Api["unicode"]["IsSymbol"] = unicode.IsSymbol
Api["unicode"]["IsTitle"] = unicode.IsTitle
Api["unicode"]["IsUpper"] = unicode.IsUpper
Api["unicode"]["SimpleFold"] = unicode.SimpleFold
Api["unicode"]["To"] = unicode.To
Api["unicode"]["ToLower"] = unicode.ToLower
Api["unicode"]["ToTitle"] = unicode.ToTitle
Api["unicode"]["ToUpper"] = unicode.ToUpper
Api["unicode"]["CaseRange"] = unicode.CaseRange{}
Api["unicode"]["Range16"] = unicode.Range16{}
Api["unicode"]["Range32"] = unicode.Range32{}
Api["unicode"]["RangeTable"] = unicode.RangeTable{}
Api["unicode"]["AzeriCase"] = unicode.AzeriCase
Api["unicode"]["TurkishCase"] = unicode.TurkishCase
Api["unicode"]["In"] = unicode.In
Api["unicode"]["Is"] = unicode.Is
Api["unicode"]["IsControl"] = unicode.IsControl
Api["unicode"]["IsDigit"] = unicode.IsDigit
Api["unicode"]["IsGraphic"] = unicode.IsGraphic
Api["unicode"]["IsLetter"] = unicode.IsLetter
Api["unicode"]["IsLower"] = unicode.IsLower
Api["unicode"]["IsMark"] = unicode.IsMark
Api["unicode"]["IsNumber"] = unicode.IsNumber
Api["unicode"]["IsOneOf"] = unicode.IsOneOf
Api["unicode"]["IsPrint"] = unicode.IsPrint
Api["unicode"]["IsPunct"] = unicode.IsPunct
Api["unicode"]["IsSpace"] = unicode.IsSpace
Api["unicode"]["IsSymbol"] = unicode.IsSymbol
Api["unicode"]["IsTitle"] = unicode.IsTitle
Api["unicode"]["IsUpper"] = unicode.IsUpper
Api["unicode"]["SimpleFold"] = unicode.SimpleFold
Api["unicode"]["To"] = unicode.To
Api["unicode"]["ToLower"] = unicode.ToLower
Api["unicode"]["ToTitle"] = unicode.ToTitle
Api["unicode"]["ToUpper"] = unicode.ToUpper

}