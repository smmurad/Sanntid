package main

/*
#cgo CFLAGS: -std=gnu99
#cgo LDFLAGS: -lcomedi -lm
#include "hardware/io.c"
#include "hardware/elev.c"
#include "hardware/main.c"
*/
import "C"

import (
	//"./network/peers"
	//"./network/runNetwork"
	//"./network/setup"
	"fmt"
	//"time"
)

type HelloMsg struct {
	Message string
	Iter    int
}
 

func main(){
	
	fmt.Println("Kjører test")
	C.testMain();
	//C.elev_init()

	/*
	peerUpdateCh := make(chan peers.PeerUpdate)
	peerTxEnableCh := make(chan bool)
	orderTxCh := make(chan setup.MyOrder)
	orderRxCh := make(chan setup.MyOrder)

	runNetwork.RunPeerNetwork(peerUpdateCh, peerTxEnableCh)
	runNetwork.RunMessageNetwork(orderTxCh, orderRxCh)

	buttonStruct := setup.ButtonStruct{Button_Type:1, Floor:2}
	newOrder := setup.MyOrder{Button:buttonStruct, Id:15}

	go func() {
		for {
			orderTxCh <- newOrder
			time.Sleep(1 * time.Second)
		}
	}()
	

	fmt.Println("Started")
	for {
		select {
		case p := <-peerUpdateCh:
			fmt.Printf("Peer update:\n")
			fmt.Printf("  Peers:    %q\n", p.Peers)
			fmt.Printf("  New:      %q\n", p.New)
			fmt.Printf("  Lost:     %q\n", p.Lost)

		case a := <-orderRxCh:
			fmt.Printf("Received: %#v\n", a)
		}
	}*/
}

