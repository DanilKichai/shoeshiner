package dhcp4

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const FilterString = "ip and udp and src port 67 and dst port 68"

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
	if dhcp4Layer := packet.Layer(layers.LayerTypeDHCPv4); dhcp4Layer != nil {
		dhcp4, _ := dhcp4Layer.(*layers.DHCPv4)

		var msgType layers.DHCPMsgType
		for _, opt := range dhcp4.Options {
			if opt.Type == layers.DHCPOptMessageType && len(opt.Data) == 1 {
				msgType = layers.DHCPMsgType(opt.Data[0])

				break
			}
		}

		if dhcp4.Operation != layers.DHCPOpReply || msgType != layers.DHCPMsgTypeAck {
			return false, nil
		}

		if err := os.MkdirAll(filepath.Join(prefix, "options"), 0755); err != nil {
			return false, fmt.Errorf("create directory structure: %v", err)
		}

		if err := os.WriteFile(filepath.Join(prefix, "nextserver"), dhcp4.NextServerIP, 0644); err != nil {
			return false, fmt.Errorf("write nextserver file: %v", err)
		}

		for _, option := range dhcp4.Options {
			if err := os.WriteFile(filepath.Join(prefix, "options", strconv.Itoa(int(option.Type))), option.Data, 0644); err != nil {
				return false, fmt.Errorf("write option data file: %v", err)
			}
		}

		return true, nil
	}

	return false, nil
}
