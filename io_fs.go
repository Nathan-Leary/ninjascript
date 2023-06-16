package ninjascript
import (
io_fs "io/fs"
)
func init() {if _, ok := Api["io/fs"]; !ok {
   Api["io/fs"] = map[string]interface{}{}
}
Api["io/fs"]["SkipAll"] = io_fs.SkipAll
Api["io/fs"]["SkipDir"] = io_fs.SkipDir
Api["io/fs"]["Glob"] = io_fs.Glob
Api["io/fs"]["ReadFile"] = io_fs.ReadFile
Api["io/fs"]["ValidPath"] = io_fs.ValidPath
Api["io/fs"]["WalkDir"] = io_fs.WalkDir
Api["io/fs"]["FileInfoToDirEntry"] = io_fs.FileInfoToDirEntry
Api["io/fs"]["ReadDir"] = io_fs.ReadDir
Api["io/fs"]["Sub"] = io_fs.Sub
Api["io/fs"]["Stat"] = io_fs.Stat
Api["io/fs"]["ModeDir"] = io_fs.ModeDir
Api["io/fs"]["ModeAppend"] = io_fs.ModeAppend
Api["io/fs"]["ModeExclusive"] = io_fs.ModeExclusive
Api["io/fs"]["ModeTemporary"] = io_fs.ModeTemporary
Api["io/fs"]["ModeSymlink"] = io_fs.ModeSymlink
Api["io/fs"]["ModeDevice"] = io_fs.ModeDevice
Api["io/fs"]["ModeNamedPipe"] = io_fs.ModeNamedPipe
Api["io/fs"]["ModeSocket"] = io_fs.ModeSocket
Api["io/fs"]["ModeSetuid"] = io_fs.ModeSetuid
Api["io/fs"]["ModeSetgid"] = io_fs.ModeSetgid
Api["io/fs"]["ModeCharDevice"] = io_fs.ModeCharDevice
Api["io/fs"]["ModeSticky"] = io_fs.ModeSticky
Api["io/fs"]["ModeIrregular"] = io_fs.ModeIrregular
Api["io/fs"]["ModeType"] = io_fs.ModeType
Api["io/fs"]["ModePerm"] = io_fs.ModePerm
Api["io/fs"]["PathError"] = io_fs.PathError{}
Api["io/fs"]["ReadDirFile"] = new(io_fs.ReadDirFile)
Api["io/fs"]["Glob"] = io_fs.Glob
Api["io/fs"]["ReadFile"] = io_fs.ReadFile
Api["io/fs"]["ValidPath"] = io_fs.ValidPath
Api["io/fs"]["WalkDir"] = io_fs.WalkDir
Api["io/fs"]["FileInfoToDirEntry"] = io_fs.FileInfoToDirEntry
Api["io/fs"]["ReadDir"] = io_fs.ReadDir
Api["io/fs"]["Sub"] = io_fs.Sub
Api["io/fs"]["Stat"] = io_fs.Stat

}