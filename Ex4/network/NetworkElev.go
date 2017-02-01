package network

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

///////TCP//////

func Tcp_send(conn net.Conn, msg string) {
	conn.Write([]byte(msg + string('\x00')))
}

func Tcp_recieve(conn net.Conn) string {
	//Wait for messafe ending in '\0'
	msg, _ := bufio.NewReader(conn).ReadString(byte('\x00'))
	return msg
}

func Tcp_connect(ip_adress string) *net.TCPConn {
	TCPAddr, err := net.ResolveTCPAddr("tcp", ip_adress)
	if err != nil {
		fmt.Println("ResolveTCPAddr failed: ", err.Error())
	}
	//Connect to tcp server
	conn, err := net.DialTCP("tcp", nil, TCPAddr)
	if err != nil {
		fmt.Println("DialTCPAddr failed: ", err.Error())
	}
	return conn
}

func Tcp_accept_connection(port string) *net.TCPConn {
	//get local tcp addr
	tcp_port, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		fmt.Println("ResolveTCPAddr failed: ", err.Error())
	}
	listener, err := net.ListenTCP("tcp", tcp_port)
	if err != nil {
		fmt.Println("Listen TCP failed: ", err.Error())
	}
	conn, err := listener.AcceptTCP()
	if err != nil {
		fmt.Println("AcceptTCP() failed: ", err.Error())
	}
	return conn
}

//////UDP//////

func Udp_create_send_socket(ip_adress string) net.Conn {
	//set up sender socket
	remoteAddr, err := net.ResolveUDPAddr("udp", ip_adress)
	if err != nil {
		fmt.Println(err.Error()) //kanskje
	}
	socketSend, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		fmt.Println(err.Error())
	}
	return socketSend
}

func Udp_create_recieve_socket(port string) *net.UDPConn {
	//set up sender socket
	udp_port, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Println(err.Error()) ////????? should you exit?
	}
	socketRecieve, err := net.ListenUDP("udp", udp_port)
	if err != nil {
		fmt.Println(err.Error())
	}
	return socketRecieve
}

func Udp_broadcast(conn net.Conn, message string) {
	//?
}

func Udp_send(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message + "\x00"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Udp_recieve(conn *net.UDPConn) string {
	buffer := make([]byte, 1024)
	n_bytes, _, _ := conn.ReadFromUDP(buffer)
	return string(buffer[0:n_bytes])
}

/////////Networking

func LocalIP() (string, error) {
	var localIP string
	if localIP == "" {
		conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: []byte{8, 8, 8, 8}, Port: 53})
		if err != nil {
			return "", err
		}
		defer conn.Close()
		localIP = strings.Split(conn.LocalAddr().String(), ":")[0]
	}
	return localIP, nil
}

func Tcp_connect_request(port string) *net.TCPConn {
	local_ip, _ := LocalIP()
	broadcast_ip := "129.241.187.255" + port
	sender_socket := Udp_create_send_socket(broadcast_ip)

	message := "IP;" + local_ip
	Udp_send(sender_socket, message)

	reciever_socket := Udp_create_recieve_socket(port)

	connected := false
	for !connected {
		recieved_message := Udp_recieve(reciever_socket)
		message_split := strings.Split(recieved_message, "!")
		if message_split[0] == "ok" {
			connection_ip := message_split[1]
			connection := Tcp_connect(connection_ip + port)
			connected = true
			return connection
		}
	}
	return nil
}

func Tcp_handle_requests(port string) *net.TCPConn {
	local_ip, _ := LocalIP()
	reciever_socket := Udp_create_recieve_socket(port)
	for true {
		recieved_message := strings.Split(Udp_recieve(reciever_socket), ";")
		if recieved_message[0] == "IP" {
			sender_socket := Udp_create_send_socket(recieved_message[1])
			accept_message := "ok!" + local_ip
			Tcp_send(sender_socket, accept_message)
			connection := Tcp_accept_connection(port)
			defer sender_socket.Close()
			return connection
		}
	}
	return nil
}

//json
