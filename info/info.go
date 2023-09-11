package info

import (
	"log"
	"net"
	"os"
)

func GetHostName() (string, error) {
	hostName, error := os.Hostname()

	return hostName, error

}

func GetIPAddress() (net.IP, error) {
	connection, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		log.Fatal(error)
	}
	defer connection.Close()
	localAddress := connection.LocalAddr().(*net.UDPAddr)
	return localAddress.IP, error
}
