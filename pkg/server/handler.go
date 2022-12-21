package server

import (
	"bufio"
	"fmt"
	"github.com/jcatala/ghdumper/pkg/config"
	"net"
	"net/http"
	"os"
)


func handleConnection(conn net.Conn, conf *config.Config){
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
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
	resp.Write(conn)
}

func handleConnectionResponseFromFile(conn net.Conn, conf *config.Config){
	// Use bufio to read the request
	bufReader := bufio.NewReader(conn)
	req, err := http.ReadRequest(bufReader)
	if err != nil {
		// handle error
		return
	}

	// Print out the request method and URL
	if conf.Verbose{
		fmt.Println("Method:", req.Method)
		fmt.Println("URL:", req.URL)
		fmt.Println("Connection from: ", req.RemoteAddr)
		fmt.Println("Request:")
	}
	rawResp, err := os.ReadFile(conf.ResponseFile)
	if err != nil{
		fmt.Println("Error reading the file")
		return
	}
	if conf.Verbose{
		fmt.Println("Sending: \n")
		fmt.Println(string(rawResp))
	}
	/* This can be used on normal handler
	if !conf.AutoCL {
		// Adding 2 to not count the \n\n
		pos := bytes.Index(rawResp, []byte("\n\n")) + 2
		length := len(rawResp[pos:len(rawResp)-1])
		fmt.Println("Position at: ", pos)
		fmt.Println("Content length used: " , length)
	}*/
	// Write a response back to the client
	//resp := http.Response{}
	//resp.Header.Set("Content-Type", "text/plain")
	//resp.Header.Set("Content-Length", "12")
	//fmt.Fprint(conn, string(rawResp))
	conn.Write(rawResp)
	conn.Close()
}