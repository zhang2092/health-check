package internal

import (
	"fmt"
	"net"
	"time"
)

func Check(destination, port string) string {
	address := destination + ":" + port
	timeout := 5 * time.Second
	var status string
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		status = fmt.Sprintf("[下线] %s 无法访问 %v", destination, err)
		return status
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	status = fmt.Sprintf("[上线] %s 访问正常\n从: %s\n到: %s", destination, conn.LocalAddr(), conn.RemoteAddr())
	return status
}
