package ninjascript
import (
net_netip "net/netip"
)
func init() {if _, ok := Api["net/netip"]; !ok {
   Api["net/netip"] = map[string]interface{}{}
}
Api["net/netip"]["Addr"] = net_netip.Addr{}
Api["net/netip"]["AddrFrom16"] = net_netip.AddrFrom16
Api["net/netip"]["AddrFrom4"] = net_netip.AddrFrom4
Api["net/netip"]["AddrFromSlice"] = net_netip.AddrFromSlice
Api["net/netip"]["IPv4Unspecified"] = net_netip.IPv4Unspecified
Api["net/netip"]["IPv6LinkLocalAllNodes"] = net_netip.IPv6LinkLocalAllNodes
Api["net/netip"]["IPv6LinkLocalAllRouters"] = net_netip.IPv6LinkLocalAllRouters
Api["net/netip"]["IPv6Loopback"] = net_netip.IPv6Loopback
Api["net/netip"]["IPv6Unspecified"] = net_netip.IPv6Unspecified
Api["net/netip"]["MustParseAddr"] = net_netip.MustParseAddr
Api["net/netip"]["ParseAddr"] = net_netip.ParseAddr
Api["net/netip"]["AddrPort"] = net_netip.AddrPort{}
Api["net/netip"]["AddrPortFrom"] = net_netip.AddrPortFrom
Api["net/netip"]["MustParseAddrPort"] = net_netip.MustParseAddrPort
Api["net/netip"]["ParseAddrPort"] = net_netip.ParseAddrPort
Api["net/netip"]["Prefix"] = net_netip.Prefix{}
Api["net/netip"]["MustParsePrefix"] = net_netip.MustParsePrefix
Api["net/netip"]["ParsePrefix"] = net_netip.ParsePrefix
Api["net/netip"]["PrefixFrom"] = net_netip.PrefixFrom
Api["net/netip"]["AddrFrom16"] = net_netip.AddrFrom16
Api["net/netip"]["AddrFrom4"] = net_netip.AddrFrom4
Api["net/netip"]["AddrFromSlice"] = net_netip.AddrFromSlice
Api["net/netip"]["IPv4Unspecified"] = net_netip.IPv4Unspecified
Api["net/netip"]["IPv6LinkLocalAllNodes"] = net_netip.IPv6LinkLocalAllNodes
Api["net/netip"]["IPv6LinkLocalAllRouters"] = net_netip.IPv6LinkLocalAllRouters
Api["net/netip"]["IPv6Loopback"] = net_netip.IPv6Loopback
Api["net/netip"]["IPv6Unspecified"] = net_netip.IPv6Unspecified
Api["net/netip"]["MustParseAddr"] = net_netip.MustParseAddr
Api["net/netip"]["ParseAddr"] = net_netip.ParseAddr
Api["net/netip"]["AddrPortFrom"] = net_netip.AddrPortFrom
Api["net/netip"]["MustParseAddrPort"] = net_netip.MustParseAddrPort
Api["net/netip"]["ParseAddrPort"] = net_netip.ParseAddrPort
Api["net/netip"]["MustParsePrefix"] = net_netip.MustParsePrefix
Api["net/netip"]["ParsePrefix"] = net_netip.ParsePrefix
Api["net/netip"]["PrefixFrom"] = net_netip.PrefixFrom

}