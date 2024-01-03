package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

//Peer will encode connection information for a peer
type peer struct {
	IP   net.IP
	Port uint16
}

// Unmarshal will parse peer IP address and their ports from a buffer
func Unmarshal(buf []byte) ([]peer, error) {
    const peersize = 6 // 4 for IP and 2 for port
    numpeers := len(buf) / peersize
    if len(buf)%peersize != 0 {
        return nil, fmt.Errorf("incorrect length size received, length of %d received, when a multiple of %d is expected", len(buf), peersize)
    }
    peers := make([]peer, numpeers)
    for i := 0; i < numpeers; i++ {
        offset := i * peersize
        peers[i].IP = net.IP(buf[offset : offset+4])
        peers[i].Port = binary.BigEndian.Uint16(buf[offset+4: offset+6])
    }
    return peers, nil
}
func (p *peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}



