package ninjascript
import (
fmt "fmt"
)
func init() {if _, ok := Api["fmt"]; !ok {
   Api["fmt"] = map[string]interface{}{}
}
Api["fmt"]["Append"] = fmt.Append
Api["fmt"]["Appendf"] = fmt.Appendf
Api["fmt"]["Appendln"] = fmt.Appendln
Api["fmt"]["Errorf"] = fmt.Errorf
Api["fmt"]["FormatString"] = fmt.FormatString
Api["fmt"]["Fprint"] = fmt.Fprint
Api["fmt"]["Fprintf"] = fmt.Fprintf
Api["fmt"]["Fprintln"] = fmt.Fprintln
Api["fmt"]["Fscan"] = fmt.Fscan
Api["fmt"]["Fscanf"] = fmt.Fscanf
Api["fmt"]["Fscanln"] = fmt.Fscanln
Api["fmt"]["Print"] = fmt.Print
Api["fmt"]["Printf"] = fmt.Printf
Api["fmt"]["Println"] = fmt.Println
Api["fmt"]["Scan"] = fmt.Scan
Api["fmt"]["Scanf"] = fmt.Scanf
Api["fmt"]["Scanln"] = fmt.Scanln
Api["fmt"]["Sprint"] = fmt.Sprint
Api["fmt"]["Sprintf"] = fmt.Sprintf
Api["fmt"]["Sprintln"] = fmt.Sprintln
Api["fmt"]["Sscan"] = fmt.Sscan
Api["fmt"]["Sscanf"] = fmt.Sscanf
Api["fmt"]["Sscanln"] = fmt.Sscanln
Api["fmt"]["Append"] = fmt.Append
Api["fmt"]["Appendf"] = fmt.Appendf
Api["fmt"]["Appendln"] = fmt.Appendln
Api["fmt"]["Errorf"] = fmt.Errorf
Api["fmt"]["FormatString"] = fmt.FormatString
Api["fmt"]["Fprint"] = fmt.Fprint
Api["fmt"]["Fprintf"] = fmt.Fprintf
Api["fmt"]["Fprintln"] = fmt.Fprintln
Api["fmt"]["Fscan"] = fmt.Fscan
Api["fmt"]["Fscanf"] = fmt.Fscanf
Api["fmt"]["Fscanln"] = fmt.Fscanln
Api["fmt"]["Print"] = fmt.Print
Api["fmt"]["Printf"] = fmt.Printf
Api["fmt"]["Println"] = fmt.Println
Api["fmt"]["Scan"] = fmt.Scan
Api["fmt"]["Scanf"] = fmt.Scanf
Api["fmt"]["Scanln"] = fmt.Scanln
Api["fmt"]["Sprint"] = fmt.Sprint
Api["fmt"]["Sprintf"] = fmt.Sprintf
Api["fmt"]["Sprintln"] = fmt.Sprintln
Api["fmt"]["Sscan"] = fmt.Sscan
Api["fmt"]["Sscanf"] = fmt.Sscanf
Api["fmt"]["Sscanln"] = fmt.Sscanln

}