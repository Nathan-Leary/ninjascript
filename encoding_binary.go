package ninjascript
import (
encoding_binary "encoding/binary"
)
func init() {if _, ok := Api["encoding/binary"]; !ok {
   Api["encoding/binary"] = map[string]interface{}{}
}
Api["encoding/binary"]["MaxVarintLen16"] = encoding_binary.MaxVarintLen16
Api["encoding/binary"]["MaxVarintLen32"] = encoding_binary.MaxVarintLen32
Api["encoding/binary"]["MaxVarintLen64"] = encoding_binary.MaxVarintLen64
Api["encoding/binary"]["BigEndian"] = encoding_binary.BigEndian
Api["encoding/binary"]["LittleEndian"] = encoding_binary.LittleEndian
Api["encoding/binary"]["AppendUvarint"] = encoding_binary.AppendUvarint
Api["encoding/binary"]["AppendVarint"] = encoding_binary.AppendVarint
Api["encoding/binary"]["PutUvarint"] = encoding_binary.PutUvarint
Api["encoding/binary"]["PutVarint"] = encoding_binary.PutVarint
Api["encoding/binary"]["Read"] = encoding_binary.Read
Api["encoding/binary"]["ReadUvarint"] = encoding_binary.ReadUvarint
Api["encoding/binary"]["ReadVarint"] = encoding_binary.ReadVarint
Api["encoding/binary"]["Size"] = encoding_binary.Size
Api["encoding/binary"]["Uvarint"] = encoding_binary.Uvarint
Api["encoding/binary"]["Varint"] = encoding_binary.Varint
Api["encoding/binary"]["Write"] = encoding_binary.Write
Api["encoding/binary"]["AppendUvarint"] = encoding_binary.AppendUvarint
Api["encoding/binary"]["AppendVarint"] = encoding_binary.AppendVarint
Api["encoding/binary"]["PutUvarint"] = encoding_binary.PutUvarint
Api["encoding/binary"]["PutVarint"] = encoding_binary.PutVarint
Api["encoding/binary"]["Read"] = encoding_binary.Read
Api["encoding/binary"]["ReadUvarint"] = encoding_binary.ReadUvarint
Api["encoding/binary"]["ReadVarint"] = encoding_binary.ReadVarint
Api["encoding/binary"]["Size"] = encoding_binary.Size
Api["encoding/binary"]["Uvarint"] = encoding_binary.Uvarint
Api["encoding/binary"]["Varint"] = encoding_binary.Varint
Api["encoding/binary"]["Write"] = encoding_binary.Write

}