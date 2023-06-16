package ninjascript
import (
text_tabwriter "text/tabwriter"
)
func init() {if _, ok := Api["text/tabwriter"]; !ok {
   Api["text/tabwriter"] = map[string]interface{}{}
}
Api["text/tabwriter"]["FilterHTML"] = text_tabwriter.FilterHTML
Api["text/tabwriter"]["StripEscape"] = text_tabwriter.StripEscape
Api["text/tabwriter"]["AlignRight"] = text_tabwriter.AlignRight
Api["text/tabwriter"]["DiscardEmptyColumns"] = text_tabwriter.DiscardEmptyColumns
Api["text/tabwriter"]["TabIndent"] = text_tabwriter.TabIndent
Api["text/tabwriter"]["Debug"] = text_tabwriter.Debug
Api["text/tabwriter"]["Writer"] = text_tabwriter.Writer{}
Api["text/tabwriter"]["NewWriter"] = text_tabwriter.NewWriter
Api["text/tabwriter"]["NewWriter"] = text_tabwriter.NewWriter

}