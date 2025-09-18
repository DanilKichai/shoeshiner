package notify

import (
	"fmt"
	"net"
	"os"
)

const socketVariableName = "NOTIFY_SOCKET"

type NotifySocket net.UnixAddr

func Create() *NotifySocket {
	socket := (*NotifySocket)(nil)

	if notifySocketName := os.Getenv(socketVariableName); notifySocketName != "" {
		socket = &NotifySocket{
			Name: notifySocketName,
			Net:  "unixgram",
		}
	}

	return socket
}
func (socket *NotifySocket) Send(state string) error {
	if socket == nil {
		return nil
	}

	conn, err := net.DialUnix(socket.Net, nil, (*net.UnixAddr)(socket))
	if err != nil {
		return fmt.Errorf("dial unix socket: %v", err)
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(state)); err != nil {
		return fmt.Errorf("write into the unix socket: %v", err)
	}

	return nil
}
