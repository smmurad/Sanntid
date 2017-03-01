package main

import (
	//"./network/peers"
	//"./network/runNetwork"
	"../setup"
	"../hardware"
	"../fsm"
	"fmt"
	//"time"
)

type HelloMsg struct {
	Message string
	Iter    int
}

var ch_button_polling = make(chan setup.ButtonStruct)

 

func main(){
	
	hardware.Init()
	hardware.SetButtonLamp(setup.ButtonCommand, 1, 1)
	go hardware.ReadButtons(ch_button_polling)
	go fsm.FsmOrderHandler(ch_button_polling)
	fmt.Println("Kjører test")

	for{
		select{
			//case msg := <- ch_button_polling:
				//fmt.Println(msg)
		}
	}


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

