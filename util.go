// canvas
package ninjascript

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"

	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	// "github.com/buke/quickjs-go"
	// "github.com/dop251/goja"
	"github.com/spf13/cast"
)

var gonum int = 0
var mut = sync.Mutex{}

func init() {

	if _, ok := Api["util"]; !ok {
		Api["util"] = map[string]interface{}{}
	}

	vmSet := func(txt string, fn interface{}) {
		// vm.Set(txt, fn)
		Api["util"][txt] = fn
	}
	//	Vm.RunString("var Util = {};")
	//	Vm.RunString("Util = {};")

	vmSet("Stop", func() {
		var c chan bool
		<-c
	})

	vmSet("Pause", func() {
		var c chan bool
		<-c
	})

	vmSet("Freeze", func() {
		var c chan bool
		<-c
	})

	vmSet("Go", func(fn func(...interface{})) func(...interface{}) {
		// goroutines = append(goroutines, fn)
		go fn()
		return fn
	})

	vmSet("Jump", func(fn func(...interface{})) func(...interface{}) {
		// goroutines = append(goroutines, fn)
		go fn()
		return fn
	})

	//	//Vm.RunString("Util['Reflect'] = Reflect;")

	vmSet("func", func(fn string) func() {
		return func() {
			Vm.RunString("(" + fn + ")()")
		}
	})

	//	//Vm.RunString("Util['func'] = func;")
	//
	vmSet("nil", nil)
	//	//Vm.RunString("Util['nil'] = nil;")
	vmSet("IsWindows", func() bool {
		return runtime.GOOS == "windows"
	})
	//Vm.RunString("Util['IsWindows'] = IsWindows;")

	vmSet("IsMac", func() bool {
		return runtime.GOOS == "darwin"
	})

	//Vm.RunString("Util['IsMac'] = IsMac;")

	vmSet("IsLinux", func() bool {
		return runtime.GOOS == "linux"
	})
	//Vm.RunString("Util['IsLinux'] = IsLinux;")

	vmSet("Getwd", os.Getwd)
	//Vm.RunString("Util['Getwd'] = Getwd;")
	vmSet("Chmod", os.Chmod)
	//Vm.RunString("Util['Chmod'] = Chmod;")
	vmSet("Chown", os.Chown)
	//Vm.RunString("Util['Chown'] = Chown;")
	vmSet("Chdir", os.Chdir)
	//Vm.RunString("Util['Chdir'] = Chdir;")
	vmSet("Create", os.Create)
	//Vm.RunString("Util['Create'] = Create;")
	vmSet("CreateTemp", os.CreateTemp)
	//Vm.RunString("Util['CreateTemp'] = CreateTemp;")
	vmSet("Executable", os.Executable)
	//Vm.RunString("Util['Executable'] = Executable;")
	vmSet("Exit", func(args interface{}) {
		// if dash != nil {
		// dash.Close()
		// }
		if a, ok := args.(float64); ok {
			os.Exit(int(a))
		} else {
			os.Exit(0)
		}

	})
	//Vm.RunString("Util['Exit'] = Exit;")

	vmSet("FindProcess", os.FindProcess)
	//Vm.RunString("Util['FindProcess'] = FindProcess;")
	vmSet("Readlink", os.Readlink)
	//Vm.RunString("Util['Readlink'] = Readlink;")
	vmSet("Link", os.Link)
	//Vm.RunString("Util['Link'] = Link;")
	vmSet("TempDir", os.TempDir)
	//Vm.RunString("Util['TempDir'] = TempDir;")
	vmSet("UserHomeDir", os.UserHomeDir)
	//Vm.RunString("Util['UserHomeDir'] = UserHomeDir;")
	vmSet("UserCacheDir", os.UserCacheDir)
	//Vm.RunString("Util['UserCacheDir'] = UserCacheDir;")
	vmSet("Symlink", os.Symlink)
	//Vm.RunString("Util['Symlink'] = Symlink;")
	vmSet("Rename", os.Rename)
	//Vm.RunString("Util['Rename'] = Rename;")
	vmSet("Remove", os.Remove)
	//Vm.RunString("Util['Remove'] = Remove;")
	vmSet("RemoveAll", os.RemoveAll)
	//Vm.RunString("Util['RemoveAll'] = RemoveAll;")
	vmSet("Mkdir", os.Mkdir)
	//Vm.RunString("Util['Mkdir'] = Mkdir;")

	vmSet("MkdirAll", os.MkdirAll)
	//Vm.RunString("Util['MkdirAll'] = MkdirAll;")
	vmSet("MkdirTemp", os.MkdirTemp)
	//Vm.RunString("Util['MkdirTemp'] = MkdirTemp;")
	vmSet("Getpid", os.Getpid)
	//Vm.RunString("Util['Getpid'] = Getpid;")
	vmSet("Hostname", os.Hostname)
	//Vm.RunString("Util['Hostname'] = Hostname;")

	vmSet("Shell", func(s ...string) []byte {

		if len(s) == 0 {
			return nil
		}
		if runtime.GOOS == "windows" {
			if len(s) > 1 {
				s = append([]string{"/c"}, s...)
				cmd := exec.Command("cmd", s...)
				out, _ := cmd.Output()

				return out
			}

			cmd := exec.Command("cmd", "/c", s[0])
			out, _ := cmd.Output()

			return out
		}
		if len(s) > 1 {
			cmd := exec.Command(s[0], s[1:]...)
			out, _ := cmd.Output()
			return out
		} else {
			shellCmd := strings.Split(s[0], `\\`)
			shellCmd[0] = strings.Join(shellCmd, "{{escaped-slash}}")
			shellCmd = strings.Split(shellCmd[0], `\ `)
			shellCmd[0] = strings.Join(shellCmd, "{{escaped-space}}")

			shellCmd = strings.Split(shellCmd[0], " ")
			for k, v := range shellCmd {
				shellCmd[k] = strings.Replace(v, "{{escaped-space}}", `\ `, -1)
				shellCmd[k] = strings.Replace(shellCmd[k], "{{escaped-slash}}", `\\`, -1)
			}

			cmd := exec.Command(shellCmd[0], shellCmd[1:]...)
			out, _ := cmd.Output()
			return out
		}
	})

	//Vm.RunString("Util['Shell'] = Shell;")

	vmSet("Int32", func(i float64) int32 {
		return int32(int(i))
	})

	//Vm.RunString("Util['Int32'] = Int32;")
	vmSet("UintPtr", func(i float64) uintptr {
		return uintptr(int(i))
	})
	//Vm.RunString("Util['UintPtr'] = UintPtr;")

	vmSet("ExePath", func() string {
		return os.Args[0]
	})

	//Vm.RunString("Util['ExePath'] = ExePath;")
	vmSet("ExeDir", func() string {
		return filepath.Dir(os.Args[0])
	})
	//Vm.RunString("Util['ExeDir'] = ExeDir;")

	vmSet("Dir", func(s string) string {
		return filepath.Dir(s)
	})
	//Vm.RunString("Util['Dir'] = Dir;")

	vmSet("WriteFile", func(s string, b []byte, i int) error {
		return ioutil.WriteFile(s, b, fs.FileMode(i))
	})

	//Vm.RunString("Util['WriteFile'] = WriteFile;")
	vmSet("ReadFile", func(s string) []byte {
		if strings.Contains(s, "://") {
			resp, err := http.Get(s)
			if err != nil {
				// handle error
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)

			return body

		}
		b, _ := ioutil.ReadFile(s)
		return b
	})
	//Vm.RunString("Util['ReadFile'] = ReadFile;")
	vmSet("ListFiles", func(s string) []string {
		files, _ := ioutil.ReadDir(s)
		sa := []string{}
		for _, file := range files {
			if !file.IsDir() {
				sa = append(sa, file.Name())
			}
		}

		return sa
	})

	//Vm.RunString("Util['ListFiles'] = ListFiles;")

	vmSet("ListFilesAndDirs", func(s string) []string {
		files, _ := ioutil.ReadDir(s)
		sa := []string{}
		for _, file := range files {
			sa = append(sa, file.Name())
		}

		return sa
	})

	//Vm.RunString("Util['ListFilesAndDirs'] = ListFilesAndDirs;")
	vmSet("ListDirs", func(s string) []string {
		files, _ := ioutil.ReadDir(s)
		sa := []string{}
		for _, file := range files {
			if file.IsDir() {
				sa = append(sa, file.Name())
			}
		}

		return sa
	})
	//Vm.RunString("Util['ListDirs'] = ListDirs;")

	vmSet("FileMode", func(s string) int64 {

		fi, err := os.Stat(s)
		if err != nil {
			return 0
		}
		// get the size
		mode := fi.Mode()

		return int64(mode)
	})

	//Vm.RunString("Util['FileMode'] = FileMode;")

	vmSet("FileIsDir", func(s string) bool {

		fi, err := os.Stat(s)
		if err != nil {
			return false
		}
		// get the size
		IsDir := fi.IsDir()

		return IsDir
	})

	//Vm.RunString("Util['FileIsDir'] = FileIsDir;")

	//Vm.RunString("Util['FileExists'] = FileExists;")
	vmSet("FileName", func(s string) string {

		fi, err := os.Stat(s)
		if err != nil {
			return ""
		}

		Name := fi.Name()

		return Name
	})
	//Vm.RunString("Util['FileName'] = FileName;")

	vmSet("FileSize", func(s string) time.Time {

		fi, err := os.Stat(s)
		if err != nil {
			return time.Time{}
		}
		// get the size
		ModTime := fi.ModTime()

		return ModTime
	})

	//Vm.RunString("Util['FileSize'] = FileSize;")

	vmSet("args", os.Args)
	//Vm.RunString("Util['Args'] = args;")
	vmSet("print", fmt.Print)
	//Vm.RunString("Util['Print'] = print;")
	vmSet("println", fmt.Println)
	//Vm.RunString("Util['Println'] = println;")
	vmSet("StringToByte", func(s string) []byte {
		return []byte(s)
	})
	//Vm.RunString("Util['StringToByte'] = StringToByte;")
	vmSet("Time", func(i interface{}) time.Time {
		return cast.ToTime(i)
	})
	//Vm.RunString("Util['Time'] = Time;")
	vmSet("Bool", func(i interface{}) bool {
		return cast.ToBool(i)
	})
	//Vm.RunString("Util['Bool'] = Bool;")
	vmSet("Float32", func(i interface{}) float32 {
		return cast.ToFloat32(i)
	})
	//Vm.RunString("Util['Float32'] = Float32;")

	vmSet("Float", func(i interface{}) float64 {
		return cast.ToFloat64(i)
	})
	//Vm.RunString("Util['Float'] = Float;")
	vmSet("Int", func(i interface{}) int {
		return cast.ToInt(i)
	})

	//Vm.RunString("Util['Int'] = Int;")
	vmSet("Byte", func(i interface{}) []byte {
		if v, ok := i.([]byte); ok {
			return v
		}

		return []byte(cast.ToString(i))
	})
	//Vm.RunString("Util['Bytes'] = Bytes;")

	vmSet("String", func(i interface{}) string {
		return cast.ToString(i)
	})
	//Vm.RunString("Util['String'] = String;")

	vmSet("ByteToString", func(b []byte) string {
		return string(b)
	})
	//Vm.RunString("Util['ByteToString'] = ByteToString;")

	vmSet("Sleep", func(t float64) {
		time.Sleep(time.Millisecond * time.Duration(t))
	})
	//Vm.RunString("Util['Sleep'] = Sleep;")
	vmSet("Println", fmt.Println)
	//Vm.RunString("Util['Println'] = Println;")
	vmSet("Print", fmt.Print)
	//Vm.RunString("Util['Print'] = Print;")
	vmSet("Sprint", fmt.Sprint)
	//Vm.RunString("Util['Sprint'] = Sprint;")
	vmSet("Sprintln", fmt.Sprintln)
	//Vm.RunString("Util['Sprintln'] = Sprintln;")

}
