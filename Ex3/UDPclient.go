package main


import (
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	serverAddr := "129.241.187.43:20009"


	//set up sender socket
	remoteAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil{
		log.Fatal(err)
	}
	socketSend, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil{
		log.Fatal(err)
	}

	//Set up recieve socket
	port, err := net.ResolveUDPAddr("udp", ":20009")
	if err != nil{
		log.Fatal(err)
	}
	socketRecieve, err := net.ListenUDP("udp", port)
	if err != nil{
		log.Fatal(err)
	}

	defer socketRecieve.Close()
	defer socketSend.Close()

	msg := "Hello terra"
	_, err = socketSend.Write([]byte(msg + "\x00"))
	if err != nil{
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)

	for{
		n_bytes, _, _ := socketRecieve.ReadFromUDP(buffer)
		fmt.Println(string(buffer[0:n_bytes]))

		msg := "Hello terra"
		_, err := socketSend.Write([]byte(msg + "\x00"))
		if err != nil{
			log.Fatal(err)
		}

		time.Sleep(time.Second)

	}
}