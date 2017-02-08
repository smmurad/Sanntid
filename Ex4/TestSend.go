package main

import (
	"./network"
	"fmt"
	//"strings"
)

func main() {
	port := ":20010"
	localAdress, _ := network.LocalIP()
	localAdress = localAdress + port
	sendSocket := network.Udp_create_send_socket(localAdress)
	network.Udp_send(sendSocket, "Hello me")
	fmt.Println("Sending complete")

	tcp_socket := network.Tcp_connect_request(port)

	network.Tcp_send(tcp_socket, "Hei, connected")

}
