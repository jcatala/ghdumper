package main

import (
	"flag"
	"fmt"
	"github.com/jcatala/ghdumper/pkg/server"
	"github.com/jcatala/ghdumper/pkg/config"
)



func main() {
	ssl := flag.Bool("ssl", false, "Usage of SSL (use with -perm and -key).")
	pemfile := flag.String("pem","cert.pem", "Path of the cert file.")
	key := flag.String("key","key.pem", "Path of the private key file.")
	verbose := flag.Bool("verbose",false, "Verbose output")
	port := flag.Int64("port", 9001, "Port to use")
	fromFile := flag.String("request","", "File to send as request response")
	autoCL := flag.Bool("autocl", true, "Auto calculate content length based on the request length, default: False")
	flag.Parse()
	run := config.Config{
		Port:    *port,
		Ssl:     *ssl,
		Pemfile: *pemfile,
		Key:     *key,
		Verbose: *verbose,
		ResponseFile: *fromFile,
		AutoCL: *autoCL,
	}
	if run.Verbose{
		fmt.Println("Starting...")
	}
	if run.Ssl {
		if *verbose{
			fmt.Println("Serving using ssl")
		}
		server.ServeSSL(&run)
	}
	if *verbose {
		fmt.Println("Serving without SSL")
	}
	server.Serve(&run)

}