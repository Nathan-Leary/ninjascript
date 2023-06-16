package gomponents
import (
github_com_maragudk_gomponents "github.com/maragudk/gomponents"
)
func Load() {
if _, ok := Api["github.com/maragudk/gomponents"]; !ok {
   Api["github.com/maragudk/gomponents"] = map[string]interface{}{}
}
Api["github.com/maragudk/gomponents"]["ElementType"] = github_com_maragudk_gomponents.ElementType
Api["github.com/maragudk/gomponents"]["AttributeType"] = github_com_maragudk_gomponents.AttributeType
Api["github.com/maragudk/gomponents"]["Attr"] = github_com_maragudk_gomponents.Attr
Api["github.com/maragudk/gomponents"]["El"] = github_com_maragudk_gomponents.El
Api["github.com/maragudk/gomponents"]["Group"] = github_com_maragudk_gomponents.Group
Api["github.com/maragudk/gomponents"]["If"] = github_com_maragudk_gomponents.If
Api["github.com/maragudk/gomponents"]["Map"] = github_com_maragudk_gomponents.Map
Api["github.com/maragudk/gomponents"]["Raw"] = github_com_maragudk_gomponents.Raw
Api["github.com/maragudk/gomponents"]["Rawf"] = github_com_maragudk_gomponents.Rawf
Api["github.com/maragudk/gomponents"]["Text"] = github_com_maragudk_gomponents.Text
Api["github.com/maragudk/gomponents"]["Textf"] = github_com_maragudk_gomponents.Textf
Api["github.com/maragudk/gomponents"]["Attr"] = github_com_maragudk_gomponents.Attr
Api["github.com/maragudk/gomponents"]["El"] = github_com_maragudk_gomponents.El
Api["github.com/maragudk/gomponents"]["Group"] = github_com_maragudk_gomponents.Group
Api["github.com/maragudk/gomponents"]["If"] = github_com_maragudk_gomponents.If
Api["github.com/maragudk/gomponents"]["Map"] = github_com_maragudk_gomponents.Map
Api["github.com/maragudk/gomponents"]["Raw"] = github_com_maragudk_gomponents.Raw
Api["github.com/maragudk/gomponents"]["Rawf"] = github_com_maragudk_gomponents.Rawf
Api["github.com/maragudk/gomponents"]["Text"] = github_com_maragudk_gomponents.Text
Api["github.com/maragudk/gomponents"]["Textf"] = github_com_maragudk_gomponents.Textf
}
