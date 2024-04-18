//go:build !linux

package dialer

import (
	"github.com/zhaofenghao/clash_go/log"
	"net"
	"sync"
)

var printMarkWarnOnce sync.Once

func printMarkWarn() {
	printMarkWarnOnce.Do(func() {
		log.Warnln("Routing mark on socket is not supported on current platform")
	})
}

func bindMarkToDialer(mark int, dialer *net.Dialer, _ string, _ net.IP) {
	printMarkWarn()
}

func bindMarkToListenConfig(mark int, lc *net.ListenConfig, _, address string) {
	printMarkWarn()
}
