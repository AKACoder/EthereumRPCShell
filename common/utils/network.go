package utils

import (
	"fmt"
	"net"
)

const maxPort = 65535

func PortInUse(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return true
	}
	defer l.Close()
	return false
}

func FindAvailablePort(start int, toFile string) string {
	port := start
	for PortInUse(port) {
		port++

		if port > maxPort {
			return ""
		}
	}

	portStr := fmt.Sprintf("%d", port)
	if toFile != "" {
		SaveFile(toFile, []byte(portStr))
	}

	return ":" + portStr
}
