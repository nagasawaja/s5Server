package main

import (
	socks5 "lakasocks5/s5"
	"log"
)

func main() {
	cred := socks5.StaticCredentials{
		"user": "passwd",
	}
	cator := socks5.UserPassAuthenticator{Credentials: cred}
	// Create a SOCKS5 server
	conf := &socks5.Config{AuthMethods: []socks5.Authenticator{cator}}
	server, err := socks5.New(conf)
	if err != nil {
		log.Printf("new server panic; err:%s", err.Error())
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", ":8000"); err != nil {
		log.Printf("listen serve fail; err:%s", err.Error())
		panic(err)
	}
}
