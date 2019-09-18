package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	da, err := ioutil.ReadFile("./localhost.cert")
	if err != nil {
		log.Println(err, "Could not read cert")
		return
	}
	rootCert := string(da)
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootCert))
	if !ok {
		log.Fatal("failed to parse root certificate")
	}
	config := &tls.Config{RootCAs: roots}

	c, err := tls.Dial("tcp", "localhost:8081", config)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	client, err := NewClient(c)
	if err != nil {
		return
	}
	var cmd *Cmd
	var data []byte
	for i := 0; i <= 100; i++ {
		cmd = client.Command("echo", "hi")
		data, err = cmd.CombinedOutput()
		if err != nil {
			if err1, ok := err.(*cmdError); ok {
				fmt.Println("cmdError source", err1.source, err1)
			} else {
				fmt.Println("normal err", err)
			}
			return
		}
		log.Println("Output", string(data))
		time.Sleep(time.Millisecond)
	}
	return
}
