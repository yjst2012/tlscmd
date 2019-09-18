package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	da1, err := ioutil.ReadFile("./localhost.cert")
	if err != nil {
		log.Println(err, "Could not read cert")
		return
	}
	serverCert := string(da1)

	da2, err := ioutil.ReadFile("./localhost.key")
	if err != nil {
		log.Println(err, "Could not read key")
		return
	}
	serverKey := string(da2)

	cer, err := tls.X509KeyPair([]byte(serverCert), []byte(serverKey))
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	fmt.Println("Launching server...")
	l, err := tls.Listen("tcp", ":8081", config)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			err := NewServer(c, nil)
			if err != nil {
				log.Println("error happen", err)
			}
			c.Close()
		}(conn)
	}
	return
}
