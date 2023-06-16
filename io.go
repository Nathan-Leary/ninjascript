package ninjascript
import (
io "io"
)
func init() {if _, ok := Api["io"]; !ok {
   Api["io"] = map[string]interface{}{}
}
Api["io"]["SeekStart"] = io.SeekStart
Api["io"]["SeekCurrent"] = io.SeekCurrent
Api["io"]["SeekEnd"] = io.SeekEnd
Api["io"]["EOF"] = io.EOF
Api["io"]["ErrNoProgress"] = io.ErrNoProgress
Api["io"]["ErrShortBuffer"] = io.ErrShortBuffer
Api["io"]["ErrShortWrite"] = io.ErrShortWrite
Api["io"]["ErrUnexpectedEOF"] = io.ErrUnexpectedEOF
Api["io"]["Copy"] = io.Copy
Api["io"]["CopyBuffer"] = io.CopyBuffer
Api["io"]["CopyN"] = io.CopyN
Api["io"]["Pipe"] = io.Pipe
Api["io"]["ReadAll"] = io.ReadAll
Api["io"]["ReadAtLeast"] = io.ReadAtLeast
Api["io"]["ReadFull"] = io.ReadFull
Api["io"]["WriteString"] = io.WriteString
Api["io"]["LimitedReader"] = io.LimitedReader{}
Api["io"]["OffsetWriter"] = io.OffsetWriter{}
Api["io"]["NewOffsetWriter"] = io.NewOffsetWriter
Api["io"]["PipeReader"] = io.PipeReader{}
Api["io"]["PipeWriter"] = io.PipeWriter{}
Api["io"]["NopCloser"] = io.NopCloser
Api["io"]["LimitReader"] = io.LimitReader
Api["io"]["MultiReader"] = io.MultiReader
Api["io"]["TeeReader"] = io.TeeReader
Api["io"]["SectionReader"] = io.SectionReader{}
Api["io"]["NewSectionReader"] = io.NewSectionReader
Api["io"]["Discard"] = io.Discard
Api["io"]["MultiWriter"] = io.MultiWriter
Api["io"]["Copy"] = io.Copy
Api["io"]["CopyBuffer"] = io.CopyBuffer
Api["io"]["CopyN"] = io.CopyN
Api["io"]["Pipe"] = io.Pipe
Api["io"]["ReadAll"] = io.ReadAll
Api["io"]["ReadAtLeast"] = io.ReadAtLeast
Api["io"]["ReadFull"] = io.ReadFull
Api["io"]["WriteString"] = io.WriteString
Api["io"]["NewOffsetWriter"] = io.NewOffsetWriter
Api["io"]["NopCloser"] = io.NopCloser
Api["io"]["LimitReader"] = io.LimitReader
Api["io"]["MultiReader"] = io.MultiReader
Api["io"]["TeeReader"] = io.TeeReader
Api["io"]["NewSectionReader"] = io.NewSectionReader
Api["io"]["MultiWriter"] = io.MultiWriter

}