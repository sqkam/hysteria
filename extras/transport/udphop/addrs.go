package udphop

import (
	"github.com/apernet/hysteria/extras/v2/utils"
	"net"
	"strings"
)

// UDPHopAddr contains an IP address and a list of ports.
type UDPHopAddrs struct {
	IPs     []net.IP
	Ports   []uint16
	PortStr string
}

func (a *UDPHopAddrs) Network() string {
	return "udphopx"
}

func (a *UDPHopAddrs) String() string {
	var ips []string
	for _, v := range a.IPs {
		ips = append(ips, v.String())
	}
	return net.JoinHostPort(strings.Join(ips, ","), a.PortStr)
}

// Addrs returns a list of net.Addr's, one for each port.
func (a *UDPHopAddrs) Addrs() ([]net.Addr, error) {
	var addrs []net.Addr
	for _, ip := range a.IPs {
		for _, port := range a.Ports {
			addr := &net.UDPAddr{
				IP:   ip,
				Port: int(port),
			}
			addrs = append(addrs, addr)
		}
	}
	return addrs, nil
}
func ResolveUDPHopAddrs(host, port string, addrs []string) (Addrs, error) {
	ips := make([]net.IP, 0)
	for _, host := range addrs {
		ip, err := net.ResolveIPAddr("ip", strings.Trim(strings.Trim(host, "["), "]"))
		if err != nil {
			return nil, err
		}
		ips = append(ips, ip.IP)
	}

	result := &UDPHopAddrs{
		IPs:     ips,
		PortStr: port,
	}

	pu := utils.ParsePortUnion(port)
	if pu == nil {
		return nil, InvalidPortError{port}
	}
	result.Ports = pu.Ports()

	return result, nil
}