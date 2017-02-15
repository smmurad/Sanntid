package main

import (
	"log"
	"net"
	"encoding/binary"
	"time"
	"os/exec"
	)

func listen(UDP *net.UDPConn, ch chan<- int){
	buffer := make([]byte, 1024)
	for{
		UDP.ReadFromUDP(buffer[:])

		rec, _ := binary.Uvarint(buffer)
		ch <- int(rec)
		time.Sleep(time.Millisecond * 100)
	}
}

func backup(UDP *net.UDPConn) int{
	lastVal := 0
	ch := make(chan int)
	go listen(UDP, ch)
	for{
		select{
			case lastVal = <- ch:
				break;
			case <-time.After(3*time.Second):
				return lastVal

		}
	}
}

func main() {
	port := ":30009"
	log.Printf("Hei, fra backup")
	
	addr, _ := net.ResolveUDPAddr("udp", port)
	listen, _ := net.ListenUDP("udp", addr)

	counter := backup(listen) + 1
	listen.Close()

	UDPaddr, err := net.ResolveUDPAddr("udp", "localhost" + port)
	if err != nil{
		log.Printf(err.Error())
	}

	connection, err := net.DialUDP("udp", nil, UDPaddr)
	if err != nil{
		log.Printf("DialUDP error: %s", err.Error())
	}

	backup := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run backup.go")
	backup.Run()

	msg := make([]byte, 1)

	for {
		log.Printf("%d", counter)

		msg[0] = byte(counter)
		_, err = connection.Write(msg)
		if err != nil{
			log.Printf("UDP write error: %s", err.Error())
		}
		counter++
		time.Sleep(time.Second)
	}	
}