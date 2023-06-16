package gomponents
import (
github_com_maragudk_gomponents "github.com/maragudk/gomponents"
)
func Load() {
if _, ok := DefaultGojaInterfaces["github.com/maragudk/gomponents"]; !ok {
   DefaultGojaInterfaces["github.com/maragudk/gomponents"] = map[string]interface{}{}
}
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["ElementType"] = github_com_maragudk_gomponents.ElementType
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["AttributeType"] = github_com_maragudk_gomponents.AttributeType
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Attr"] = github_com_maragudk_gomponents.Attr
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["El"] = github_com_maragudk_gomponents.El
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Group"] = github_com_maragudk_gomponents.Group
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["If"] = github_com_maragudk_gomponents.If
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Map"] = github_com_maragudk_gomponents.Map
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Raw"] = github_com_maragudk_gomponents.Raw
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Rawf"] = github_com_maragudk_gomponents.Rawf
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Text"] = github_com_maragudk_gomponents.Text
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Textf"] = github_com_maragudk_gomponents.Textf
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Attr"] = github_com_maragudk_gomponents.Attr
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["El"] = github_com_maragudk_gomponents.El
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Group"] = github_com_maragudk_gomponents.Group
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["If"] = github_com_maragudk_gomponents.If
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Map"] = github_com_maragudk_gomponents.Map
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Raw"] = github_com_maragudk_gomponents.Raw
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Rawf"] = github_com_maragudk_gomponents.Rawf
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Text"] = github_com_maragudk_gomponents.Text
DefaultGojaInterfaces["github.com/maragudk/gomponents"]["Textf"] = github_com_maragudk_gomponents.Textf
}
