package notify

import "fmt"

func (socket *NotifySocket) IsEnabled() bool {
	return socket != nil
}

func (socket *NotifySocket) Ready() error {
	return socket.Send("READY=1")
}

func (socket *NotifySocket) Reloading() error {
	return socket.Send("RELOADING=1")
}

func (socket *NotifySocket) Stopping() error {
	return socket.Send("STOPPING=1")
}

func (socket *NotifySocket) Status(status string) error {
	return socket.Send(fmt.Sprintf("STATUS=%s", status))
}

func (socket *NotifySocket) ErrNo(errno int) error {
	return socket.Send(fmt.Sprintf("ERRNO=%d", errno))
}

func (socket *NotifySocket) BusError(buserror string) error {
	return socket.Send(fmt.Sprintf("BUSERROR=%s", buserror))
}

func (socket *NotifySocket) MainPID(mainpid int) error {
	return socket.Send(fmt.Sprintf("MAINPID=%d", mainpid))
}

func (socket *NotifySocket) WatchDog() error {
	return socket.Send("WATCHDOG=1")
}

func (socket *NotifySocket) WatchDogUSec(usec int64) error {
	return socket.Send(fmt.Sprintf("WATCHDOG_USEC=%d", usec))
}
