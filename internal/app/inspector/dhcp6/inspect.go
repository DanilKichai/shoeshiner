package dhcp6

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const FilterString = "ip6 and udp and src port 547 and dst port 546"

func Prepare(prefix string) error {
	if err := os.MkdirAll(prefix, 0755); err != nil {
		return fmt.Errorf("prepare directory structure: %v", err)
	}

	if err := syscall.Mkfifo(filepath.Join(prefix, "trigger"), 0600); err != nil {
		return fmt.Errorf("make trigger pipe: %v", err)
	}

	return nil
}

func Inspect(packet gopacket.Packet, prefix string) (bool, error) {
	if dhcp6Layer := packet.Layer(layers.LayerTypeDHCPv6); dhcp6Layer != nil {
		dhcp6, _ := dhcp6Layer.(*layers.DHCPv6)

		if dhcp6.MsgType != layers.DHCPv6MsgTypeReply {
			return false, nil
		}

		if err := os.MkdirAll(filepath.Join(prefix, "options"), 0755); err != nil {
			return false, fmt.Errorf("create directory structure: %v", err)
		}

		for _, option := range dhcp6.Options {
			if err := os.WriteFile(filepath.Join(prefix, "options", strconv.Itoa(int(option.Code))), option.Data, 0644); err != nil {
				return false, fmt.Errorf("write option data file: %v", err)
			}
		}

		return true, nil
	}

	return false, nil
}
