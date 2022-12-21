package server

import (
	"crypto/tls"
	"fmt"
	"github.com/jcatala/ghdumper/pkg/config"
	"log"
	"net"
	"strconv"
)



func Serve(conf *config.Config){
	// listen and serve
	p := ":" + strconv.Itoa(int(conf.Port))
	ln, err := net.Listen("tcp", p)
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
		if conf.ResponseFile != "" {
			go handleConnectionResponseFromFile(conn, conf)
			continue
		}
		go handleConnection(conn, conf)
	}
}


func ServeSSL(conf *config.Config){
	// SSL things
	cert, err := tls.LoadX509KeyPair(conf.Pemfile, conf.Key)
	if err != nil{
		fmt.Println("Error on getting the certificates")
		log.Fatalln(err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	// listen and serve
	p := ":" + strconv.Itoa(int(conf.Port))
	ln, err := tls.Listen("tcp", p , tlsConfig)
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
		if conf.ResponseFile != "" {
			go handleConnectionResponseFromFile(conn, conf)
			continue
		}
		go handleConnection(conn, conf)
	}
}
