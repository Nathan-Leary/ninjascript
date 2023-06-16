package goquery
import (
github_com_PuerkitoBio_goquery "github.com/PuerkitoBio/goquery"
)
func Load() {
if _, ok := DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]; !ok {
   DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"] = map[string]interface{}{}
}
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["ToEnd"] = github_com_PuerkitoBio_goquery.ToEnd
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["NodeName"] = github_com_PuerkitoBio_goquery.NodeName
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["OuterHtml"] = github_com_PuerkitoBio_goquery.OuterHtml
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["Render"] = github_com_PuerkitoBio_goquery.Render
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["Document"] = github_com_PuerkitoBio_goquery.Document{}
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["CloneDocument"] = github_com_PuerkitoBio_goquery.CloneDocument
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["NewDocument"] = github_com_PuerkitoBio_goquery.NewDocument
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["NewDocumentFromNode"] = github_com_PuerkitoBio_goquery.NewDocumentFromNode
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["NewDocumentFromReader"] = github_com_PuerkitoBio_goquery.NewDocumentFromReader
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["NewDocumentFromResponse"] = github_com_PuerkitoBio_goquery.NewDocumentFromResponse
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["Single"] = github_com_PuerkitoBio_goquery.Single
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["SingleMatcher"] = github_com_PuerkitoBio_goquery.SingleMatcher
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["Selection"] = github_com_PuerkitoBio_goquery.Selection{}
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["NodeName"] = github_com_PuerkitoBio_goquery.NodeName
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["OuterHtml"] = github_com_PuerkitoBio_goquery.OuterHtml
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["Render"] = github_com_PuerkitoBio_goquery.Render
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["CloneDocument"] = github_com_PuerkitoBio_goquery.CloneDocument
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["NewDocument"] = github_com_PuerkitoBio_goquery.NewDocument
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["NewDocumentFromNode"] = github_com_PuerkitoBio_goquery.NewDocumentFromNode
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["NewDocumentFromReader"] = github_com_PuerkitoBio_goquery.NewDocumentFromReader
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["NewDocumentFromResponse"] = github_com_PuerkitoBio_goquery.NewDocumentFromResponse
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["Single"] = github_com_PuerkitoBio_goquery.Single
DefaultGojaInterfaces["github.com/PuerkitoBio/goquery"]["SingleMatcher"] = github_com_PuerkitoBio_goquery.SingleMatcher
}
