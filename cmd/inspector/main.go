package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"shoeshiner/internal/app/inspector/capture"
	"shoeshiner/internal/app/inspector/dhcp4"
	"shoeshiner/internal/app/inspector/dhcp6"
	"shoeshiner/internal/app/inspector/systemd/notify"
	"time"

	"github.com/google/gopacket"
)

func main() {
	proto := flag.String("protocol", "dhcp4", "inspection protocol (dhcp4|dhcp6)")
	iface := flag.String("interface", "eth0", "interface to capture packets from")
	promisc := flag.Bool("promiscuous", false, "promiscuous mode")
	snaplen := flag.Int("length", 1600, "packet capture snapshot size")
	timeout := flag.Int("timeout", 5, "packet capture aggregation timeout seconds")
	dir := flag.String("directory", "/var/lib/inspector", "inspector logging directory")

	flag.Parse()

	prefix := filepath.Join(*dir, *proto)

	var inspect func(packet gopacket.Packet, dir string) (bool, error)
	filter := ""

	log.Printf("carrying out preparations for the protocol: %s", *proto)

	switch *proto {
	case "dhcp4":
		filter = dhcp4.FilterString
		inspect = dhcp4.Inspect
		dhcp4.Prepare(prefix)

	case "dhcp6":
		filter = dhcp6.FilterString
		inspect = dhcp6.Inspect
		dhcp6.Prepare(prefix)

	default:
		log.Fatalf("invalid protocol: %s", *proto)
	}

	log.Printf("creating BPF capture handle for interface: %s", *iface)

	captureHandle, err := capture.Start(*iface, *promisc, filter, *snaplen, time.Duration(*timeout)*time.Second)
	if err != nil {
		log.Fatalf("start capture: %v", err)
	}
	defer captureHandle.Close()

	log.Printf("start event listening for capture")

	packetSource := gopacket.NewPacketSource(captureHandle, captureHandle.LinkType())

	dNotify := notify.Create()

	log.Printf("inform systemd daemon (READY=1)")

	if err = dNotify.Ready(); err != nil {
		log.Fatalf("notify systemd of the ready status: %v", err)
	}

	for packet := range packetSource.Packets() {
		log.Printf("got packet with BPF filter")

		succ, err := inspect(packet, prefix)
		if err != nil {
			log.Fatalf("inspection: %v", err)
		}

		if succ {
			log.Printf("parsed data was written successfully")

			os.WriteFile(filepath.Join(prefix, "trigger"), []byte("activate\n"), 0600)
			log.Printf("trigger fifo was processed")

			return
		}
	}
}
