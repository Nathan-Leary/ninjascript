package zip
import (
archive_zip "archive/zip"
)
func Load() {
if _, ok := Api["archive/zip"]; !ok {
   Api["archive/zip"] = map[string]interface{}{}
}
Api["archive/zip"]["Store"] = archive_zip.Store
Api["archive/zip"]["Deflate"] = archive_zip.Deflate
Api["archive/zip"]["RegisterCompressor"] = archive_zip.RegisterCompressor
Api["archive/zip"]["RegisterDecompressor"] = archive_zip.RegisterDecompressor
Api["archive/zip"]["File"] = archive_zip.File{}
Api["archive/zip"]["FileHeader"] = archive_zip.FileHeader{}
Api["archive/zip"]["FileInfoHeader"] = archive_zip.FileInfoHeader
Api["archive/zip"]["ReadCloser"] = archive_zip.ReadCloser{}
Api["archive/zip"]["OpenReader"] = archive_zip.OpenReader
Api["archive/zip"]["Reader"] = archive_zip.Reader{}
Api["archive/zip"]["NewReader"] = archive_zip.NewReader
Api["archive/zip"]["Writer"] = archive_zip.Writer{}
Api["archive/zip"]["NewWriter"] = archive_zip.NewWriter
Api["archive/zip"]["RegisterCompressor"] = archive_zip.RegisterCompressor
Api["archive/zip"]["RegisterDecompressor"] = archive_zip.RegisterDecompressor
Api["archive/zip"]["FileInfoHeader"] = archive_zip.FileInfoHeader
Api["archive/zip"]["OpenReader"] = archive_zip.OpenReader
Api["archive/zip"]["NewReader"] = archive_zip.NewReader
Api["archive/zip"]["NewWriter"] = archive_zip.NewWriter
}
