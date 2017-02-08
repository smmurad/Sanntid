package runNetwork

import (
	"../bcast"
	"../localip"
	"../peers"
	"fmt"
	"os"
)

type ButtonType int
const(
	ButtonUp = 0
	ButtonDown = 1
	ButtonCommand = 2
)

type ButtonStruct struct{
	Button_Type ButtonType
	Floor int
}

type MyOrder struct {
	Button ButtonStruct
	Id int
}



const PeerPort = 20009
const BcastPort = 30009

func RunPeerNetwork(peerUpdateCh chan<- peers.PeerUpdate, peerTxEnable <-chan bool){
		
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id := fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())

	go peers.Transmitter(PeerPort, id, peerTxEnable)
	go peers.Receiver(PeerPort, peerUpdateCh)
}

func RunMessageNetwork(orderTxCh chan MyOrder, orderRxCh chan MyOrder){
	go bcast.Transmitter(BcastPort, orderTxCh)
	go bcast.Receiver(BcastPort, orderRxCh)
}