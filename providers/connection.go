package providers

import (
	"crypto/tls"
	"fmt"
	"log"
	"github.com/jcatala/ghdumper/providers/"
)

func handleSsl(pem string, key string){
	// SSL things
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.unencrypted.pem")
	if err != nil{
		fmt.Println("Error on getting the certificates")
		log.Fatalln(err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	// listen and serve
	ln, err := tls.Listen("tcp", ":9001", tlsConfig)
	if err != nil {
		fmt.Println("Error while starting the server ")
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn)
	}
}
