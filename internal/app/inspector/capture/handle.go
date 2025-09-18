package capture

import (
	"fmt"
	"time"

	"github.com/google/gopacket/pcap"
)

func Start(iface string, promisc bool, filterString string, len int, timeout time.Duration) (*pcap.Handle, error) {
	handle, err := pcap.OpenLive(iface, int32(len), promisc, timeout)
	if err != nil {
		return nil, fmt.Errorf("opening device: %v", err)
	}

	err = handle.SetBPFFilter(filterString)
	if err != nil {
		return nil, fmt.Errorf("setting BPF filter: %v", err)
	}

	return handle, nil
}
