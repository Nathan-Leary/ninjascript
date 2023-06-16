package goquery
import (
github_com_PuerkitoBio_goquery "github.com/PuerkitoBio/goquery"
)
func Load() {
if _, ok := Api["github.com/PuerkitoBio/goquery"]; !ok {
   Api["github.com/PuerkitoBio/goquery"] = map[string]interface{}{}
}
Api["github.com/PuerkitoBio/goquery"]["ToEnd"] = github_com_PuerkitoBio_goquery.ToEnd
Api["github.com/PuerkitoBio/goquery"]["NodeName"] = github_com_PuerkitoBio_goquery.NodeName
Api["github.com/PuerkitoBio/goquery"]["OuterHtml"] = github_com_PuerkitoBio_goquery.OuterHtml
Api["github.com/PuerkitoBio/goquery"]["Render"] = github_com_PuerkitoBio_goquery.Render
Api["github.com/PuerkitoBio/goquery"]["Document"] = github_com_PuerkitoBio_goquery.Document{}
Api["github.com/PuerkitoBio/goquery"]["CloneDocument"] = github_com_PuerkitoBio_goquery.CloneDocument
Api["github.com/PuerkitoBio/goquery"]["NewDocument"] = github_com_PuerkitoBio_goquery.NewDocument
Api["github.com/PuerkitoBio/goquery"]["NewDocumentFromNode"] = github_com_PuerkitoBio_goquery.NewDocumentFromNode
Api["github.com/PuerkitoBio/goquery"]["NewDocumentFromReader"] = github_com_PuerkitoBio_goquery.NewDocumentFromReader
Api["github.com/PuerkitoBio/goquery"]["NewDocumentFromResponse"] = github_com_PuerkitoBio_goquery.NewDocumentFromResponse
Api["github.com/PuerkitoBio/goquery"]["Single"] = github_com_PuerkitoBio_goquery.Single
Api["github.com/PuerkitoBio/goquery"]["SingleMatcher"] = github_com_PuerkitoBio_goquery.SingleMatcher
Api["github.com/PuerkitoBio/goquery"]["Selection"] = github_com_PuerkitoBio_goquery.Selection{}
Api["github.com/PuerkitoBio/goquery"]["NodeName"] = github_com_PuerkitoBio_goquery.NodeName
Api["github.com/PuerkitoBio/goquery"]["OuterHtml"] = github_com_PuerkitoBio_goquery.OuterHtml
Api["github.com/PuerkitoBio/goquery"]["Render"] = github_com_PuerkitoBio_goquery.Render
Api["github.com/PuerkitoBio/goquery"]["CloneDocument"] = github_com_PuerkitoBio_goquery.CloneDocument
Api["github.com/PuerkitoBio/goquery"]["NewDocument"] = github_com_PuerkitoBio_goquery.NewDocument
Api["github.com/PuerkitoBio/goquery"]["NewDocumentFromNode"] = github_com_PuerkitoBio_goquery.NewDocumentFromNode
Api["github.com/PuerkitoBio/goquery"]["NewDocumentFromReader"] = github_com_PuerkitoBio_goquery.NewDocumentFromReader
Api["github.com/PuerkitoBio/goquery"]["NewDocumentFromResponse"] = github_com_PuerkitoBio_goquery.NewDocumentFromResponse
Api["github.com/PuerkitoBio/goquery"]["Single"] = github_com_PuerkitoBio_goquery.Single
Api["github.com/PuerkitoBio/goquery"]["SingleMatcher"] = github_com_PuerkitoBio_goquery.SingleMatcher
}
