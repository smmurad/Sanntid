package main


import (
	"fmt"
	"log"
	"net"
	"time"
	"bufio"
	"os"
)

func TCP_recieve(conn net.Conn) string{
	//Wait for messafe ending in '\0'
	msg, _ := bufio.NewReader(conn).ReadString(byte('\x00'))
	return msg
}

func TCP_send(conn net.Conn, msg string){
	conn.Write([]byte(msg + string('\x00')))
}


func main() {

	serverAddr := "129.241.187.43:33546"
	localAddr := ":34588"

	fmt.Println("Launching serverino ...")
	input := bufio.NewReader(os.Stdin)

	//Get a servers TCP addr
	TCPAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil{
		log.Fatal("ResolveTCPAddr failed: ", err.Error())
	}

	//get local tcp addr
	ListenAddr, err := net.ResolveTCPAddr("tcp", localAddr)

	//Set up a listener socket
	listener, err := net.ListenTCP("tcp", ListenAddr)

	//Connect to tcp server
	conn, err := net.DialTCP("tcp", nil, TCPAddr)
	if err != nil{
		log.Fatal("DialTCPAddr failed: ", err.Error())
	}
	//fmt.Printf("conn: %s\n", conn.LocalAddr());

	//Recieve welcome msg
	msg_rcpt := TCP_recieve(conn)
	fmt.Println(msg_rcpt)

	//Calling for a connection from server
	msg_connect := "Connect to: 129.241.187.156:34588"
	TCP_send(conn, msg_connect)
	
	defer conn.Close()

	//Listen for connection from server
	connServer, err := listener.AcceptTCP()
	//Recieve welcome msg
	msg_rcpt = TCP_recieve(connServer)
	fmt.Println(msg_rcpt)

	defer connServer.Close()



	for{
		//sends message to server

		fmt.Print("Enter message to send: ...")
		msg_send, _ := input.ReadString('\n')

		TCP_send(connServer, msg_send)

		msg_rcpt := TCP_recieve(connServer)

		//Print msg

		fmt.Println("Message Recieved: ", string(msg_rcpt))

		//sleep
		time.Sleep(1)

	}

}