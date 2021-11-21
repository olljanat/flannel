package ip

import (
    "fmt"
	"net"
)

type IP struct {
	IP4		  IP4
	IP6		  IP6
	IPv6Only  bool
}

type IPNet struct {
	IP4Net	  IP4Net
	IP6Net	  IP6Net
	IPv6Only  bool
}

func (n IPNet) StringSep(prefixSep string) string {
	if n.IPv6Only {
		return n.IP6Net.StringSep(":",fmt.Sprintf("%d",n.IP6Net.PrefixLen))
	}
	return n.IP6Net.StringSep(".",fmt.Sprintf("%d",n.IP4Net.PrefixLen))
}

func (n IPNet) Equal(other IPNet) bool {
	if n.IPv6Only {
		return n.IP6Net.Equal(other.IP6Net)
	}
	return n.IP4Net.Equal(other.IP4Net)
}

func (n IPNet) String() string {
	if n.IPv6Only {
		return n.IP6Net.String()
	}
	return n.IP4Net.String()
}

func (n IPNet) ToIPNet() *net.IPNet {
	if n.IPv6Only {
		return n.IP6Net.ToIPNet()
	}
	return n.IP4Net.ToIPNet()
}
