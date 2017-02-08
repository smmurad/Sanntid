package main

import (
	"./network"
	"fmt"
)

func main() {
	port := ":20010"
	recieverSocket := network.Udp_create_recieve_socket(port)
	fmt.Println(network.Udp_recieve(recieverSocket))

	tcp_socket := network.Tcp_handle_requests(port)

	recieved_message := network.Tcp_recieve(tcp_socket)
	fmt.Println(recieved_message)
}
