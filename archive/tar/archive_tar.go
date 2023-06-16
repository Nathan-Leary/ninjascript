package tar
import (
archive_tar "archive/tar"
)
func Load() {
if _, ok := Api["archive/tar"]; !ok {
   Api["archive/tar"] = map[string]interface{}{}
}
Api["archive/tar"]["TypeReg"] = archive_tar.TypeReg
Api["archive/tar"]["TypeRegA"] = archive_tar.TypeRegA
Api["archive/tar"]["TypeLink"] = archive_tar.TypeLink
Api["archive/tar"]["TypeSymlink"] = archive_tar.TypeSymlink
Api["archive/tar"]["TypeChar"] = archive_tar.TypeChar
Api["archive/tar"]["TypeBlock"] = archive_tar.TypeBlock
Api["archive/tar"]["TypeDir"] = archive_tar.TypeDir
Api["archive/tar"]["TypeFifo"] = archive_tar.TypeFifo
Api["archive/tar"]["TypeCont"] = archive_tar.TypeCont
Api["archive/tar"]["TypeXHeader"] = archive_tar.TypeXHeader
Api["archive/tar"]["TypeXGlobalHeader"] = archive_tar.TypeXGlobalHeader
Api["archive/tar"]["TypeGNUSparse"] = archive_tar.TypeGNUSparse
Api["archive/tar"]["TypeGNULongName"] = archive_tar.TypeGNULongName
Api["archive/tar"]["TypeGNULongLink"] = archive_tar.TypeGNULongLink
Api["archive/tar"]["FormatUnknown"] = archive_tar.FormatUnknown
Api["archive/tar"]["FormatUSTAR"] = archive_tar.FormatUSTAR
Api["archive/tar"]["FormatPAX"] = archive_tar.FormatPAX
Api["archive/tar"]["FormatGNU"] = archive_tar.FormatGNU
Api["archive/tar"]["FileInfoHeader"] = archive_tar.FileInfoHeader
Api["archive/tar"]["Reader"] = archive_tar.Reader{}
Api["archive/tar"]["NewReader"] = archive_tar.NewReader
Api["archive/tar"]["Writer"] = archive_tar.Writer{}
Api["archive/tar"]["NewWriter"] = archive_tar.NewWriter
Api["archive/tar"]["FileInfoHeader"] = archive_tar.FileInfoHeader
Api["archive/tar"]["NewReader"] = archive_tar.NewReader
Api["archive/tar"]["NewWriter"] = archive_tar.NewWriter
}
