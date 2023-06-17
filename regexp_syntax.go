package ninjascript
import (
regexp_syntax "regexp/syntax"
)
func init() {if _, ok := Api["regexp/syntax"]; !ok {
   Api["regexp/syntax"] = map[string]interface{}{}
}
Api["regexp/syntax"]["IsWordChar"] = regexp_syntax.IsWordChar
Api["regexp/syntax"]["EmptyBeginLine"] = regexp_syntax.EmptyBeginLine
Api["regexp/syntax"]["EmptyEndLine"] = regexp_syntax.EmptyEndLine
Api["regexp/syntax"]["EmptyBeginText"] = regexp_syntax.EmptyBeginText
Api["regexp/syntax"]["EmptyEndText"] = regexp_syntax.EmptyEndText
Api["regexp/syntax"]["EmptyWordBoundary"] = regexp_syntax.EmptyWordBoundary
Api["regexp/syntax"]["EmptyNoWordBoundary"] = regexp_syntax.EmptyNoWordBoundary
Api["regexp/syntax"]["EmptyOpContext"] = regexp_syntax.EmptyOpContext
Api["regexp/syntax"]["Error"] = regexp_syntax.Error{}
Api["regexp/syntax"]["ErrInternalError"] = regexp_syntax.ErrInternalError
Api["regexp/syntax"]["ErrInvalidCharClass"] = regexp_syntax.ErrInvalidCharClass
Api["regexp/syntax"]["ErrInvalidCharRange"] = regexp_syntax.ErrInvalidCharRange
Api["regexp/syntax"]["ErrInvalidEscape"] = regexp_syntax.ErrInvalidEscape
Api["regexp/syntax"]["ErrInvalidNamedCapture"] = regexp_syntax.ErrInvalidNamedCapture
Api["regexp/syntax"]["ErrInvalidPerlOp"] = regexp_syntax.ErrInvalidPerlOp
Api["regexp/syntax"]["ErrInvalidRepeatOp"] = regexp_syntax.ErrInvalidRepeatOp
Api["regexp/syntax"]["ErrInvalidRepeatSize"] = regexp_syntax.ErrInvalidRepeatSize
Api["regexp/syntax"]["ErrInvalidUTF8"] = regexp_syntax.ErrInvalidUTF8
Api["regexp/syntax"]["ErrMissingBracket"] = regexp_syntax.ErrMissingBracket
Api["regexp/syntax"]["ErrMissingParen"] = regexp_syntax.ErrMissingParen
Api["regexp/syntax"]["ErrMissingRepeatArgument"] = regexp_syntax.ErrMissingRepeatArgument
Api["regexp/syntax"]["ErrTrailingBackslash"] = regexp_syntax.ErrTrailingBackslash
Api["regexp/syntax"]["ErrUnexpectedParen"] = regexp_syntax.ErrUnexpectedParen
Api["regexp/syntax"]["ErrNestingDepth"] = regexp_syntax.ErrNestingDepth
Api["regexp/syntax"]["ErrLarge"] = regexp_syntax.ErrLarge
Api["regexp/syntax"]["FoldCase"] = regexp_syntax.FoldCase
Api["regexp/syntax"]["Literal"] = regexp_syntax.Literal
Api["regexp/syntax"]["ClassNL"] = regexp_syntax.ClassNL
Api["regexp/syntax"]["DotNL"] = regexp_syntax.DotNL
Api["regexp/syntax"]["OneLine"] = regexp_syntax.OneLine
Api["regexp/syntax"]["NonGreedy"] = regexp_syntax.NonGreedy
Api["regexp/syntax"]["PerlX"] = regexp_syntax.PerlX
Api["regexp/syntax"]["UnicodeGroups"] = regexp_syntax.UnicodeGroups
Api["regexp/syntax"]["WasDollar"] = regexp_syntax.WasDollar
Api["regexp/syntax"]["Simple"] = regexp_syntax.Simple
Api["regexp/syntax"]["MatchNL"] = regexp_syntax.MatchNL
Api["regexp/syntax"]["Perl"] = regexp_syntax.Perl
Api["regexp/syntax"]["POSIX"] = regexp_syntax.POSIX
Api["regexp/syntax"]["Inst"] = regexp_syntax.Inst{}
Api["regexp/syntax"]["InstAlt"] = regexp_syntax.InstAlt
Api["regexp/syntax"]["InstAltMatch"] = regexp_syntax.InstAltMatch
Api["regexp/syntax"]["InstCapture"] = regexp_syntax.InstCapture
Api["regexp/syntax"]["InstEmptyWidth"] = regexp_syntax.InstEmptyWidth
Api["regexp/syntax"]["InstMatch"] = regexp_syntax.InstMatch
Api["regexp/syntax"]["InstFail"] = regexp_syntax.InstFail
Api["regexp/syntax"]["InstNop"] = regexp_syntax.InstNop
Api["regexp/syntax"]["InstRune"] = regexp_syntax.InstRune
Api["regexp/syntax"]["InstRune1"] = regexp_syntax.InstRune1
Api["regexp/syntax"]["InstRuneAny"] = regexp_syntax.InstRuneAny
Api["regexp/syntax"]["InstRuneAnyNotNL"] = regexp_syntax.InstRuneAnyNotNL
Api["regexp/syntax"]["OpNoMatch"] = regexp_syntax.OpNoMatch
Api["regexp/syntax"]["OpEmptyMatch"] = regexp_syntax.OpEmptyMatch
Api["regexp/syntax"]["OpLiteral"] = regexp_syntax.OpLiteral
Api["regexp/syntax"]["OpCharClass"] = regexp_syntax.OpCharClass
Api["regexp/syntax"]["OpAnyCharNotNL"] = regexp_syntax.OpAnyCharNotNL
Api["regexp/syntax"]["OpAnyChar"] = regexp_syntax.OpAnyChar
Api["regexp/syntax"]["OpBeginLine"] = regexp_syntax.OpBeginLine
Api["regexp/syntax"]["OpEndLine"] = regexp_syntax.OpEndLine
Api["regexp/syntax"]["OpBeginText"] = regexp_syntax.OpBeginText
Api["regexp/syntax"]["OpEndText"] = regexp_syntax.OpEndText
Api["regexp/syntax"]["OpWordBoundary"] = regexp_syntax.OpWordBoundary
Api["regexp/syntax"]["OpNoWordBoundary"] = regexp_syntax.OpNoWordBoundary
Api["regexp/syntax"]["OpCapture"] = regexp_syntax.OpCapture
Api["regexp/syntax"]["OpStar"] = regexp_syntax.OpStar
Api["regexp/syntax"]["OpPlus"] = regexp_syntax.OpPlus
Api["regexp/syntax"]["OpQuest"] = regexp_syntax.OpQuest
Api["regexp/syntax"]["OpRepeat"] = regexp_syntax.OpRepeat
Api["regexp/syntax"]["OpConcat"] = regexp_syntax.OpConcat
Api["regexp/syntax"]["OpAlternate"] = regexp_syntax.OpAlternate
Api["regexp/syntax"]["Prog"] = regexp_syntax.Prog{}
Api["regexp/syntax"]["Compile"] = regexp_syntax.Compile
Api["regexp/syntax"]["Regexp"] = regexp_syntax.Regexp{}
Api["regexp/syntax"]["Parse"] = regexp_syntax.Parse
Api["regexp/syntax"]["IsWordChar"] = regexp_syntax.IsWordChar
Api["regexp/syntax"]["EmptyOpContext"] = regexp_syntax.EmptyOpContext
Api["regexp/syntax"]["Compile"] = regexp_syntax.Compile
Api["regexp/syntax"]["Parse"] = regexp_syntax.Parse

}