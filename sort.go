package ninjascript
import (
sort "sort"
)
func init() {if _, ok := Api["sort"]; !ok {
   Api["sort"] = map[string]interface{}{}
}
Api["sort"]["Find"] = sort.Find
Api["sort"]["Float64s"] = sort.Float64s
Api["sort"]["Float64sAreSorted"] = sort.Float64sAreSorted
Api["sort"]["Ints"] = sort.Ints
Api["sort"]["IntsAreSorted"] = sort.IntsAreSorted
Api["sort"]["IsSorted"] = sort.IsSorted
Api["sort"]["Search"] = sort.Search
Api["sort"]["SearchFloat64s"] = sort.SearchFloat64s
Api["sort"]["SearchInts"] = sort.SearchInts
Api["sort"]["SearchStrings"] = sort.SearchStrings
Api["sort"]["Slice"] = sort.Slice
Api["sort"]["SliceIsSorted"] = sort.SliceIsSorted
Api["sort"]["SliceStable"] = sort.SliceStable
Api["sort"]["Sort"] = sort.Sort
Api["sort"]["Stable"] = sort.Stable
Api["sort"]["Strings"] = sort.Strings
Api["sort"]["StringsAreSorted"] = sort.StringsAreSorted
Api["sort"]["Reverse"] = sort.Reverse
Api["sort"]["Find"] = sort.Find
Api["sort"]["Float64s"] = sort.Float64s
Api["sort"]["Float64sAreSorted"] = sort.Float64sAreSorted
Api["sort"]["Ints"] = sort.Ints
Api["sort"]["IntsAreSorted"] = sort.IntsAreSorted
Api["sort"]["IsSorted"] = sort.IsSorted
Api["sort"]["Search"] = sort.Search
Api["sort"]["SearchFloat64s"] = sort.SearchFloat64s
Api["sort"]["SearchInts"] = sort.SearchInts
Api["sort"]["SearchStrings"] = sort.SearchStrings
Api["sort"]["Slice"] = sort.Slice
Api["sort"]["SliceIsSorted"] = sort.SliceIsSorted
Api["sort"]["SliceStable"] = sort.SliceStable
Api["sort"]["Sort"] = sort.Sort
Api["sort"]["Stable"] = sort.Stable
Api["sort"]["Strings"] = sort.Strings
Api["sort"]["StringsAreSorted"] = sort.StringsAreSorted
Api["sort"]["Reverse"] = sort.Reverse

}