package daemon

import "github.com/openshift/linuxptp-daemon/pkg/config"
import "net"

type process interface {
	Name() string
	Stopped() bool
	CmdStop()
	CmdInit()
	ProcessStatus(int64)
	CmdRun(c **net.Conn)
	MonitorProcess(p config.ProcessConfig)
	ExitCh() chan struct{}
}
