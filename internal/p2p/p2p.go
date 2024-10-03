package p2p

import (
	"fmt"
	"net"
)

func StartNode(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Erro ao iniciar :", err)
		return
	}
	defer ln.Close()
	fmt.Println("Listando a port:", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Nova conexão estabelecida")
	conn.Close()
}
