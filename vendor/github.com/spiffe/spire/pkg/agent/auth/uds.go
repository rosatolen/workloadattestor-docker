// +build linux

package auth

import (
	"net"
	"syscall"
)

// TODO: Figure out portability - can this work elsewhere? FreeBSD supports SO_PEERCRED
func FromUDSConn(conn net.Conn) CallerInfo {
	var info CallerInfo

	uconn, ok := conn.(*net.UnixConn)
	if !ok {
		info.Err = ErrInvalidConnection
		return info
	}

	file, err := uconn.File()
	if err != nil {
		info.Err = err
		return info
	}
	defer file.Close()

	ucred, err := syscall.GetsockoptUcred(int(file.Fd()), syscall.SOL_SOCKET, syscall.SO_PEERCRED)
	if err != nil {
		info.Err = err
		return info
	}

	info.Addr = uconn.RemoteAddr()
	info.PID = int32(ucred.Pid)
	return info
}
