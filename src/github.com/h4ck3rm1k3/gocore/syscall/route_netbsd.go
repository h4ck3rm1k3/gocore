// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import "github.com/h4ck3rm1k3/gocore/unsafe"

func (any *anyMessage) toRoutingMessage(b []byte) RoutingMessage {
	switch any.Type {
	case RTM_ADD, RTM_DELETE, RTM_CHANGE, RTM_GET, RTM_LOSING, RTM_REDIRECT, RTM_MISS, RTM_LOCK, RTM_RESOLVE:
		p := (*RouteMessage)(unsafe.Pointer(any))
		// We don't support sockaddr_mpls for now.
		p.Header.Addrs &= RTA_DST | RTA_GATEWAY | RTA_NETMASK | RTA_GENMASK | RTA_IFA | RTA_IFP | RTA_BRD | RTA_AUTHOR
		return &RouteMessage{Header: p.Header, Data: b[SizeofRtMsghdr:any.Msglen]}
	case RTM_IFINFO:
		p := (*InterfaceMessage)(unsafe.Pointer(any))
		return &InterfaceMessage{Header: p.Header, Data: b[SizeofIfMsghdr:any.Msglen]}
	case RTM_IFANNOUNCE:
		p := (*InterfaceAnnounceMessage)(unsafe.Pointer(any))
		return &InterfaceAnnounceMessage{Header: p.Header}
	case RTM_NEWADDR, RTM_DELADDR:
		p := (*InterfaceAddrMessage)(unsafe.Pointer(any))
		return &InterfaceAddrMessage{Header: p.Header, Data: b[SizeofIfaMsghdr:any.Msglen]}
	}
	return nil
}

// InterfaceAnnounceMessage represents a routing message containing
// network interface arrival and departure information.
type InterfaceAnnounceMessage struct {
	Header IfAnnounceMsghdr
}

func (m *InterfaceAnnounceMessage) sockaddr() ([]Sockaddr, error) { return nil, nil }
