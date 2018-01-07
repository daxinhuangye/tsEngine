package tsMac

import (
	"net"
)

func MacList()(data []string, err error){
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	
	for _, inter := range interfaces {
		data = append(data, inter.HardwareAddr.String())
	}
	return
}