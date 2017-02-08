package runNetwork

import (
	"./network/bcast"
	"./network/localip"
	"./network/peers"
)

type MyString struct {
	Message string
	Iter    int
}

const PeerPort := 20009
const BcastPort := 30009

func RunNetwork(){
		if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	}

		// We make a channel for receiving updates on the id's of the peers that are
	//  alive on the network
	peerUpdateCh := make(chan peers.PeerUpdate)
	// We can disable/enable the transmitter after it has been started.
	// This could be used to signal that we are somehow "unavailable".
	peerTxEnable := make(chan bool)
	go peers.Transmitter(msgPort, id, peerTxEnable)
	go peers.Receiver(msgPort, peerUpdateCh)

}