package main

import (
    "log"
    "net"
    "fmt"
)

func main(){
    ListenAddr, err := net.ResolveUDPAddr("udp", ":30000")
    if err != nil{
        log.Fatal(err)
    }
    
    buffer := make([]byte,1024)
    listenConn, err := net.ListenUDP("udp", ListenAddr)
    if err != nil{
        log.Fatal(err)
    }

    fmt.Println("IP: ", listenConn)

    defer listenConn.Close()
    
    //for{
        n_bytes, addr, err := listenConn.ReadFromUDP(buffer)
        fmt.Println("Received: ", string(buffer[0:n_bytes]), " from ", addr)
        
        if err != nil{
            log.Fatal(err)
        }
    //}

        //server ip: "129.241.187.43:51504"
        //129.241.187.156/24
}