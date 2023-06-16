package goquery
import (
github_com_PuerkitoBio_goquery "github.com/PuerkitoBio/goquery"
)
func Load(Interfaces ...interface{}) {
	if len(Interfaces) > 5 {
		if i, ok := Interfaces[0].(map[string]map[string]interface{}) {
			 DefaultGojaInterfaces: = i
			 if i, ok := Interfaces[1].(*goja.Runtime) {
				vm := i
				if i, ok := Interfaces[2].(*quickjs.Context) {
					ctx := i	
					if i, ok := Interfaces[3].(func(*quickjs.Context, interface{}) (quickjs.Value, string)) {
						ConvertInterfaceToQuickJS := i
						if i, ok := Interfaces[4].(func(interface{}, ...interface{}) interface{}) {
							ExecuteFunction := i
if _, ok := DefaultGojaInterfaces["goquery"]; !ok {
   DefaultGojaInterfaces["goquery"] = map[string]interface{}{}
}
DefaultGojaInterfaces["goquery"]["ToEnd"] = github_com_PuerkitoBio_goquery.ToEnd
DefaultGojaInterfaces["goquery"]["NodeName"] = github_com_PuerkitoBio_goquery.NodeName
DefaultGojaInterfaces["goquery"]["OuterHtml"] = github_com_PuerkitoBio_goquery.OuterHtml
DefaultGojaInterfaces["goquery"]["Render"] = github_com_PuerkitoBio_goquery.Render
DefaultGojaInterfaces["goquery"]["Document"] = github_com_PuerkitoBio_goquery.Document{}
DefaultGojaInterfaces["goquery"]["CloneDocument"] = github_com_PuerkitoBio_goquery.CloneDocument
DefaultGojaInterfaces["goquery"]["NewDocument"] = github_com_PuerkitoBio_goquery.NewDocument
DefaultGojaInterfaces["goquery"]["NewDocumentFromNode"] = github_com_PuerkitoBio_goquery.NewDocumentFromNode
DefaultGojaInterfaces["goquery"]["NewDocumentFromReader"] = github_com_PuerkitoBio_goquery.NewDocumentFromReader
DefaultGojaInterfaces["goquery"]["NewDocumentFromResponse"] = github_com_PuerkitoBio_goquery.NewDocumentFromResponse
DefaultGojaInterfaces["goquery"]["Single"] = github_com_PuerkitoBio_goquery.Single
DefaultGojaInterfaces["goquery"]["SingleMatcher"] = github_com_PuerkitoBio_goquery.SingleMatcher
DefaultGojaInterfaces["goquery"]["Selection"] = github_com_PuerkitoBio_goquery.Selection{}
DefaultGojaInterfaces["goquery"]["NodeName"] = github_com_PuerkitoBio_goquery.NodeName
DefaultGojaInterfaces["goquery"]["OuterHtml"] = github_com_PuerkitoBio_goquery.OuterHtml
DefaultGojaInterfaces["goquery"]["Render"] = github_com_PuerkitoBio_goquery.Render
DefaultGojaInterfaces["goquery"]["CloneDocument"] = github_com_PuerkitoBio_goquery.CloneDocument
DefaultGojaInterfaces["goquery"]["NewDocument"] = github_com_PuerkitoBio_goquery.NewDocument
DefaultGojaInterfaces["goquery"]["NewDocumentFromNode"] = github_com_PuerkitoBio_goquery.NewDocumentFromNode
DefaultGojaInterfaces["goquery"]["NewDocumentFromReader"] = github_com_PuerkitoBio_goquery.NewDocumentFromReader
DefaultGojaInterfaces["goquery"]["NewDocumentFromResponse"] = github_com_PuerkitoBio_goquery.NewDocumentFromResponse
DefaultGojaInterfaces["goquery"]["Single"] = github_com_PuerkitoBio_goquery.Single
DefaultGojaInterfaces["goquery"]["SingleMatcher"] = github_com_PuerkitoBio_goquery.SingleMatcher

			 			}	
			 		}
			 	}
			 }
		}
	}
}