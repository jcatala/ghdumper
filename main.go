package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
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

func main() {
	ssl := flag.Bool("ssl", false, "Usage of SSL (use with -perm and -key).")
	pemfile := flag.String("pem","cert.pem", "Path of the cert file.")
	key := flag.String("key","key.pem", "Path of the private key file.")
	if *ssl {
		handleSsl(*pemfile, *key)
	}

}

func handleConnection(conn net.Conn) {
	// Use bufio to read the request
	bufReader := bufio.NewReader(conn)
	req, err := http.ReadRequest(bufReader)
	if err != nil {
		// handle error
		return
	}

	// Print out the request method and URL
	fmt.Println("Method:", req.Method)
	fmt.Println("URL:", req.URL)

	// Write a response back to the client
	resp := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     make(http.Header),
	}
	//resp.Header.Set("Content-Type", "text/plain")
	//resp.Header.Set("Content-Length", "12")
	body := `deadbeef`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: 4\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
	resp.Write(conn)
}